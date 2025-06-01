package rootless

import (
	"errors"
	"os"
	"syscall"
)

// CheckPermissions verifies if the current user has the necessary permissions
// to perform container operations in a rootless environment.
func CheckPermissions() error {
	// Check if the user is running as root
	if os.Geteuid() == 0 {
		return errors.New("running as root is not allowed in rootless mode")
	}

	// Check if the user has the necessary capabilities
	capabilities := []string{"CAP_SYS_ADMIN", "CAP_NET_ADMIN"}
	for _, cap := range capabilities {
		if !hasCapability(cap) {
			return errors.New("missing required capability: " + cap)
		}
	}

	return nil
}

// hasCapability checks if the current user has the specified capability.
func hasCapability(cap string) bool {
	// This is a placeholder for actual capability checking logic.
	// In a real implementation, you would use a library or system call
	// to check the user's capabilities.
	return true
}

// SetUserNamespace sets up the user namespace for rootless container operations.
func SetUserNamespace() error {
	// Create a new user namespace
	if err := syscall.Unshare(syscall.CLONE_NEWUSER); err != nil {
		return err
	}

	// Additional setup for user namespace can be added here

	return nil
}