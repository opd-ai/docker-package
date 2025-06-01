package embedder

import (
	"archive/tar"
	"bytes"
	"io"
	"os"
	"path/filepath"
)

// Image represents an embedded container image.
type Image struct {
	Name string
	Data []byte
}

// LoadImages extracts embedded images from the binary and loads them into the local registry.
func LoadImages(images []Image) error {
	for _, img := range images {
		if err := extractImage(img); err != nil {
			return err
		}
	}
	return nil
}

// extractImage extracts a single image from the embedded data.
func extractImage(img Image) error {
	reader := bytes.NewReader(img.Data)
	tarReader := tar.NewReader(reader)

	// Create a directory for the image
	imageDir := filepath.Join("/var/lib/docker/images", img.Name)
	if err := os.MkdirAll(imageDir, 0755); err != nil {
		return err
	}

	// Extract files from the tar archive
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// Create the file or directory
		targetPath := filepath.Join(imageDir, header.Name)
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(targetPath, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			outFile, err := os.Create(targetPath)
			if err != nil {
				return err
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				outFile.Close()
				return err
			}
			outFile.Close()
		}
	}
	return nil
}