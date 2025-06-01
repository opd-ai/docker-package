package engine

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Image represents a container image with its metadata.
type Image struct {
	Name    string
	Version string
	Path    string
}

// LoadImage loads an embedded container image from the specified path.
func LoadImage(imagePath string) (*Image, error) {
	if imagePath == "" {
		return nil, errors.New("image path cannot be empty")
	}

	// Check if the image file exists
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("image file does not exist: %s", imagePath)
	}

	// Read the image file
	data, err := os.ReadFile(imagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read image file: %w", err)
	}

	// Validate minimum file size (empty files are invalid)
	if len(data) == 0 {
		return nil, fmt.Errorf("image file is empty: %s", imagePath)
	}

	// Extract image name from filename (remove extension)
	fileName := filepath.Base(imagePath)
	imageName := fileName
	if ext := filepath.Ext(fileName); ext != "" {
		imageName = fileName[:len(fileName)-len(ext)]
	}

	// Parse version from filename if present (format: name:version or name-version)
	version := "latest"
	if idx := findVersionSeparator(imageName); idx != -1 {
		version = imageName[idx+1:]
		imageName = imageName[:idx]
	}

	// Validate extracted name
	if imageName == "" {
		return nil, fmt.Errorf("could not extract valid image name from path: %s", imagePath)
	}

	// Create image object with extracted metadata
	image := &Image{
		Name:    imageName,
		Version: version,
		Path:    imagePath,
	}

	return image, nil
}

// findVersionSeparator finds the index of version separator (':' or '-') in image name
func findVersionSeparator(name string) int {
	// Look for ':' first (Docker convention)
	if idx := len(name) - 1; idx >= 0 {
		for i := idx; i >= 0; i-- {
			if name[i] == ':' {
				return i
			}
		}
	}

	// Fallback to '-' separator
	if idx := len(name) - 1; idx >= 0 {
		for i := idx; i >= 0; i-- {
			if name[i] == '-' {
				return i
			}
		}
	}

	return -1
}

// VerifyImage checks if the image is valid and can be used.
func VerifyImage(image *Image) error {
	if image == nil {
		return errors.New("image cannot be nil")
	}
	if image.Name == "" || image.Version == "" {
		return errors.New("image name and version must be specified")
	}
	// Additional verification logic can be added here
	return nil
}

// SaveImage saves the image to the specified destination.
func SaveImage(image *Image, destPath string) error {
	if image == nil {
		return errors.New("image cannot be nil")
	}
	if destPath == "" {
		return errors.New("destination path cannot be empty")
	}

	// Simulate saving the image (in a real implementation, you would write the image data)
	err := ioutil.WriteFile(destPath, []byte("dummy image data"), 0644)
	if err != nil {
		return fmt.Errorf("failed to save image: %w", err)
	}

	return nil
}

// ListImages lists all available images in the specified directory.
func ListImages(dir string) ([]*Image, error) {
	var images []*Image

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			images = append(images, &Image{Name: info.Name(), Path: path})
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to list images: %w", err)
	}

	return images, nil
}
