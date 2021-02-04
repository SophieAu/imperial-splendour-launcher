package backend

import (
	"io/ioutil"
	"os"
)

// Handler .
type Handler interface {
	WriteFile(filePath string, data []byte, perm os.FileMode) error
	ReadFile(filePath string) ([]byte, error)
	MoveFile(source, destination string) error
}

// FileHandler is an abstraction of ioutil.WriteFile and ioutil.ReadFile
type FileHandler struct {
}

// WriteFile implements the Handler interface that's been created so that ioutil.WriteFile can be mocked
func (w *FileHandler) WriteFile(filePath string, data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(filePath, data, perm)
}

// ReadFile implements the Handler interface that's been created so that ioutil.ReadFile can be mocked
func (w *FileHandler) ReadFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

// MoveFile implements the Handler interface that's been created so that os.MoveFile can be mocked
func (w *FileHandler) MoveFile(source, destination string) error {
	return os.Rename(source, destination)
}
