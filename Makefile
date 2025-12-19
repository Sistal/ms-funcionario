.PHONY: help build run test docker-build docker-up docker-down clean

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the application
	go build -o bin/ms-funcionario ./cmd/api

run: ## Run the application locally
	go run ./cmd/api/main.go

test: ## Run tests
	go test -v ./...

docker-build: ## Build Docker image
	docker compose build

docker-up: ## Start services with Docker Compose
	docker compose up -d

docker-down: ## Stop services with Docker Compose
	docker compose down

docker-logs: ## Show Docker logs
	docker compose logs -f

clean: ## Clean build artifacts
	rm -rf bin/
	go clean
