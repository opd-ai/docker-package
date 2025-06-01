package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/opd-ai/docker-package/internal/engine"
	"github.com/opd-ai/docker-package/internal/rootless"
)

func main() {
	// Initialize command-line flags
	configPath := flag.String("config", "configs/default.yaml", "Path to the configuration file")
	flag.Parse()

	// Set up rootless environment
	if err := rootless.Setup(); err != nil {
		log.Fatalf("Failed to set up rootless environment: %v", err)
	}

	// Initialize the Docker engine
	dockerEngine, err := engine.NewEngine(*configPath)
	if err != nil {
		log.Fatalf("Failed to initialize Docker engine: %v", err)
	}

	// Start the Docker engine
	if err := dockerEngine.Start(); err != nil {
		log.Fatalf("Failed to start Docker engine: %v", err)
	}

	fmt.Println("Docker engine started successfully in rootless mode.")

	// Wait for termination signal
	// (Implementation of signal handling would go here)

	// Clean up and shutdown
	if err := dockerEngine.Stop(); err != nil {
		log.Fatalf("Failed to stop Docker engine: %v", err)
	}

	fmt.Println("Docker engine stopped.")
}
