.DEFAULT_GOAL := help

.PHONY: run
run: ## Run the application passing the arguments
	@echo "> Running gotor..."
	go run cmd/main.go $(ARGS)

.PHONY: test
test: ## Run tests
	@echo "> Testing..."
	go test -v ./...

.PHONY: postgres-up
postgres-up: ## Create a postgres container for testing
	@echo "> Creating a postgres container for testing..."
	docker-compose -f deploy/docker-compose.yml up -d

.PHONY: postgres-down
postgres-down: ## Stop the postgres container for testing
	@echo "> Stopping the postgres container for testing..."
	docker-compose -f deploy/docker-compose.yml down

.PHONY: tidy
tidy: ## Clean and format Go code
	@echo "> Tidying..."
	go mod tidy
	go fmt ./...
	@echo "> Done!"

.PHONY: fmt
fmt: ## Format Go code
	go fmt ./...

.PHONY: lint-host
lint-host: ## Run golangci-lint directly on host
	@echo "> Linting..."
	golangci-lint run -c .golangci.yml -v
	@echo "> Done!"

.PHONY: help
help: ## Show this help
	@echo "make run - Run the application"
	@echo "make test - Run tests"
	@echo "make postgres-up - Create a postgres container for testing"
	@echo "make postgres-down - Stop the postgres container for testing"
	@echo "make tidy - Clean and format Go code"
	@echo "make fmt - Format Go code"
	@echo "make lint-host - Run golangci-lint directly on host"
	@echo "make help - Show this help"
