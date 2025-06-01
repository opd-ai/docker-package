# Makefile

# This Makefile defines build targets and automation commands for the docker-package project.

.PHONY: all build clean test package

# Default target
all: build

# Build the static executable
build:
	go build -o bin/docker-package ./cmd/docker-package/main.go
	go build -o bin/builder ./cmd/builder/main.go

# Clean up build artifacts
clean:
	rm -rf bin/*

# Run tests
test:
	go test ./...

# Package the application for distribution
package: clean build
	./scripts/package.sh