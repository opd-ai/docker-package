package api

// Container represents a container's configuration and state.
type Container struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Image   string `json:"image"`
	Status  string `json:"status"`
	Ports   []Port `json:"ports"`
	Created int64  `json:"created"`
}

// Port represents a network port configuration for a container.
type Port struct {
	HostPort      int    `json:"host_port"`
	ContainerPort int    `json:"container_port"`
	Protocol      string `json:"protocol"`
}

// Image represents a container image's metadata.
type Image struct {
	ID          string `json:"id"`
	Repository  string `json:"repository"`
	Tag         string `json:"tag"`
	Size        int64  `json:"size"`
	Created     int64  `json:"created"`
	Architecture string `json:"architecture"`
}

// Network represents a network configuration for containers.
type Network struct {
	Name   string `json:"name"`
	Driver string `json:"driver"`
}

// HealthCheck represents the health check configuration for a container.
type HealthCheck struct {
	Interval int `json:"interval"`
	Timeout  int `json:"timeout"`
	Retries  int `json:"retries"`
}