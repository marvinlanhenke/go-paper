.PHONY: docker
docker:
	docker compose down --remove-orphans && \
	docker rmi ml/go-paper; \
	docker compose build --force-rm && \
	docker image prune -f && \
	docker compose up -d

