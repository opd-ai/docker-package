package containerd

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/errdefs"
	"github.com/containerd/containerd/images"
	"github.com/containerd/containerd/namespaces"
)

// Client wraps the containerd client for managing containers.
type Client struct {
	ctrClient *containerd.Client
	namespace string
}

// NewClient creates a new containerd client in the specified namespace.
func NewClient(namespace string) (*Client, error) {
	// Connect to the containerd daemon
	ctrClient, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to containerd: %w", err)
	}

	return &Client{
		ctrClient: ctrClient,
		namespace: namespace,
	}, nil
}

// Close closes the containerd client connection.
func (c *Client) Close() error {
	return c.ctrClient.Close()
}

// CreateContainer creates a new container with the specified image.
func (c *Client) CreateContainer(ctx context.Context, name, imageRef string) error {
	ctx = namespaces.WithNamespace(ctx, c.namespace)

	// Pull the image
	image, err := c.ctrClient.Pull(ctx, imageRef)
	if err != nil {
		return fmt.Errorf("failed to pull image: %w", err)
	}

	// Create a new container
	_, err = c.ctrClient.NewContainer(ctx, name, containerd.WithImage(image))
	if err != nil {
		return fmt.Errorf("failed to create container: %w", err)
	}

	return nil
}

// StartContainer starts the specified container.
func (c *Client) StartContainer(ctx context.Context, name string) error {
	ctx = namespaces.WithNamespace(ctx, c.namespace)

	// Get the container
	container, err := c.ctrClient.LoadContainer(ctx, name)
	if err != nil {
		return fmt.Errorf("failed to load container: %w", err)
	}

	// Start the container
	task, err := container.NewTask(ctx, containerd.NewIO(os.Stdin, os.Stdout, os.Stderr))
	if err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}

	if err := task.Start(ctx); err != nil {
		return fmt.Errorf("failed to start task: %w", err)
	}

	return nil
}

// StopContainer stops the specified container.
func (c *Client) StopContainer(ctx context.Context, name string) error {
	ctx = namespaces.WithNamespace(ctx, c.namespace)

	// Get the container
	container, err := c.ctrClient.LoadContainer(ctx, name)
	if err != nil {
		return fmt.Errorf("failed to load container: %w", err)
	}

	// Stop the container
	task, err := container.Task(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to get task: %w", err)
	}

	if err := task.Kill(ctx, 0); err != nil && !errdefs.IsNotFound(err) {
		return fmt.Errorf("failed to kill task: %w", err)
	}

	return nil
}

// ListImages lists all images in the containerd namespace.
func (c *Client) ListImages(ctx context.Context) ([]images.Image, error) {
	ctx = namespaces.WithNamespace(ctx, c.namespace)

	// List images
	imageList, err := c.ctrClient.ImageService().List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list images: %w", err)
	}

	return imageList, nil
}

// SaveImage saves the specified image to a tar file.
func (c *Client) SaveImage(ctx context.Context, imageRef, outputPath string) error {
	ctx = namespaces.WithNamespace(ctx, c.namespace)

	// Get the image
	image, err := c.ctrClient.GetImage(ctx, imageRef)
	if err != nil {
		return fmt.Errorf("failed to get image: %w", err)
	}

	// Save the image to a tar file
	imageTar, err := image.Export(ctx)
	if err != nil {
		return fmt.Errorf("failed to export image: %w", err)
	}
	defer imageTar.Close()

	data, err := ioutil.ReadAll(imageTar)
	if err != nil {
		return fmt.Errorf("failed to read image tar: %w", err)
	}

	if err := ioutil.WriteFile(outputPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write image tar to file: %w", err)
	}

	return nil
}
