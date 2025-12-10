.PHONY: build install test testacc clean fmt lint

GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
BINARY_NAME=terraform-provider-mailgun
VERSION?=0.9.0
OS_ARCH?=linux_amd64

default: build

build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME)

install: build
	@echo "Installing $(BINARY_NAME) locally..."
	@mkdir -p ~/.terraform.d/plugins/registry.terraform.io/murad-heydarov/mailgun/$(VERSION)/$(OS_ARCH)
	@cp $(BINARY_NAME) ~/.terraform.d/plugins/registry.terraform.io/murad-heydarov/mailgun/$(VERSION)/$(OS_ARCH)/

test:
	@echo "Running tests..."
	@go test -v ./... -timeout 120s

testacc:
	@echo "Running acceptance tests..."
	@TF_ACC=1 go test -v ./... -timeout 120m

fmt:
	@echo "Formatting Go code..."
	@gofmt -w $(GOFMT_FILES)

lint:
	@echo "Running linter..."
	@golangci-lint run

clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)
	@go clean

deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build     - Build the provider binary"
	@echo "  install   - Install provider locally for testing"
	@echo "  test      - Run unit tests"
	@echo "  testacc   - Run acceptance tests (requires MAILGUN_API_KEY)"
	@echo "  fmt       - Format Go code"
	@echo "  lint      - Run linter"
	@echo "  clean     - Clean build artifacts"
	@echo "  deps      - Install dependencies"
