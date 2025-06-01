package containerd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// Runtime manages the containerd runtime environment.
type Runtime struct {
	rootDir string
}

// NewRuntime creates a new Runtime instance with the specified root directory.
func NewRuntime(rootDir string) *Runtime {
	return &Runtime{rootDir: rootDir}
}

// Init initializes the containerd runtime environment.
func (r *Runtime) Init() error {
	if err := os.MkdirAll(r.rootDir, 0755); err != nil {
		return fmt.Errorf("failed to create root directory: %w", err)
	}

	// Initialize containerd
	cmd := exec.Command("containerd", "--root", r.rootDir)
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start containerd: %w", err)
	}

	return nil
}

// Shutdown gracefully stops the containerd runtime.
func (r *Runtime) Shutdown() error {
	cmd := exec.Command("pkill", "containerd")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to stop containerd: %w", err)
	}
	return nil
}

// GetRuntimeDir returns the root directory for the containerd runtime.
func (r *Runtime) GetRuntimeDir() string {
	return r.rootDir
}

// Example usage of the Runtime
func Example() {
	runtime := NewRuntime(filepath.Join(os.TempDir(), "containerd"))
	if err := runtime.Init(); err != nil {
		fmt.Println("Error initializing runtime:", err)
		return
	}
	defer runtime.Shutdown()

	fmt.Println("Containerd runtime initialized at:", runtime.GetRuntimeDir())
}
