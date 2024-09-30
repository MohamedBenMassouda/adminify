# Variables
APP_NAME ?= adminify           # Default app name if not provided
BUILD_DIR = bin
TEST_DIR = ./...
GO_FILES = $(shell find . -name '*.go')

# Default target
.PHONY: help
help:  ## Show help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: $(BUILD_DIR)  ## Build Go app
	@go build -o $(BUILD_DIR)/$(APP_NAME) || exit 1

$(BUILD_DIR):  ## Create bin directory if it doesn't exist
	@mkdir -p $(BUILD_DIR)

.PHONY: test
test:  ## Run tests
	@go test -v $(TEST_DIR) || exit 1

.PHONY: format
format:  ## Format Go code
	@go fmt $(GO_FILES)

.PHONY: lint
lint:  ## Run linter
	@golangci-lint run || exit 1

.PHONY: clean
clean:  ## Clean up
	@rm -rf $(BUILD_DIR)

.PHONY: tidy
tidy:  ## Clean up go.mod dependencies
	@go mod tidy


