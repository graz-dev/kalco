# Kalco Makefile
# This file provides common development and release commands

.PHONY: help build test clean release docker-build docker-push install uninstall lint security-check

# Variables
BINARY_NAME=kalco
VERSION=$(shell git describe --tags --always --dirty)
COMMIT=$(shell git rev-parse HEAD)
DATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
LDFLAGS=-s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)

# Default target
help:
	@echo "🚀 Kalco Development Commands"
	@echo "============================="
	@echo ""
	@echo "📦 Building:"
	@echo "  build          - Build kalco binary"
	@echo "  build-all      - Build for all platforms"
	@echo "  clean          - Clean build artifacts"
	@echo ""
	@echo "🧪 Testing:"
	@echo "  test           - Run tests"
	@echo "  test-race      - Run tests with race detection"
	@echo "  test-coverage  - Run tests with coverage"
	@echo ""
	@echo "🔧 Development:"
	@echo "  lint           - Run linters"
	@echo "  security-check - Run security scans"
	@echo "  install        - Install kalco to system"
	@echo "  uninstall      - Remove kalco from system"
	@echo ""
	@echo "🐳 Docker:"
	@echo "  docker-build   - Build Docker image"
	@echo "  docker-push    - Push Docker image"
	@echo ""
	@echo "🚀 Release:"
	@echo "  release        - Create a new release"
	@echo "  changelog      - Generate changelog"
	@echo ""
	@echo "📋 Information:"
	@echo "  version        - Show current version"
	@echo "  deps           - Show dependency information"

# Build targets
build:
	@echo "🔨 Building kalco..."
	go build -ldflags "$(LDFLAGS)" -o $(BINARY_NAME)
	@echo "✅ Build complete: $(BINARY_NAME)"

build-all: clean
	@echo "🔨 Building for all platforms..."
	@mkdir -p dist
	
	# Linux
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o dist/$(BINARY_NAME)-linux-amd64
	GOOS=linux GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o dist/$(BINARY_NAME)-linux-arm64
	
	# macOS
	GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o dist/$(BINARY_NAME)-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o dist/$(BINARY_NAME)-darwin-arm64
	
	# Windows
	GOOS=windows GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o dist/$(BINARY_NAME)-windows-amd64.exe
	GOOS=windows GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o dist/$(BINARY_NAME)-windows-arm64.exe
	
	@echo "✅ Multi-platform build complete in dist/ directory"

clean:
	@echo "🧹 Cleaning build artifacts..."
	@rm -f $(BINARY_NAME)
	@rm -rf dist/
	@rm -f *.tar.gz *.zip
	@rm -f coverage.txt
	@echo "✅ Clean complete"

# Test targets
test:
	@echo "🧪 Running tests..."
	go test -v ./...

test-race:
	@echo "🧪 Running tests with race detection..."
	go test -race -v ./...

test-coverage:
	@echo "🧪 Running tests with coverage..."
	go test -v -coverprofile=coverage.txt -covermode=atomic ./...
	@echo "📊 Coverage report generated: coverage.txt"



# Development targets
lint:
	@echo "🔍 Running linters..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "⚠️  golangci-lint not installed. Installing..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
		golangci-lint run; \
	fi

security-check:
	@echo "🔒 Running security checks..."
	@if command -v trivy >/dev/null 2>&1; then \
		trivy fs .; \
	else \
		echo "⚠️  Trivy not installed. Install from: https://aquasecurity.github.io/trivy/"; \
	fi

# Installation targets
install: build
	@echo "📦 Installing kalco to system..."
	@sudo mv $(BINARY_NAME) /usr/local/bin/
	@echo "✅ kalco installed to /usr/local/bin/"

uninstall:
	@echo "🗑️  Removing kalco from system..."
	@sudo rm -f /usr/local/bin/$(BINARY_NAME)
	@echo "✅ kalco removed from system"

# Docker targets
docker-build:
	@echo "🐳 Building Docker image..."
	docker build \
		--build-arg VERSION=$(VERSION) \
		--build-arg COMMIT=$(COMMIT) \
		--build-arg DATE=$(DATE) \
		-t kalco:$(VERSION) \
		-t kalco:latest .

docker-push:
	@echo "🐳 Pushing Docker image..."
	@echo "⚠️  Make sure you're logged in to your registry first"
	docker push kalco:$(VERSION)
	docker push kalco:latest

# Release targets
release:
	@echo "🚀 Creating release..."
	@echo "Current version: $(VERSION)"
	@echo "Current commit: $(COMMIT)"
	@echo ""
	@echo "To create a release:"
	@echo "1. git tag v1.0.0"
	@echo "2. git push origin v1.0.0"
	@echo "3. GitHub Actions will automatically build and release"
	@echo ""
	@echo "Or run: make release-local"

release-local: build-all
	@echo "🚀 Creating local release packages..."
	@cd dist && \
	for file in *; do \
		if [[ $$file == *.exe ]]; then \
			zip $${file%.exe}.zip $$file; \
		else \
			tar -czf $$file.tar.gz $$file; \
		fi; \
	done
	@echo "✅ Local release packages created in dist/ directory"

changelog:
	@echo "📋 Generating changelog..."
	@if command -v conventional-changelog >/dev/null 2>&1; then \
		conventional-changelog -p angular -i CHANGELOG.md -s; \
	else \
		echo "⚠️  conventional-changelog not installed. Install with: npm install -g conventional-changelog-cli"; \
	fi

# Information targets
version:
	@echo "📋 Version Information:"
	@echo "  Binary: $(BINARY_NAME)"
	@echo "  Version: $(VERSION)"
	@echo "  Commit: $(COMMIT)"
	@echo "  Date: $(DATE)"

deps:
	@echo "📦 Dependency Information:"
	@go mod graph
	@echo ""
	@echo "📊 Go version:"
	@go version
	@echo ""
	@echo "🔍 Outdated dependencies:"
	@go list -u -m all

# Development helpers
dev-setup:
	@echo "🔧 Setting up development environment..."
	@go mod download
	@go mod verify
	@echo "✅ Development environment ready"

fmt:
	@echo "🎨 Formatting code..."
	@go fmt ./...
	@echo "✅ Code formatted"

vet:
	@echo "🔍 Running go vet..."
	@go vet ./...
	@echo "✅ Go vet complete"

# Quick commands
all: clean build test lint
	@echo "🎉 All checks passed!"

.PHONY: help build test clean release docker-build docker-push install uninstall lint security-check
