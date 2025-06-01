#!/bin/bash

# This script automates the build process for the docker-package application.

set -e

# Define variables
APP_NAME="docker-package"
BUILD_DIR="build"
OUTPUT_DIR="$BUILD_DIR/output"
VERSION="1.0.0"

# Create output directory
mkdir -p $OUTPUT_DIR

# Build the application
echo "Building the $APP_NAME application..."
go build -o $OUTPUT_DIR/$APP_NAME ./cmd/docker-package

# Check if the build was successful
if [ $? -ne 0 ]; then
    echo "Build failed!"
    exit 1
fi

# Package the application
echo "Packaging the application..."
tar -czf $OUTPUT_DIR/$APP_NAME-$VERSION.tar.gz -C $OUTPUT_DIR $APP_NAME

# Check if the packaging was successful
if [ $? -ne 0 ]; then
    echo "Packaging failed!"
    exit 1
fi

echo "Build and packaging completed successfully!"
echo "Output: $OUTPUT_DIR/$APP_NAME-$VERSION.tar.gz"