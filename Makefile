.PHONY: build clean run install test cross-compile

# Variables
BINARY_NAME=atm
VERSION=1.0.0
BUILD_DIR=bin
MAIN_PATH=cmd/atm/main.go

# Build flags
LDFLAGS=-ldflags="-s -w"

# Default target
all: build

# Build for current platform
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME).exe $(MAIN_PATH)
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME).exe"

# Run without building
run:
	@go run $(MAIN_PATH)

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete"

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy
	@echo "Dependencies installed"

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Cross-compile for multiple platforms
cross-compile:
	@echo "Cross-compiling for multiple platforms..."
	@mkdir -p $(BUILD_DIR)

	@echo "Building for Windows (amd64)..."
	@GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PATH)

	@echo "Building for Linux (amd64)..."
	@GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_PATH)

	@echo "Building for macOS (amd64)..."
	@GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(MAIN_PATH)

	@echo "Building for macOS (arm64)..."
	@GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 $(MAIN_PATH)

	@echo "Cross-compilation complete"

# Install to system (Windows)
install: build
	@echo "Installing to C:\Windows\System32..."
	@copy $(BUILD_DIR)\$(BINARY_NAME).exe C:\Windows\System32\
	@echo "Installation complete. You can now run 'atm' from anywhere."

# Display help
help:
	@echo "Available targets:"
	@echo "  build          - Build for current platform"
	@echo "  run            - Run without building"
	@echo "  clean          - Remove build artifacts"
	@echo "  deps           - Install dependencies"
	@echo "  test           - Run tests"
	@echo "  cross-compile  - Build for all platforms"
	@echo "  install        - Install to system (Windows)"
	@echo "  help           - Display this help message"
