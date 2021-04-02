package backend_test

import (
	"errors"
	"imperial-splendour-launcher/backend"
	"imperial-splendour-launcher/backend/mock"
	"imperial-splendour-launcher/backend/testHelpers"

	"testing"

	"github.com/stretchr/testify/assert"
	testifyMock "github.com/stretchr/testify/mock"
)

func TestInit(t *testing.T) {
	logger := &mock.Logger{}
	logger.On("Infof", testifyMock.Anything, testifyMock.Anything).Return()
	logger.On("Warnf", testifyMock.Anything, testifyMock.Anything).Return()

	t.Run("Cannot get current exe's directory", func(t *testing.T) {
		sysHandler := &mock.SystemHandler{}
		api := &backend.API{}
		sysHandler.On("Executable").Return("", errors.New("Error getting exe dir")).Once()

		err := api.Init(&mock.Browser{}, &mock.Window{}, logger, sysHandler)

		assert.EqualError(t, err, "Error getting exe dir")
		sysHandler.AssertNotCalled(t, "ReadFile", testifyMock.Anything)
	})

	t.Run("Cannot get app data directory", func(t *testing.T) {
		sysHandler := &mock.SystemHandler{}
		api := &backend.API{}
		sysHandler.On("Executable").Return(".", nil).Once()
		sysHandler.On("Getenv", "APPDATA").Return("").Once()

		err := api.Init(&mock.Browser{}, &mock.Window{}, logger, sysHandler)

		assert.EqualError(t, err, "Couldn't get user's APPDATA dir")
		sysHandler.AssertNotCalled(t, "ReadFile", testifyMock.Anything)
	})

	// load from info file
	t.Run("Cannot read info file", func(t *testing.T) {
		sysHandler := &mock.SystemHandler{}
		api := &backend.API{}
		sysHandler.On("Executable").Return(".", nil).Once()
		sysHandler.On("Getenv", "APPDATA").Return("APPDATA")
		sysHandler.On("ReadFile", testifyMock.Anything).Return(nil, errors.New("FileNotFound")).Once()

		err := api.Init(&mock.Browser{}, &mock.Window{}, logger, sysHandler)

		sysHandler.AssertCalled(t, "ReadFile", "./IS_Files/IS_info.json")
		assert.EqualError(t, err, "FileNotFound")
	})

	t.Run("Cannot unmarshal info file", func(t *testing.T) {
		sysHandler := &mock.SystemHandler{}
		api := &backend.API{}
		sysHandler.On("Executable").Return(".", nil).Once()
		sysHandler.On("Getenv", "APPDATA").Return("APPDATA")
		sysHandler.On("ReadFile", testifyMock.Anything).Return([]byte{}, nil).Once()

		err := api.Init(&mock.Browser{}, &mock.Window{}, logger, sysHandler)

		sysHandler.AssertCalled(t, "ReadFile", "./IS_Files/IS_info.json")
		assert.EqualError(t, err, "unexpected end of JSON input")
	})

	t.Run("No user script checksum in info file", func(t *testing.T) {
		sysHandler := &mock.SystemHandler{}
		api := &backend.API{}
		sysHandler.On("Executable").Return(".", nil).Once()
		sysHandler.On("Getenv", "APPDATA").Return("APPDATA")
		sysHandler.On("ReadFile", testifyMock.Anything).Return(testHelpers.FmtInfoFile(true, "2.0", ""), nil).Once()

		err := api.Init(&mock.Browser{}, &mock.Window{}, logger, sysHandler)

		sysHandler.AssertCalled(t, "ReadFile", "./IS_Files/IS_info.json")
		assert.EqualError(t, err, "Corrupt Info File")
	})

	t.Run("No version in info file", func(t *testing.T) {
		sysHandler := &mock.SystemHandler{}
		api := &backend.API{}
		sysHandler.On("Executable").Return(".", nil).Once()
		sysHandler.On("Getenv", "APPDATA").Return("APPDATA")
		sysHandler.On("ReadFile", testifyMock.Anything).Return(testHelpers.FmtInfoFile(true, "", "test"), nil).Once()

		err := api.Init(&mock.Browser{}, &mock.Window{}, logger, sysHandler)

		sysHandler.AssertCalled(t, "ReadFile", "./IS_Files/IS_info.json")
		assert.EqualError(t, err, "Corrupt Info File")
	})

	t.Run("Successfully initialize the Launcher", func(t *testing.T) {
		sysHandler := &mock.SystemHandler{}
		logger := &mock.Logger{}
		api := &backend.API{}
		logger.On("Infof", testifyMock.Anything, testifyMock.Anything).Return()
		sysHandler.On("Executable").Return(".", nil).Once()
		sysHandler.On("Getenv", "APPDATA").Return("APPDATA")
		sysHandler.On("ReadFile", testifyMock.Anything).Return(testHelpers.FmtInfoFile(true, "2.0", "test"), nil).Once()

		err := api.Init(&mock.Browser{}, &mock.Window{}, logger, sysHandler)

		logger.AssertCalled(t, "Infof", testHelpers.ExpectFmt("Info loaded %v", "{true 2.0 test}"))
		assert.Nil(t, err)
	})
}
