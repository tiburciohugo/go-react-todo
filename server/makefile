.PHONY: default run build test docs clean
# Variables
APP_NAME=webserver

# Tasks
default: run

run:
	@echo "Running server..."
	@go run main.go
build:
	@echo "Building server..."
	@go build -o $(APP_NAME) main.go
test:
	@echo "Testing server..."
	@go test ./...
docs:
	@echo "Generating docs..."
	@swag init
clean:
	@echo "Cleaning up..."
	@rm -rf $(APP_NAME) docs
	@rm -rf ./docs
	@go clean
