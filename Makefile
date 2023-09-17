include configs/.env

dep:
	@go mod tidy

init:
	@docker-compose up -d

run: dep
	@go run cmd/main.go

test: dep
	@go test ./...

help:
	@echo Below are the commands provided by the `Makefile` file.
	@echo Note: The commands are already arranged in the recommended execution order.
	@echo
	@echo `dep` - download `go.mod` dependicies. This is required command for a newly installed project.
	@echo `init` - download and start the MySQL server by Docker. This is required command for a newly installed project.
	@echo `run` - start the server.
	@echo `test` - run on the tests.