package rootless

import (
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/sys/unix"
)

// Namespace represents a user namespace for rootless container operation.
type Namespace struct {
	uid  int
	gid  int
	path string
}

// NewNamespace creates a new user namespace for rootless operation.
func NewNamespace() (*Namespace, error) {
	uid, gid := os.Getuid(), os.Getgid()
	nsPath := fmt.Sprintf("/proc/self/ns/user")

	// Create the user namespace using unshare command
	cmd := exec.Command("unshare", "--user", "--map-root-user", "--mount-proc", "--fork", "bash")
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to create user namespace: %w", err)
	}

	return &Namespace{
		uid:  uid,
		gid:  gid,
		path: nsPath,
	}, nil
}

// Enter enters the user namespace.
func (n *Namespace) Enter() error {
	file, err := os.Open(n.path)
	if err != nil {
		return fmt.Errorf("failed to open namespace file: %w", err)
	}
	defer file.Close()

	if err := unix.Setns(int(file.Fd()), unix.CLONE_NEWUSER); err != nil {
		return fmt.Errorf("failed to enter user namespace: %w", err)
	}
	return nil
}

// Exit cleans up the namespace when done.
func (n *Namespace) Exit() error {
	// Cleanup logic can be added here if necessary
	return nil
}
