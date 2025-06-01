# Docker Package

This project provides a self-contained Go application that embeds a Docker engine, compose configuration, and container images into a single static executable. It allows for rootless container deployment without requiring Docker installation on target systems.

## Features

- **Rootless Operation**: Run containers without requiring root privileges.
- **Embedded Docker Engine**: Includes a lightweight Docker engine embedded within the executable.
- **Docker Compose Support**: Supports Docker Compose configurations for multi-container applications.
- **Static Binary**: Compiles into a single static binary for easy distribution and deployment.

## Installation

To build the application, run the following command:

```bash
make build
```

This will create a static executable in the `bin` directory.

## Usage

To run the application, use the following command:

```bash
./bin/docker-package
```

You can specify a Docker Compose file using the `-f` flag:

```bash
./bin/docker-package -f examples/simple-web/docker-compose.yml
```

## Examples

- **Simple Web Application**: See the `examples/simple-web` directory for a sample Docker Compose configuration and its corresponding settings.
- **Microservices**: Check the `examples/microservices` directory for a more complex setup involving multiple services.

## Configuration

Configuration files are located in the `configs` directory. You can customize the settings in `default.yaml` or `rootless.yaml` as needed.

## Testing

To run the tests, execute:

```bash
make test
```

This will run all unit and integration tests defined in the `test` directory.

## Documentation

For more detailed information, refer to the documentation in the `docs` directory:

- [Architecture](docs/architecture.md)
- [Usage](docs/usage.md)
- [API Documentation](docs/api.md)

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for details.