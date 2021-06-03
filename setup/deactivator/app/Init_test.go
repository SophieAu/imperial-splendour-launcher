package app_test

import (
	"errors"
	"deactivator/app"
	"deactivator/app/test"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInit(t *testing.T) {
	t.Run("Cannot get current exe's directory", func(t *testing.T) {
		sysHandler := &test.MockSystemHandler{}
		api := &app.API{}
		sysHandler.On("Executable").Return("", errors.New("Error getting exe dir")).Once()

		err := api.Init(sysHandler)

		assert.EqualError(t, err, "Error getting exe dir")
		sysHandler.AssertNotCalled(t, "ReadFile", mock.Anything)
	})

	t.Run("Cannot get app data directory", func(t *testing.T) {
		sysHandler := &test.MockSystemHandler{}
		api := &app.API{}
		sysHandler.On("Executable").Return(".", nil).Once()
		sysHandler.On("Getenv", "APPDATA").Return("").Once()

		err := api.Init(sysHandler)

		assert.EqualError(t, err, "Couldn't get user's APPDATA dir")
		sysHandler.AssertNotCalled(t, "ReadFile", mock.Anything)
	})

	// NOTE: Extensive testing is done in the Launcher
	t.Run("Cannot read info file", func(t *testing.T) {
		sysHandler := &test.MockSystemHandler{}
		api := &app.API{}
		sysHandler.On("Executable").Return("./IS_Uninstall", nil).Once()
		sysHandler.On("Getenv", "APPDATA").Return("APPDATA")
		sysHandler.On("ReadFile", mock.Anything).Return(nil, errors.New("FileNotFound")).Once()

		err := api.Init(sysHandler)

		assert.EqualError(t, err, "FileNotFound")
		sysHandler.AssertCalled(t, "ReadFile", "./IS_Files/IS_info.json")
	})

	t.Run("Successfully initialize the Launcher", func(t *testing.T) {
		sysHandler := &test.MockSystemHandler{}
		api := &app.API{}
		sysHandler.On("Executable").Return(".", nil).Once()
		sysHandler.On("Getenv", "APPDATA").Return("APPDATA")
		sysHandler.On("ReadFile", mock.Anything).Return(test.FmtInfoFile(true, "2.0", "test"), nil).Once()

		err := api.Init(sysHandler)

		assert.Nil(t, err)
		sysHandler.AssertCalled(t, "ReadFile", "./IS_Files/IS_info.json")
	})
}
