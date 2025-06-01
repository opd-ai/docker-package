package builder

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Packager is responsible for packaging the application and its dependencies into a single static binary.
type Packager struct {
	// Add any necessary fields for the Packager
}

// NewPackager creates a new instance of Packager.
func NewPackager() *Packager {
	return &Packager{}
}

// Package creates a tar.gz archive of the specified files and directories.
func (p *Packager) Package(output string, sources []string) error {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	tw := tar.NewWriter(gz)

	for _, source := range sources {
		if err := p.addSourceToTar(tw, source); err != nil {
			return err
		}
	}

	if err := tw.Close(); err != nil {
		return err
	}
	if err := gz.Close(); err != nil {
		return err
	}

	return ioutil.WriteFile(output, buf.Bytes(), 0644)
}

// addSourceToTar adds a file or directory to the tar writer.
func (p *Packager) addSourceToTar(tw *tar.Writer, source string) error {
	info, err := os.Stat(source)
	if err != nil {
		return fmt.Errorf("failed to stat %s: %w", source, err)
	}

	if info.IsDir() {
		return p.addDirectoryToTar(tw, source)
	}
	return p.addFileToTar(tw, source)
}

// addDirectoryToTar adds a directory and its contents to the tar writer.
func (p *Packager) addDirectoryToTar(tw *tar.Writer, dir string) error {
	err := filepath.Walk(dir, func(file string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		header.Name, _ = filepath.Rel(filepath.Dir(dir), file)

		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		if !info.IsDir() {
			data, err := ioutil.ReadFile(file)
			if err != nil {
				return err
			}
			if _, err := tw.Write(data); err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

// addFileToTar adds a single file to the tar writer.
func (p *Packager) addFileToTar(tw *tar.Writer, file string) error {
	info, err := os.Stat(file)
	if err != nil {
		return fmt.Errorf("failed to stat file %s: %w", file, err)
	}

	header, err := tar.FileInfoHeader(info, "")
	if err != nil {
		return err
	}
	header.Name = filepath.Base(file)

	if err := tw.WriteHeader(header); err != nil {
		return err
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	_, err = tw.Write(data)
	return err
}