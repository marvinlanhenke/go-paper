FRONTEND_DIR := client
BACKEND_DIR := server
INFRA_DIR := infra
ECR_REPO := ml-sa-go-paper-backend
BUCKET_NAME := ml-sa-s3-static-site
AWS_PROFILE := default
AWS_REGION := eu-central-1

API_URL := http://localhost:8080/v1

.PHONY: backend-build
backend-build:
	cd $(BACKEND_DIR) && \
	docker build -t $(ECR_REPO):latest .

.PHONY: backend-push
backend-push:
	URL=$$(cd $(INFRA_DIR) && terraform output -raw ecr_repository_url) && \
	aws ecr get-login-password --region $(AWS_REGION) | docker login --username AWS --password-stdin $$URL && \
	docker tag $(ECR_REPO):latest $$URL:latest && \
	docker push $$URL:latest

.PHONY: frontend-build
frontend-build:
	cd $(FRONTEND_DIR) && \
	VITE_API_URL=$(API_URL) \
	npm run build

.PHONY: frontend-preview
frontend-preview:
	cd $(FRONTEND_DIR) && \
	npm run preview

.PHONY: upload-s3
upload-s3:
	aws s3 sync ./client/dist s3://$(BUCKET_NAME)/ --delete
