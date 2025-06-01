FROM golang:1.20 AS builder

# Set the working directory
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the entire application source code
COPY . .

# Build the static executable
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o docker-package ./cmd/docker-package/main.go

# Create a minimal image for the final executable
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the static executable from the builder stage
COPY --from=builder /app/docker-package .

# Set the entry point for the container
ENTRYPOINT ["./docker-package"]