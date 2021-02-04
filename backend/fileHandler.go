package backend

import (
	"io/ioutil"
	"os"
)

// Handler .
type Handler interface {
	WriteFile(filePath string, data []byte) error
	ReadFile(filePath string) ([]byte, error)
	MoveFile(source, destination string) error
	Executable() (string, error)
	Getenv(key string) string
}

// SystemHandler is an abstraction of all direct system interaction
type SystemHandler struct {
}

// WriteFile implements the Handler interface that's been created so that ioutil.WriteFile can be mocked
func (w *SystemHandler) WriteFile(filePath string, data []byte) error {
	return ioutil.WriteFile(filePath, data, os.FileMode(0644))
}

// ReadFile implements the Handler interface that's been created so that ioutil.ReadFile can be mocked
func (w *SystemHandler) ReadFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

// MoveFile implements the Handler interface that's been created so that os.MoveFile can be mocked
func (w *SystemHandler) MoveFile(source, destination string) error {
	return os.Rename(source, destination)
}

// Executable implements the Handler interface that's been created so that ios.Executable can be mocked
func (w *SystemHandler) Executable() (string, error) {
	return os.Executable()
}

// Getenv implements the Handler interface that's been created so that os.Getenv can be mocked
func (w *SystemHandler) Getenv(key string) string {
	return os.Getenv(key)
}
