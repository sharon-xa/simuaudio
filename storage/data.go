package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
)

// The package includes:
//
// Thread-safe operations using sync.RWMutex
// Error handling for file operations
// Simple interface with Save, Load, and Clear methods
// Support for any struct that can be marshaled to JSON

type FileStorage struct {
	filepath string
	mu       sync.RWMutex
}

// NewFileStorage creates a new FileStorage instance
// appName is your application name, filename is the name of the JSON file
func NewFileStorage(appName, filename string) (*FileStorage, error) {
	dataDir := os.Getenv("XDG_DATA_HOME")
	if dataDir == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		dataDir = filepath.Join(homeDir, ".local", "share")
	}

	appDir := filepath.Join(dataDir, appName)

	if err := os.MkdirAll(appDir, 0755); err != nil {
		return nil, err
	}

	return &FileStorage{
		filepath: filepath.Join(appDir, filename),
	}, nil
}

// Save marshals the data to JSON and writes it to the file
func (fs *FileStorage) Save(data interface{}) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(fs.filepath, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Load reads JSON from the file and unmarshals it into the provided data structure
func (fs *FileStorage) Load(data interface{}) error {
	fs.mu.RLock()
	defer fs.mu.RUnlock()

	if _, err := os.Stat(fs.filepath); os.IsNotExist(err) {
		return err
	}

	jsonData, err := os.ReadFile(fs.filepath)
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonData, &data)
}

// Clear empties the file by writing an empty byte slice
func (fs *FileStorage) Clear() error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	return os.WriteFile(fs.filepath, []byte{}, 0644)
}
