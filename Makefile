BUCKET_NAME=ml-sa-s3-static-site

.PHONY: docker
docker:
	docker compose down --remove-orphans && \
	docker rmi ml/go-paper; \
	docker compose build --force-rm && \
	docker image prune -f && \
	docker compose up -d

.PHONY: upload-s3
upload-s3:
	aws s3 sync ./client/dist s3://$(BUCKET_NAME)/ --delete
