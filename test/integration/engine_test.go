package integration_test

import (
    "os"
    "os/exec"
    "testing"
)

func TestEngineIntegration(t *testing.T) {
    // Set up the environment for the test
    if err := os.Setenv("DOCKER_PACKAGE_TEST", "1"); err != nil {
        t.Fatalf("Failed to set environment variable: %v", err)
    }

    // Run the embedded Docker engine
    cmd := exec.Command("./docker-package", "run", "--rm", "hello-world")
    output, err := cmd.CombinedOutput()
    if err != nil {
        t.Fatalf("Failed to run embedded Docker engine: %v\nOutput: %s", err, output)
    }

    // Check if the output contains the expected message
    expectedOutput := "Hello from Docker!"
    if !contains(string(output), expectedOutput) {
        t.Errorf("Expected output to contain %q, got %q", expectedOutput, string(output))
    }
}

// contains checks if a string contains a substring
func contains(s, substr string) bool {
    return len(s) >= len(substr) && s[:len(substr)] == substr
}