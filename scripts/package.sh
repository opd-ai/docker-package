#!/bin/bash

# This script packages the application for distribution as a static executable.
# It builds the Go application, embeds necessary resources, and prepares the final binary.

set -e

# Define variables
APP_NAME="docker-package"
BUILD_DIR="build"
OUTPUT_DIR="$BUILD_DIR/output"
VERSION=$(git describe --tags --abbrev=0)

# Create output directory
mkdir -p $OUTPUT_DIR

# Build the application
echo "Building the application..."
GOOS=linux GOARCH=amd64 go build -o $OUTPUT_DIR/$APP_NAME-$VERSION .

# Package the application with embedded resources
echo "Packaging the application..."
# Here you would include commands to embed Docker images and compose files
# For example, using the embedder package to handle this

# Compress the output binary
echo "Compressing the output..."
tar -czf $OUTPUT_DIR/$APP_NAME-$VERSION.tar.gz -C $OUTPUT_DIR $APP_NAME-$VERSION

echo "Packaging complete. Output located in $OUTPUT_DIR."