# API Documentation for Docker Package

## Overview

This document provides an overview of the API endpoints available in the Docker Package application. The API allows users to interact with the embedded Docker engine, manage containers, and perform various operations related to container lifecycle management.

## Base URL

The base URL for the API is:

```
http://localhost:8080/api
```

## Endpoints

### 1. Create Container

- **POST** `/containers/create`
- **Description**: Creates a new container based on the provided configuration.
- **Request Body**:
  ```json
  {
    "image": "string",
    "name": "string",
    "cmd": ["string"],
    "env": {
      "key": "value"
    },
    "ports": {
      "containerPort": "hostPort"
    }
  }
  ```
- **Response**:
  - **201 Created**: Returns the ID of the created container.
  - **400 Bad Request**: If the request body is invalid.

### 2. Start Container

- **POST** `/containers/{id}/start`
- **Description**: Starts a previously created container.
- **Path Parameters**:
  - `id`: The ID of the container to start.
- **Response**:
  - **204 No Content**: If the container starts successfully.
  - **404 Not Found**: If the container ID does not exist.

### 3. Stop Container

- **POST** `/containers/{id}/stop`
- **Description**: Stops a running container.
- **Path Parameters**:
  - `id`: The ID of the container to stop.
- **Response**:
  - **204 No Content**: If the container stops successfully.
  - **404 Not Found**: If the container ID does not exist.

### 4. Remove Container

- **DELETE** `/containers/{id}`
- **Description**: Removes a container by its ID.
- **Path Parameters**:
  - `id`: The ID of the container to remove.
- **Response**:
  - **204 No Content**: If the container is removed successfully.
  - **404 Not Found**: If the container ID does not exist.

### 5. List Containers

- **GET** `/containers`
- **Description**: Lists all containers, including their status.
- **Response**:
  - **200 OK**: Returns an array of container objects.
  ```json
  [
    {
      "id": "string",
      "name": "string",
      "status": "running|stopped"
    }
  ]
  ```

## Error Handling

All API responses include an error message in the following format:

```json
{
  "error": "string"
}
```

## Conclusion

This API provides a simple interface for managing containers in a rootless Docker environment. For further details on usage and examples, please refer to the usage documentation.