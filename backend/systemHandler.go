package backend

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
)

// Handler is an abstraction of all direct system interaction
type Handler interface {
	WriteFile(filePath string, data []byte) error
	ReadFile(filePath string) ([]byte, error)
	MoveFile(source, destination string) error
	Remove(path string) error
	Executable() (string, error)
	Getenv(key string) string
}

type SystemHandler struct {
}

func rename(source, destination string) error {
	return os.Rename(source, destination)
}

func moveAcrossDrives(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return errors.New("Couldn't open source file: " + err.Error())
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return errors.New("Couldn't open dest file: " + err.Error())
	}

	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return errors.New("Writing to output file failed: " + err.Error())
	}

	// The copy was successful, so now delete the original file
	if os.Remove(sourcePath); err != nil {
		return errors.New("Failed removing original file: " + err.Error())
	}

	return nil
}

func (w *SystemHandler) WriteFile(filePath string, data []byte) error {
	return ioutil.WriteFile(filePath, data, os.FileMode(0644))
}

func (w *SystemHandler) ReadFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

func (w *SystemHandler) Remove(path string) error {
	return os.RemoveAll(path)
}

func (w *SystemHandler) Executable() (string, error) {
	return os.Executable()
}

func (w *SystemHandler) Getenv(key string) string {
	return os.Getenv(key)
}

func (w *SystemHandler) MoveFile(source, destination string) error {
	var moveFunc func(sourcePath string, destPath string) error
	sourceDrive := source
	targetDrive := destination

	if sourceDrive != targetDrive {
		moveFunc = moveAcrossDrives
	} else {
		moveFunc = rename
	}

	return moveFunc(source, destination)
}
