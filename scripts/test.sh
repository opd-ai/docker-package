#!/bin/bash

# This script runs the tests for the application.

set -e

# Run unit tests
echo "Running unit tests..."
go test ./... -v

# Run integration tests
echo "Running integration tests..."
go test ./test/integration/... -v

# Check for any test failures
if [ $? -ne 0 ]; then
    echo "Some tests failed. Please check the output above."
    exit 1
fi

echo "All tests passed successfully!"