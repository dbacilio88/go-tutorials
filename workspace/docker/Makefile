all: up

down:
	@docker compose down
	@docker image rm docker-server-ssh
	@docker image rm docker-postgres-db

up:
	@docker compose up -d

delete:
	@docker image rm docker-server-ssh