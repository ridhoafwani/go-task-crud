# Include environment variables from .env file
include .env

# Export environment variables
export $(shell sed 's/=.*//' .env)

# Construct POSTGRESQL_URL
POSTGRESQL_URL := postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DATABASE)?sslmode=disable

# Build the application
all: build test

build:
	@echo "Building..."
	
	
	@go build -o main cmd/main.go

# Run the application
run:
	@go run cmd/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v


# Live Reload
watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

migrate-create:
	@ migrate create -ext sql -dir scripts/migrations -seq $(name)

migrate-up:
	@ echo "Connecting to PostgreSQL at ${POSTGRESQL_URL}"
	@ migrate -path scripts/migrations -database "${POSTGRESQL_URL}" -verbose up

migrate-down:
	@ migrate -path scripts/migrations -database "${POSTGRESQL_URL}" -verbose down

# Generate the swagger docs
swag:
	@swag init -g cmd/main.go

.PHONY: all build run test watch swag