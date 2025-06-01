package embedder

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// ComposeConfig represents the structure of a docker-compose YAML configuration.
type ComposeConfig struct {
	Version  string                 `yaml:"version"`
	Services map[string]ServiceConfig `yaml:"services"`
}

// ServiceConfig represents the configuration for a single service in the compose file.
type ServiceConfig struct {
	Image   string   `yaml:"image"`
	Ports   []string `yaml:"ports,omitempty"`
	Volumes []string `yaml:"volumes,omitempty"`
}

// LoadComposeFile reads a docker-compose YAML file and unmarshals it into a ComposeConfig struct.
func LoadComposeFile(filePath string) (*ComposeConfig, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read compose file: %w", err)
	}

	var config ComposeConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal compose file: %w", err)
	}

	return &config, nil
}

// SaveComposeFile marshals a ComposeConfig struct and writes it to a YAML file.
func SaveComposeFile(filePath string, config *ComposeConfig) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal compose config: %w", err)
	}

	if err := ioutil.WriteFile(filePath, data, os.ModePerm); err != nil {
		return fmt.Errorf("failed to write compose file: %w", err)
	}

	return nil
}