package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// EnsureDir creates a directory if it does not exist.
func EnsureDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, os.ModePerm)
	}
	return nil
}

// ReadFile reads the contents of a file and returns it as a byte slice.
func ReadFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

// WriteFile writes data to a file, creating it if it does not exist.
func WriteFile(filePath string, data []byte) error {
	return ioutil.WriteFile(filePath, data, os.ModePerm)
}

// DeleteFile removes a file from the filesystem.
func DeleteFile(filePath string) error {
	return os.Remove(filePath)
}

// WalkDir traverses a directory and applies a function to each file.
func WalkDir(root string, fn func(string) error) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return fn(path)
		}
		return nil
	})
}