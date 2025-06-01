package main

import (
	"fmt"
	"os"

	"github.com/opd-ai/docker-package/internal/builder"
	"github.com/opd-ai/docker-package/internal/embedder"
)

func main() {
	// Initialize the embedder to load Docker images and compose configurations
	if err := embedder.LoadResources(); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading resources: %v\n", err)
		os.Exit(1)
	}

	// Build the static executable with embedded resources
	if err := builder.Build(); err != nil {
		fmt.Fprintf(os.Stderr, "Error building the executable: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Static executable built successfully.")
}
