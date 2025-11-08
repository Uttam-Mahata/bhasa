# Makefile for building Bhasa binaries for multiple platforms

# Binary name
BINARY_NAME=bhasa

# Version info (can be overridden)
VERSION?=latest

# Build directory
BUILD_DIR=bin

# Go build flags
LDFLAGS=-ldflags "-s -w"

# Platforms to build for
.PHONY: all clean linux windows darwin linux-amd64 linux-arm64 windows-amd64 windows-arm64 darwin-amd64 darwin-arm64 help

help: ## Show this help message
	@echo "Bhasa Build System - Available targets:"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

all: linux windows darwin ## Build binaries for all platforms

linux: linux-amd64 linux-arm64 ## Build all Linux binaries

windows: windows-amd64 windows-arm64 ## Build all Windows binaries

darwin: darwin-amd64 darwin-arm64 ## Build all macOS binaries

linux-amd64: ## Build for Linux AMD64
	@echo "Building for Linux AMD64..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 .

linux-arm64: ## Build for Linux ARM64
	@echo "Building for Linux ARM64..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 .

windows-amd64: ## Build for Windows AMD64
	@echo "Building for Windows AMD64..."
	@mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe .

windows-arm64: ## Build for Windows ARM64
	@echo "Building for Windows ARM64..."
	@mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-arm64.exe .

darwin-amd64: ## Build for macOS AMD64 (Intel)
	@echo "Building for macOS AMD64..."
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 .

darwin-arm64: ## Build for macOS ARM64 (Apple Silicon)
	@echo "Building for macOS ARM64..."
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 .

build: ## Build for current platform
	@echo "Building for current platform..."
	go build $(LDFLAGS) -o $(BINARY_NAME) .

clean: ## Remove build artifacts
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@rm -f $(BINARY_NAME) $(BINARY_NAME).exe

test: ## Run tests
	@echo "Running tests..."
	go test -v ./...

.DEFAULT_GOAL := help
