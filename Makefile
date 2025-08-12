.PHONY: build clean test tidy help

# Binary name
BINARY_NAME=kalco

# Build flags
LDFLAGS=-ldflags "-X main.Version=$(shell git describe --tags --always --dirty)"

# Default target
all: build

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	go build $(LDFLAGS) -o $(BINARY_NAME) .

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -f $(BINARY_NAME)
	rm -rf kalco-dump-*

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Tidy go modules
tidy:
	@echo "Tidying go modules..."
	go mod tidy
	go mod verify

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod download

# Run linter
lint:
	@echo "Running linter..."
	golangci-lint run

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Vet code
vet:
	@echo "Vetting code..."
	go vet ./...

# Build for multiple platforms
build-all: clean
	@echo "Building for multiple platforms..."
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY_NAME)-linux-amd64 .
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY_NAME)-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(BINARY_NAME)-darwin-arm64 .
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY_NAME)-windows-amd64.exe

# Show help
help:
	@echo "Available targets:"
	@echo "  build        - Build the binary"
	@echo "  clean        - Clean build artifacts"
	@echo "  test         - Run tests"
	@echo "  test-coverage- Run tests with coverage"
	@echo "  tidy         - Tidy go modules"
	@echo "  deps         - Install dependencies"
	@echo "  lint         - Run linter"
	@echo "  fmt          - Format code"
	@echo "  vet          - Vet code"
	@echo "  build-all    - Build for multiple platforms"
	@echo "  help         - Show this help"
