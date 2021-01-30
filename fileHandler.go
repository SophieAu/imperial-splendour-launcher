package main

import (
	"io/ioutil"
	"os"
	"path"
)

func getTodosterDir() string {
	todosterDir := os.Getenv("TODOSTER_DIR")
	if todosterDir == "" {
		panic("TODOSTER_DIR env var not set")
	}

	return path.Clean(todosterDir)
}

// Handler .
type Handler interface {
	WriteFile(filename string, data []byte, perm os.FileMode) error
	ReadFile(filename string) ([]byte, error)
}

// FileHandler is an abstraction of ioutil.WriteFile and ioutil.ReadFile
type FileHandler struct {
}

// WriteFile implements the Handler interface that's been created so that ioutil.WriteFile can be mocked
func (w *FileHandler) WriteFile(filename string, data []byte, perm os.FileMode) error {
	todosterDir := getTodosterDir()

	return ioutil.WriteFile(todosterDir+"/"+filename, data, perm)
}

// ReadFile implements the Handler interface that's been created so that ioutil.ReadFile can be mocked
func (w *FileHandler) ReadFile(filename string) ([]byte, error) {
	todosterDir := getTodosterDir()

	return ioutil.ReadFile(todosterDir + "/" + filename)
}
