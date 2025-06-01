package embedder

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
)

// Embedder is responsible for embedding Docker images and compose files into the executable.
type Embedder struct {
	Images      map[string][]byte
	ComposeFile []byte
}

// NewEmbedder initializes a new Embedder instance.
func NewEmbedder() *Embedder {
	return &Embedder{
		Images: make(map[string][]byte),
	}
}

// LoadComposeFile loads the docker-compose file from the specified path.
func (e *Embedder) LoadComposeFile(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	e.ComposeFile = data
	return nil
}

// AddImage adds an image to the embedder.
func (e *Embedder) AddImage(name string, imagePath string) error {
	data, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return err
	}
	e.Images[name] = data
	return nil
}

// SaveEmbeddedResources saves the embedded resources to a specified output path.
func (e *Embedder) SaveEmbeddedResources(outputPath string) error {
	var buffer bytes.Buffer

	// Serialize the compose file
	if err := json.NewEncoder(&buffer).Encode(e.ComposeFile); err != nil {
		return err
	}

	// Serialize the images
	for name, image := range e.Images {
		if err := json.NewEncoder(&buffer).Encode(map[string]interface{}{
			"name":  name,
			"image": image,
		}); err != nil {
			return err
		}
	}

	// Compress the buffer
	compressedData, err := compressData(buffer.Bytes())
	if err != nil {
		return err
	}

	// Write to output file
	return ioutil.WriteFile(outputPath, compressedData, 0644)
}

// compressData compresses the given data using gzip.
func compressData(data []byte) ([]byte, error) {
	var buffer bytes.Buffer
	writer := gzip.NewWriter(&buffer)
	if _, err := writer.Write(data); err != nil {
		return nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// ExtractResources extracts the embedded resources from the specified path.
func ExtractResources(path string) (*Embedder, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Decompress the data
	decompressedData, err := decompressData(data)
	if err != nil {
		return nil, err
	}

	embedder := NewEmbedder()
	if err := json.Unmarshal(decompressedData, embedder); err != nil {
		return nil, err
	}

	return embedder, nil
}

// decompressData decompresses the given gzip data.
func decompressData(data []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	return ioutil.ReadAll(reader)
}
