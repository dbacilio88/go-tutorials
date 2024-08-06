up:
	docker compose -f workspace/docker/docker-compose.yml up -d

down:
	docker compose -f workspace/docker/docker-compose.yml down
	@docker image rm docker-server-ssh
	@docker image rm docker-postgres-db

create:
	migrate create -ext sql -dir workspace/docker/db/migration -seq init_schema

run:
	migrate -path workspace/docker/db/migration -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose up

delete:
	migrate -path workspace/docker/db/migration -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose down

generate:
	#sqlc generate -f ./workspace/docker/db/sqlc.yml
	sqlc generate

tests:
	go test -v -cover ./...

.PHONY: up down run delete