package backend_test

import (
	"errors"
	"imperial-splendour-launcher/backend/customErrors"
	"imperial-splendour-launcher/backend/testHelpers"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	testifyMock "github.com/stretchr/testify/mock"

	"testing"
)

func TestUninstall(t *testing.T) {
	t.Run("Cannot deactivate and therefore uninstall Imperial Splendour", func(t *testing.T) {
		// NOTE: Deactivating Imperial Splendour is tested more thoroughly in the Switch_test.go file
		api, _, _, _, sysHandler := testHelpers.VariableBefore("2.0", true, "test")
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return(nil, errors.New("Error Deactivating")).Once()

		err := api.Uninstall()

		assert.Equal(t, err, customErrors.Deactivation)
		sysHandler.AssertNotCalled(t, "Remove", testifyMock.Anything)

		testHelpers.After(*api)
	})

	t.Run("Can deactivate but cannot uninstall Imperial Splendour", func(t *testing.T) {
		api, _, _, _, sysHandler := testHelpers.VariableBefore("2.0", true, "test")
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte("merp.pack"), nil)
		sysHandler.On("MoveFile", testifyMock.Anything, testifyMock.Anything).Return(nil)
		sysHandler.On("WriteFile", "./IS_Files/IS_info.json", mock.Anything).Return(nil)
		sysHandler.On("Remove", testifyMock.Anything, testifyMock.Anything).Return(errors.New("Could not delete files"))

		err := api.Uninstall()

		assert.Equal(t, err, customErrors.Uninstall)
		sysHandler.AssertCalled(t, "MoveFile", "./data/merp.pack", "./IS_Files/merp.pack")
		sysHandler.AssertCalled(t, "Remove", "./IS_Files/")

		testHelpers.After(*api)
	})

	t.Run("Deactivates and uninstalls Imperial Splendour", func(t *testing.T) {
		api, _, _, _, sysHandler := testHelpers.VariableBefore("2.0", true, "test")
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte("merp.pack"), nil)
		sysHandler.On("MoveFile", testifyMock.Anything, testifyMock.Anything).Return(nil)
		sysHandler.On("WriteFile", "./IS_Files/IS_info.json", mock.Anything).Return(nil)
		sysHandler.On("Remove", testifyMock.Anything, testifyMock.Anything).Return(nil)

		err := api.Uninstall()

		assert.Nil(t, err)
		sysHandler.AssertCalled(t, "MoveFile", "./data/merp.pack", "./IS_Files/merp.pack")
		sysHandler.AssertCalled(t, "Remove", "./IS_Files/")

		testHelpers.After(*api)
	})

	t.Run("Doesn't need to deactivate Imperial Splendour but cannot uninstall it", func(t *testing.T) {
		api, _, _, _, sysHandler := testHelpers.VariableBefore("2.0", false, "test")
		sysHandler.On("Remove", testifyMock.Anything, testifyMock.Anything).Return(errors.New("Could not delete files"))

		err := api.Uninstall()

		assert.Equal(t, err, customErrors.Uninstall)
		sysHandler.AssertNotCalled(t, "MoveFile", testifyMock.Anything, testifyMock.Anything)
		sysHandler.AssertCalled(t, "Remove", "./IS_Files/")

		testHelpers.After(*api)
	})

	t.Run("Doesn't need to deactivate Imperial Splendour and uninstalls it", func(t *testing.T) {
		api, _, _, _, sysHandler := testHelpers.VariableBefore("2.0", false, "test")
		sysHandler.On("Remove", testifyMock.Anything, testifyMock.Anything).Return(nil)

		err := api.Uninstall()

		assert.Nil(t, err)
		sysHandler.AssertNotCalled(t, "MoveFile", testifyMock.Anything, testifyMock.Anything)
		sysHandler.AssertCalled(t, "Remove", "./IS_Files/")

		testHelpers.After(*api)
	})
}
