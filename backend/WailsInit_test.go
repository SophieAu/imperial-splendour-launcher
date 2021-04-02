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
	mockL := &mock.Logger{}
	mockL.On("Infof", testifyMock.Anything, testifyMock.Anything).Return()
	mockL.On("Warnf", testifyMock.Anything, testifyMock.Anything).Return()

	t.Run("Cannot get current exe's directory", func(t *testing.T) {
		mockS := &mock.SystemHandler{}
		api := &backend.API{}
		mockS.On("Executable").Return("", errors.New("Error getting exe dir")).Once()

		err := api.Init(&mock.Browser{}, &mock.Window{}, mockL, mockS)

		assert.Equal(t, "Error getting exe dir", err.Error())
		mockS.AssertNotCalled(t, "ReadFile", testifyMock.Anything)
	})

	t.Run("Cannot get app data directory", func(t *testing.T) {
		mockS := &mock.SystemHandler{}
		api := &backend.API{}
		mockS.On("Executable").Return(".", nil).Once()
		mockS.On("Getenv", "APPDATA").Return("").Once()

		err := api.Init(&mock.Browser{}, &mock.Window{}, mockL, mockS)

		assert.Equal(t, "Couldn't get user's APPDATA dir", err.Error())
		mockS.AssertNotCalled(t, "ReadFile", testifyMock.Anything)
	})

	// load from info file
	t.Run("Cannot read info file", func(t *testing.T) {
		mockS := &mock.SystemHandler{}
		api := &backend.API{}
		mockS.On("Executable").Return(".", nil).Once()
		mockS.On("Getenv", "APPDATA").Return("APPDATA")
		mockS.On("ReadFile", testifyMock.Anything).Return(nil, errors.New("FileNotFound")).Once()

		err := api.Init(&mock.Browser{}, &mock.Window{}, mockL, mockS)

		mockS.AssertCalled(t, "ReadFile", "./IS_Files/IS_info.json")
		assert.Equal(t, "FileNotFound", err.Error())
	})

	t.Run("Cannot unmarshal info file", func(t *testing.T) {
		mockS := &mock.SystemHandler{}
		api := &backend.API{}
		mockS.On("Executable").Return(".", nil).Once()
		mockS.On("Getenv", "APPDATA").Return("APPDATA")
		mockS.On("ReadFile", testifyMock.Anything).Return([]byte{}, nil).Once()

		err := api.Init(&mock.Browser{}, &mock.Window{}, mockL, mockS)

		mockS.AssertCalled(t, "ReadFile", "./IS_Files/IS_info.json")
		assert.Equal(t, "unexpected end of JSON input", err.Error())
	})

	t.Run("No user script checksum in info file", func(t *testing.T) {
		mockS := &mock.SystemHandler{}
		api := &backend.API{}
		mockS.On("Executable").Return(".", nil).Once()
		mockS.On("Getenv", "APPDATA").Return("APPDATA")
		mockS.On("ReadFile", testifyMock.Anything).Return(testHelpers.FmtInfoFile(true, "2.0", ""), nil).Once()

		err := api.Init(&mock.Browser{}, &mock.Window{}, mockL, mockS)

		mockS.AssertCalled(t, "ReadFile", "./IS_Files/IS_info.json")
		assert.Equal(t, "Corrupt Info File", err.Error())
	})

	t.Run("No version in info file", func(t *testing.T) {
		mockS := &mock.SystemHandler{}
		api := &backend.API{}
		mockS.On("Executable").Return(".", nil).Once()
		mockS.On("Getenv", "APPDATA").Return("APPDATA")
		mockS.On("ReadFile", testifyMock.Anything).Return(testHelpers.FmtInfoFile(true, "", "test"), nil).Once()

		err := api.Init(&mock.Browser{}, &mock.Window{}, mockL, mockS)

		mockS.AssertCalled(t, "ReadFile", "./IS_Files/IS_info.json")
		assert.Equal(t, "Corrupt Info File", err.Error())
	})

	t.Run("Successfully initialize the Launcher", func(t *testing.T) {
		mockS := &mock.SystemHandler{}
		mockL := &mock.Logger{}
		api := &backend.API{}
		mockL.On("Infof", testifyMock.Anything, testifyMock.Anything).Return()
		mockS.On("Executable").Return(".", nil).Once()
		mockS.On("Getenv", "APPDATA").Return("APPDATA")
		mockS.On("ReadFile", testifyMock.Anything).Return(testHelpers.FmtInfoFile(true, "2.0", "test"), nil).Once()

		err := api.Init(&mock.Browser{}, &mock.Window{}, mockL, mockS)

		mockL.AssertCalled(t, "Infof", testHelpers.ExpectFmt("Info loaded %v", "{true 2.0 test}"))
		assert.Nil(t, err)
	})
}
