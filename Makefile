API_URL=http://localhost:8080/v1
BUCKET_NAME=ml-sa-s3-static-site

.PHONY: backend
backend:
	docker compose down --remove-orphans && \
	docker rmi ml/go-paper; \
	docker compose build --force-rm && \
	docker image prune -f && \
	docker compose up -d

.PHONY: frontend-build
frontend-build:
	cd ./client && \
	VITE_API_URL=$(API_URL) \
	npm run build

.PHONY: frontend-preview
frontend-preview:
	cd ./client && \
	npm run preview

.PHONY: run-local
run-local: backend frontend-build frontend-preview


.PHONY: upload-s3
upload-s3:
	aws s3 sync ./client/dist s3://$(BUCKET_NAME)/ --delete
