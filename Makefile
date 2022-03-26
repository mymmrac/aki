# Adds $GOPATH/bit to $PATH
export PATH := $(PATH):$(shell go env GOPATH)/bin

help: ## Display this help message
	@echo "Usage:"
	@grep -E "^[a-zA-Z_-]+:.*? ## .+$$" $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-24s\033[0m %s\n", $$1, $$2}'

lint: ## Run golangci-lint
	golangci-lint run

lint-install: ## Install golangci-lint
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

test: ## Run tests
	go test -coverprofile cover.out \
	$(shell go list ./...)

cover: test ## Run tests & show coverage
	go tool cover -func cover.out

race: ## Run tests with race flag
	go test -race -count=1 ./...

pre-commit: test lint ## Run tests and linter

.PHONY: help lint lint-install test cover race pre-commit
