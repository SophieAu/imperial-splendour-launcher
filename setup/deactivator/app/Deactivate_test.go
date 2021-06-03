package app_test

import (
	"deactivator/app/test"
	"errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testing"
)

// NOTE: Deactivating Imperial Splendour is tested more thoroughly in the Launcher in the Switch_test.go file
func TestUninstall(t *testing.T) {
	t.Run("Doesn't need to deactivate Imperial Splendour", func(t *testing.T) {
		api, sysHandler := test.VariableBefore("2.0", false, "test")

		err := api.Deactivate()

		assert.Nil(t, err)
		sysHandler.AssertNotCalled(t, "MoveFile", mock.Anything, mock.Anything)

		test.After(*api)
	})

	t.Run("Cannot deactivate Imperial Splendour", func(t *testing.T) {
		api, sysHandler := test.VariableBefore("2.0", true, "test")
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return(nil, errors.New("Error Deactivating")).Once()

		err := api.Deactivate()

		assert.EqualError(t, err, "FileListError")
		sysHandler.AssertNotCalled(t, "MoveFile", mock.Anything, mock.Anything)

		test.After(*api)
	})

	t.Run("Can deactivate Imperial Splendour but has silent errors", func(t *testing.T) {
		api, sysHandler := test.VariableBefore("2.0", true, "test")
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte("merp.pack\nderp.tga"), nil)
		sysHandler.On("MoveFile", mock.Anything, mock.Anything).Return(nil)
		err := api.Deactivate()

		assert.Nil(t, err)
		sysHandler.AssertCalled(t, "MoveFile", "./data/merp.pack", "./IS_Files/merp.pack")
		sysHandler.AssertCalled(t, "MoveFile", "./data/campaigns/imperial_splendour/derp.tga", "./IS_Files/derp.tga")
		sysHandler.AssertCalled(t, "MoveFile", "APPDATA/The Creative Assembly/Empire/scripts/user.empire_script.txt", "./IS_Files/user.empire_script.txt")

		test.After(*api)
	})

	t.Run("Can deactivate Imperial Splendour", func(t *testing.T) {
		api, sysHandler := test.VariableBefore("2.0", true, "test")
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte("merp.pack"), nil)
		sysHandler.On("MoveFile", mock.Anything, mock.Anything).Return(nil)

		err := api.Deactivate()

		assert.Nil(t, err)
		sysHandler.AssertCalled(t, "MoveFile", "./data/merp.pack", "./IS_Files/merp.pack")
		sysHandler.AssertCalled(t, "MoveFile", "APPDATA/The Creative Assembly/Empire/scripts/user.empire_script.txt", "./IS_Files/user.empire_script.txt")

		test.After(*api)
	})
}
