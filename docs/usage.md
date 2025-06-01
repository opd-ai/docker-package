# Usage Instructions for Docker Package

## Overview

The Docker Package is a self-contained Go application that allows you to run Docker containers in a rootless environment without requiring a separate Docker installation. This document provides instructions on how to use the application, including setup, commands, and examples.

## Installation

To install the Docker Package, you need to build the static executable from the source code. Follow these steps:

1. Clone the repository:
   ```
   git clone https://github.com/opd-ai/docker-package.git
   cd docker-package
   ```

2. Build the application:
   ```
   ./scripts/build.sh
   ```

3. The static executable will be located in the `bin` directory.

## Running the Application

To run the Docker Package, use the following command:

```
./bin/docker-package [OPTIONS]
```

### Options

- `--config`: Specify the path to the configuration file (default: `configs/default.yaml`).
- `--compose-file`: Specify the path to the Docker Compose file (default: `docker-compose.yml` in the current directory).
- `--help`: Display help information.

## Example Usage

### Running a Simple Web Application

1. Navigate to the example directory:
   ```
   cd examples/simple-web
   ```

2. Run the application using the provided Docker Compose file:
   ```
   ../bin/docker-package --compose-file docker-compose.yml
   ```

3. Access the application in your web browser at `http://localhost:8080`.

### Running a Microservices Application

1. Navigate to the microservices example directory:
   ```
   cd examples/microservices
   ```

2. Start the microservices using the Docker Compose file:
   ```
   ../bin/docker-package --compose-file docker-compose.yml
   ```

3. Access the services as defined in the Compose file.

## Troubleshooting

- Ensure that you have the necessary permissions to run the application in a rootless environment.
- Check the logs for any errors during container startup. Logs can be found in the `logs` directory.

## Conclusion

The Docker Package provides a convenient way to run Docker containers without the overhead of a traditional Docker installation. For further information, refer to the [API documentation](api.md) and [architecture documentation](architecture.md).