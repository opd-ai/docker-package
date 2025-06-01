package integration

import (
    "os"
    "os/exec"
    "testing"
)

func TestRootlessContainer(t *testing.T) {
    // Set up a temporary directory for the rootless environment
    tempDir, err := os.MkdirTemp("", "rootless-test")
    if err != nil {
        t.Fatalf("failed to create temp directory: %v", err)
    }
    defer os.RemoveAll(tempDir)

    // Set environment variables for rootless operation
    os.Setenv("XDG_RUNTIME_DIR", tempDir)
    os.Setenv("DOCKER_HOST", "unix://" + tempDir + "/docker.sock")

    // Start the rootless Docker daemon
    cmd := exec.Command("dockerd-rootless.sh", "--data-root", tempDir)
    if err := cmd.Start(); err != nil {
        t.Fatalf("failed to start rootless Docker daemon: %v", err)
    }
    defer cmd.Process.Kill()

    // Wait for the daemon to start
    if err := cmd.Wait(); err != nil {
        t.Fatalf("Docker daemon exited with error: %v", err)
    }

    // Here you would add tests for container creation, execution, etc.
    // For example, running a simple container
    runCmd := exec.Command("docker", "run", "--rm", "hello-world")
    output, err := runCmd.CombinedOutput()
    if err != nil {
        t.Fatalf("failed to run container: %v, output: %s", err, output)
    }

    // Check the output for expected results
    expectedOutput := "Hello from Docker!"
    if !contains(string(output), expectedOutput) {
        t.Errorf("unexpected output: got %s, want %s", output, expectedOutput)
    }
}

func contains(s, substr string) bool {
    return len(s) >= len(substr) && s[:len(substr)] == substr
}