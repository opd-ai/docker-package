package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Client is a struct that represents the API client for interacting with the embedded Docker daemon.
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new API client with the specified base URL.
func NewClient(baseURL string) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

// Container represents a container in the API.
type Container struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// ListContainers retrieves a list of containers from the embedded Docker daemon.
func (c *Client) ListContainers() ([]Container, error) {
	resp, err := c.httpClient.Get(fmt.Sprintf("%s/containers", c.baseURL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to list containers: %s", resp.Status)
	}

	var containers []Container
	if err := json.NewDecoder(resp.Body).Decode(&containers); err != nil {
		return nil, err
	}

	return containers, nil
}

// StartContainer starts a container with the specified ID.
func (c *Client) StartContainer(containerID string) error {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/containers/%s/start", c.baseURL, containerID), nil)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		var buf bytes.Buffer
		buf.ReadFrom(resp.Body)
		return fmt.Errorf("failed to start container: %s", buf.String())
	}

	return nil
}

// StopContainer stops a container with the specified ID.
func (c *Client) StopContainer(containerID string) error {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/containers/%s/stop", c.baseURL, containerID), nil)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		var buf bytes.Buffer
		buf.ReadFrom(resp.Body)
		return fmt.Errorf("failed to stop container: %s", buf.String())
	}

	return nil
}