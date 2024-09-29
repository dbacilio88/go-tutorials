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

mock:
	mockgen -source pkg/adapters/db/db.go -destination mocks/pkg/adapters/db/db.go
	#mockgen -source pkg/adapters/db/modles.go -destination mocks/pkg/adapters/db/models.go
	#mockgen -source pkg/adapters/db/store.go -destination mocks/pkg/adapters/db/store.go
	#mockgen -source pkg/adapters/db/querier.go -destination mocks/pkg/adapters/db/querier.go
	mockgen -source pkg/adapters/ssh/ssh.go -destination mocks/pkg/adapters/ssh/ssh.go
	mockgen -source pkg/adapters/ftp/ftp.go -destination mocks/pkg/adapters/ftp/ftp.go
	mockgen -source pkg/adapters/mq/rabbitmq.go -destination mocks/pkg/adapters/mq/rabbitmq.go
	mockgen -source pkg/components/helpers/random.go -destination mocks/pkg/components/helpers/random.go
	mockgen -source pkg/config/configuration.go -destination mocks/pkg/config/configuration.go
	mockgen -source pkg/server/server.go -destination mocks/pkg/server/server.go
	mockgen -source pkg/task/task.go -destination mocks/pkg/task/task.go

.PHONY: up down run delete