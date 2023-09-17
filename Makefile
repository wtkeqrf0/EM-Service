include configs/.env

dep:
	@go mod download

init:
	@docker-compose up -d

run:
	@go run main.go

generate:
	@go generate ./...

test:
	@go test ./...

migration_up:
	migrate -path database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose up

migration_down:
	migrate -path database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose down

migration_fix:
	migrate -path database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" force VERSION

help:
	@echo Below are the commands provided by the `Makefile` file.
	@echo Note: The commands are already arranged in the recommended execution order.
	@echo
	@echo `dep` - download `go.mod` dependicies. This is required command for a newly installed project.
	@echo `init` - download and start the MySQL server by Docker. This is required command for a newly installed project.
	@echo `run` - start the gRPC server.
	@echo `test` - run on the tests.
	@echo `test_coverage` - run on the tests and generate coverage file.