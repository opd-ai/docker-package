package builder

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"os"
)

// Compress compresses the given data using gzip and returns the compressed data.
func Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)
	if _, err := writer.Write(data); err != nil {
		return nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Decompress decompresses the given gzip-compressed data and returns the original data.
func Decompress(compressedData []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(compressedData))
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	return ioutil.ReadAll(reader)
}

// SaveCompressedFile saves the compressed data to a file.
func SaveCompressedFile(filename string, data []byte) error {
	compressedData, err := Compress(data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, compressedData, os.ModePerm)
}

// LoadCompressedFile loads compressed data from a file and decompresses it.
func LoadCompressedFile(filename string) ([]byte, error) {
	compressedData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return Decompress(compressedData)
}