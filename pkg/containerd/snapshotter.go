package containerd

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/snapshots"
	"github.com/pkg/errors"
)

// Snapshotter is a struct that manages snapshotting of container filesystems.
type Snapshotter struct {
	client   *containerd.Client
	snapshot string
}

// NewSnapshotter creates a new Snapshotter instance.
func NewSnapshotter(ctx context.Context, address string, snapshotterName string) (*Snapshotter, error) {
	client, err := containerd.New(address)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to containerd")
	}

	return &Snapshotter{
		client:   client,
		snapshot: snapshotterName,
	}, nil
}

// CreateSnapshot creates a new snapshot of the specified container filesystem.
func (s *Snapshotter) CreateSnapshot(ctx context.Context, snapshotID string, rootfsPath string) error {
	snapshotter := s.client.SnapshotService(s.snapshot)

	// Read the root filesystem to be snapshotted
	rootfs, err := ioutil.ReadDir(rootfsPath)
	if err != nil {
		return errors.Wrap(err, "failed to read root filesystem")
	}

	// Create a new snapshot
	_, err = snapshotter.Prepare(ctx, snapshotID, snapshots.WithLabels(map[string]string{"created": "true"}))
	if err != nil {
		return errors.Wrap(err, "failed to create snapshot")
	}

	// Here you would typically copy the root filesystem to the snapshot
	// This is a placeholder for the actual snapshotting logic
	fmt.Printf("Snapshot created for %s with ID %s\n", rootfs, snapshotID)

	return nil
}

// RemoveSnapshot removes the specified snapshot.
func (s *Snapshotter) RemoveSnapshot(ctx context.Context, snapshotID string) error {
	snapshotter := s.client.SnapshotService(s.snapshot)

	if err := snapshotter.Remove(ctx, snapshotID); err != nil {
		return errors.Wrap(err, "failed to remove snapshot")
	}

	fmt.Printf("Snapshot %s removed\n", snapshotID)
	return nil
}

// Cleanup cleans up any resources used by the Snapshotter.
func (s *Snapshotter) Cleanup() error {
	if err := s.client.Close(); err != nil {
		return errors.Wrap(err, "failed to close containerd client")
	}
	return nil
}
