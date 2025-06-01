package builder

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Bundler is responsible for bundling embedded resources and images into the executable.
type Bundler struct {
	// Add fields as necessary for bundling resources
}

// NewBundler creates a new instance of Bundler.
func NewBundler() *Bundler {
	return &Bundler{}
}

// BundleResources bundles the specified resources into a single tar.gz file.
func (b *Bundler) BundleResources(resourcePaths []string, outputPath string) error {
	// Create a buffer to write our tar file to
	var buf bytes.Buffer
	gzw := gzip.NewWriter(&buf)
	tw := tar.NewWriter(gzw)

	for _, resourcePath := range resourcePaths {
		if err := b.addFileToTar(tw, resourcePath); err != nil {
			return err
		}
	}

	if err := tw.Close(); err != nil {
		return err
	}
	if err := gzw.Close(); err != nil {
		return err
	}

	return ioutil.WriteFile(outputPath, buf.Bytes(), 0644)
}

// addFileToTar adds a file to the tar writer.
func (b *Bundler) addFileToTar(tw *tar.Writer, filePath string) error {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	header, err := tar.FileInfoHeader(fileInfo, "")
	if err != nil {
		return err
	}
	header.Name = filepath.Base(filePath)

	if err := tw.WriteHeader(header); err != nil {
		return err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := io.Copy(tw, file); err != nil {
		return err
	}

	return nil
}
