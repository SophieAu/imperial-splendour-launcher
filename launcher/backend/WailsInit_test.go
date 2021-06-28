package backend_test

import (
	"errors"
	"imperial-splendour-launcher/backend"
	"imperial-splendour-launcher/backend/mocks"
	"imperial-splendour-launcher/backend/test"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInit(t *testing.T) {
	logger := &mocks.MockLogger{}
	logger.On("Infof", mock.Anything, mock.Anything).Return()
	logger.On("Warnf", mock.Anything, mock.Anything).Return()

	t.Run("Cannot get current exe's directory", func(t *testing.T) {
		sysHandler := &mocks.MockSystemHandler{}
		api := &backend.API{}
		sysHandler.On("Executable").Return("", errors.New("Error getting exe dir")).Once()

		err := api.Init(&mocks.MockBrowser{}, &mocks.MockWindow{}, logger, sysHandler)

		assert.EqualError(t, err, "Error getting exe dir")
		sysHandler.AssertNotCalled(t, "ReadFile", mock.Anything)
	})

	t.Run("Cannot get app data directory", func(t *testing.T) {
		sysHandler := &mocks.MockSystemHandler{}
		api := &backend.API{}
		sysHandler.On("Executable").Return(".", nil).Once()
		sysHandler.On("Getenv", "APPDATA").Return("").Once()

		err := api.Init(&mocks.MockBrowser{}, &mocks.MockWindow{}, logger, sysHandler)

		assert.EqualError(t, err, "Couldn't get user's APPDATA dir")
		sysHandler.AssertNotCalled(t, "ReadFile", mock.Anything)
	})

	// load from info file
	t.Run("Cannot read info file", func(t *testing.T) {
		sysHandler := &mocks.MockSystemHandler{}
		api := &backend.API{}
		sysHandler.On("Executable").Return(".", nil).Once()
		sysHandler.On("Getenv", "APPDATA").Return("APPDATA")
		sysHandler.On("ReadFile", mock.Anything).Return(nil, errors.New("FileNotFound")).Once()

		err := api.Init(&mocks.MockBrowser{}, &mocks.MockWindow{}, logger, sysHandler)

		sysHandler.AssertCalled(t, "ReadFile", "./IS_Files/IS_Info.json")
		assert.EqualError(t, err, "FileNotFound")
	})

	t.Run("Cannot unmarshal info file", func(t *testing.T) {
		sysHandler := &mocks.MockSystemHandler{}
		api := &backend.API{}
		sysHandler.On("Executable").Return(".", nil).Once()
		sysHandler.On("Getenv", "APPDATA").Return("APPDATA")
		sysHandler.On("ReadFile", mock.Anything).Return([]byte{}, nil).Once()

		err := api.Init(&mocks.MockBrowser{}, &mocks.MockWindow{}, logger, sysHandler)

		sysHandler.AssertCalled(t, "ReadFile", "./IS_Files/IS_Info.json")
		assert.EqualError(t, err, "unexpected end of JSON input")
	})

	t.Run("No user script checksum in info file", func(t *testing.T) {
		sysHandler := &mocks.MockSystemHandler{}
		api := &backend.API{}
		sysHandler.On("Executable").Return(".", nil).Once()
		sysHandler.On("Getenv", "APPDATA").Return("APPDATA")
		sysHandler.On("ReadFile", mock.Anything).Return(test.FmtInfoFile(true, "2.0", ""), nil).Once()

		err := api.Init(&mocks.MockBrowser{}, &mocks.MockWindow{}, logger, sysHandler)

		sysHandler.AssertCalled(t, "ReadFile", "./IS_Files/IS_Info.json")
		assert.EqualError(t, err, "Corrupt Info File")
	})

	t.Run("No version in info file", func(t *testing.T) {
		sysHandler := &mocks.MockSystemHandler{}
		api := &backend.API{}
		sysHandler.On("Executable").Return(".", nil).Once()
		sysHandler.On("Getenv", "APPDATA").Return("APPDATA")
		sysHandler.On("ReadFile", mock.Anything).Return(test.FmtInfoFile(true, "", "test"), nil).Once()

		err := api.Init(&mocks.MockBrowser{}, &mocks.MockWindow{}, logger, sysHandler)

		sysHandler.AssertCalled(t, "ReadFile", "./IS_Files/IS_Info.json")
		assert.EqualError(t, err, "Corrupt Info File")
	})

	t.Run("Successfully initialize the Launcher", func(t *testing.T) {
		sysHandler := &mocks.MockSystemHandler{}
		logger := &mocks.MockLogger{}
		api := &backend.API{}
		logger.On("Infof", mock.Anything, mock.Anything).Return()
		sysHandler.On("Executable").Return(".", nil).Once()
		sysHandler.On("Getenv", "APPDATA").Return("APPDATA")
		sysHandler.On("ReadFile", mock.Anything).Return(test.FmtInfoFile(true, "2.0", "test"), nil).Once()

		err := api.Init(&mocks.MockBrowser{}, &mocks.MockWindow{}, logger, sysHandler)

		logger.AssertCalled(t, "Infof", test.ExpectFmt("Info loaded %v", "{true 2.0 test}"))
		assert.Nil(t, err)
	})
}
