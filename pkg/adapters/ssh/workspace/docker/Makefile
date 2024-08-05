all: up

down:
	@docker compose down
	@docker image rm docker-server-ssh

up:
	@docker compose up -d

delete:
	@docker image rm docker-server-ssh