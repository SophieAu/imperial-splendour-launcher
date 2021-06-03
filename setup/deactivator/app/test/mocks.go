package test

import (
	"github.com/stretchr/testify/mock"
)

// System Handler
type MockSystemHandler struct {
	mock.Mock
}

func (sh *MockSystemHandler) WriteFile(filePath string, data []byte) error {
	args := sh.Called(filePath, data)
	return args.Error(0)
}

func (sh *MockSystemHandler) ReadFile(filePath string) ([]byte, error) {
	args := sh.Called(filePath)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]byte), args.Error(1)
}

func (sh *MockSystemHandler) MoveFile(source, destination string) error {
	args := sh.Called(source, destination)
	return args.Error(0)
}

func (sh *MockSystemHandler) Remove(path string) error {
	args := sh.Called(path)
	return args.Error(0)
}

func (sh *MockSystemHandler) Executable() (string, error) {
	args := sh.Called()
	return args.Get(0).(string), args.Error(1)
}

func (sh *MockSystemHandler) Getenv(key string) string {
	args := sh.Called(key)
	return args.Get(0).(string)
}
