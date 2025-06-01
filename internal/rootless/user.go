package rootless

import (
	"os"
	"os/user"
)

// User represents the user information for rootless container operations.
type User struct {
	Username string
	UID      string
	GID      string
}

// NewUser initializes a new User instance with the current user's information.
func NewUser() (*User, error) {
	currentUser, err := user.Current()
	if err != nil {
		return nil, err
	}

	return &User{
		Username: currentUser.Username,
		UID:      currentUser.Uid,
		GID:      currentUser.Gid,
	}, nil
}

// HasPermission checks if the user has the necessary permissions for container operations.
func (u *User) HasPermission() bool {
	// Implement permission checks based on user UID and GID.
	// This is a placeholder for actual permission logic.
	return true
}

// GetHomeDir returns the home directory of the user.
func (u *User) GetHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return homeDir
}