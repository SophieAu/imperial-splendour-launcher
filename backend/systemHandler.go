package backend

import (
	"io/ioutil"
	"os"
)

// Handler is an abstraction of all direct system interaction
type Handler interface {
	WriteFile(filePath string, data []byte) error
	ReadFile(filePath string) ([]byte, error)
	MoveFile(source, destination string) error
	Executable() (string, error)
	Getenv(key string) string
}

type SystemHandler struct {
}

func (w *SystemHandler) WriteFile(filePath string, data []byte) error {
	return ioutil.WriteFile(filePath, data, os.FileMode(0644))
}

func (w *SystemHandler) ReadFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

func (w *SystemHandler) MoveFile(source, destination string) error {
	return os.Rename(source, destination)
}

func (w *SystemHandler) Executable() (string, error) {
	return os.Executable()
}

func (w *SystemHandler) Getenv(key string) string {
	return os.Getenv(key)
}
