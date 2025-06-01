package engine

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

// Container represents a container managed by the embedded Docker engine.
type Container struct {
	ID         string
	Image      string
	Command    []string
	Env        []string
	WorkingDir string
}

// NewContainer creates a new container instance.
func NewContainer(image string, command []string, env []string, workingDir string) *Container {
	return &Container{
		ID:         generateContainerID(),
		Image:      image,
		Command:    command,
		Env:        env,
		WorkingDir: workingDir,
	}
}

// Start launches the container.
func (c *Container) Start(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, "docker", "run", "--rm",
		"--name", c.ID,
		"-w", c.WorkingDir,
		c.Image)
	cmd.Env = c.Env

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to start container %s: %v, output: %s", c.ID, err, output)
	}

	return nil
}

// Stop stops the container.
func (c *Container) Stop() error {
	cmd := exec.Command("docker", "stop", c.ID)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to stop container %s: %v, output: %s", c.ID, err, output)
	}

	return nil
}

// generateContainerID generates a unique container ID.
func generateContainerID() string {
	return fmt.Sprintf("container-%d", os.Getpid())
}

// Cleanup removes the container resources.
func (c *Container) Cleanup() error {
	cmd := exec.Command("docker", "rm", c.ID)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to cleanup container %s: %v, output: %s", c.ID, err, output)
	}

	return nil
}
