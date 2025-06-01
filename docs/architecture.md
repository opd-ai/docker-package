# Architecture Overview of the Docker Package

## Introduction
The Docker Package is designed to provide a self-contained, rootless container deployment solution. It embeds a Docker engine, compose configuration, and container images into a single static executable, eliminating the need for a separate Docker installation on target systems. This document outlines the architecture and key components of the application.

## Architecture Components

### 1. Command Line Interface (CLI)
The application provides a command-line interface for users to interact with the embedded Docker engine. The CLI is implemented in the `cmd/docker-package/main.go` file, which initializes the rootless environment and manages command-line flags.

### 2. Builder
The builder component, located in `cmd/builder/main.go`, is responsible for creating the static executable. It handles the embedding of Docker images and compose configurations, ensuring that all necessary resources are included in the final binary.

### 3. Engine
The engine is the core of the application, managing the runtime environment for containers. It consists of several files:
- **runtime.go**: Initializes and configures the runtime directory for the embedded Docker daemon.
- **container.go**: Manages the lifecycle of containers, including launching, monitoring, and shutting down based on the docker-compose configuration.
- **image.go**: Handles the loading and management of embedded container images, including verification and multi-architecture support.
- **network.go**: Manages networking namespaces for container isolation, ensuring that each container operates within its own network environment.

### 4. Embedder
The embedder component is responsible for embedding resources into the executable:
- **embedder.go**: Provides functionality for embedding Docker images and compose files.
- **compose.go**: Parses and processes the docker-compose YAML configuration.
- **images.go**: Manages embedded image archives, including extraction and loading into the local registry.

### 5. Rootless Operation
To support rootless container execution, the application includes:
- **namespace.go**: Manages user namespaces for rootless operation.
- **user.go**: Handles user permissions and configurations necessary for running the application without elevated privileges.
- **permissions.go**: Checks and manages permissions required for container operations.

### 6. Builder Utilities
The builder utilities are responsible for packaging the application and its dependencies:
- **packager.go**: Packages the application into a single static binary.
- **bundler.go**: Handles the bundling of embedded resources and images.
- **compression.go**: Provides utilities for compressing and decompressing embedded images.

### 7. API Layer
The API layer allows interaction with the embedded Docker daemon:
- **client.go**: Defines the API client for managing container operations.
- **server.go**: Sets up the API server for handling requests and health checks.
- **types.go**: Defines data structures and types used in the API.

### 8. Utility Functions
Utility functions are provided for various operations:
- **fs.go**: Filesystem operations.
- **archive.go**: Handling tar archives.
- **network.go**: Network operations.

## Conclusion
The Docker Package architecture is designed for efficiency and ease of use, providing a complete container solution without the need for external dependencies. By embedding all necessary components into a single executable, it simplifies deployment and enhances security for rootless container operations.