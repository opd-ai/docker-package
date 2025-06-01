package utils

import (
	"net"
	"strings"
)

// GetAvailablePort finds an available port on the local machine.
func GetAvailablePort() (int, error) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}
	defer listener.Close()
	return listener.Addr().(*net.TCPAddr).Port, nil
}

// IsPrivateIP checks if the given IP address is a private IP address.
func IsPrivateIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}
	return parsedIP.IsPrivate()
}

// ParseCIDR parses a CIDR notation string and returns the IP network.
func ParseCIDR(cidr string) (*net.IPNet, error) {
	_, network, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}
	return network, nil
}

// GetHostIP retrieves the host's IP address.
func GetHostIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String(), nil
		}
	}
	return "", nil
}

// SplitHostPort splits a host:port string into host and port components.
func SplitHostPort(hostPort string) (string, string, error) {
	host, port, err := net.SplitHostPort(hostPort)
	if err != nil {
		return "", "", err
	}
	return strings.TrimSpace(host), strings.TrimSpace(port), nil
}