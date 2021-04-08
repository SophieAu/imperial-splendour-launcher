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

	t.Run("Can deactivate but cannot uninstall Imperial Splendour (deleting mod files)", func(t *testing.T) {
		api, _, _, _, sysHandler := testHelpers.VariableBefore("2.0", true, "test")
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte("merp.pack"), nil)
		sysHandler.On("MoveFile", testifyMock.Anything, testifyMock.Anything).Return(nil)
		sysHandler.On("WriteFile", "./IS_Files/IS_info.json", mock.Anything).Return(nil)
		sysHandler.On("Remove", testifyMock.Anything).Return(errors.New("Could not delete files")).Once()
		sysHandler.On("Remove", testifyMock.Anything).Return(nil).Once()
		sysHandler.On("Remove", testifyMock.Anything).Return(nil).Once()
		sysHandler.On("Remove", testifyMock.Anything).Return(nil).Once()
		sysHandler.On("Getenv", "USERPROFILE").Return("").Once()
		sysHandler.On("Getenv", "APPDATA").Return("").Once()

		err := api.Uninstall()

		assert.Equal(t, err, customErrors.Uninstall)
		sysHandler.AssertCalled(t, "MoveFile", "./data/merp.pack", "./IS_Files/merp.pack")
		sysHandler.AssertCalled(t, "Remove", "./IS_Files/")
		sysHandler.AssertCalled(t, "Remove", "./data/campaigns/imperial_splendour/")
		sysHandler.AssertCalled(t, "Remove", "/Desktop/Imperial Splendour.lnk")
		sysHandler.AssertCalled(t, "Remove", "/Microsoft/Windows/Start Menu/Programs/Imperial Splendour")

		testHelpers.After(*api)
	})

	t.Run("Can deactivate but cannot uninstall Imperial Splendour (deleting campaign folder)", func(t *testing.T) {
		api, _, _, _, sysHandler := testHelpers.VariableBefore("2.0", true, "test")
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte("merp.pack"), nil)
		sysHandler.On("MoveFile", testifyMock.Anything, testifyMock.Anything).Return(nil)
		sysHandler.On("WriteFile", "./IS_Files/IS_info.json", mock.Anything).Return(nil)
		sysHandler.On("Remove", testifyMock.Anything).Return(nil).Once()
		sysHandler.On("Remove", testifyMock.Anything).Return(errors.New("Could not delete folder")).Once()
		sysHandler.On("Remove", testifyMock.Anything).Return(nil).Once()
		sysHandler.On("Remove", testifyMock.Anything).Return(nil).Once()
		sysHandler.On("Getenv", "USERPROFILE").Return("").Once()
		sysHandler.On("Getenv", "APPDATA").Return("").Once()

		err := api.Uninstall()

		assert.Equal(t, err, customErrors.Uninstall)
		sysHandler.AssertCalled(t, "MoveFile", "./data/merp.pack", "./IS_Files/merp.pack")
		sysHandler.AssertCalled(t, "Remove", "./IS_Files/")
		sysHandler.AssertCalled(t, "Remove", "./data/campaigns/imperial_splendour/")
		sysHandler.AssertCalled(t, "Remove", "/Desktop/Imperial Splendour.lnk")
		sysHandler.AssertCalled(t, "Remove", "/Microsoft/Windows/Start Menu/Programs/Imperial Splendour")

		testHelpers.After(*api)
	})

	t.Run("Can deactivate but cannot uninstall Imperial Splendour (deleting campaign folder)", func(t *testing.T) {
		api, _, _, _, sysHandler := testHelpers.VariableBefore("2.0", true, "test")
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte("merp.pack"), nil)
		sysHandler.On("MoveFile", testifyMock.Anything, testifyMock.Anything).Return(nil)
		sysHandler.On("WriteFile", "./IS_Files/IS_info.json", mock.Anything).Return(nil)
		sysHandler.On("Remove", testifyMock.Anything).Return(errors.New("Could not delete files")).Once()
		sysHandler.On("Remove", testifyMock.Anything).Return(errors.New("Could not delete folder")).Once()
		sysHandler.On("Remove", testifyMock.Anything).Return(nil).Once()
		sysHandler.On("Remove", testifyMock.Anything).Return(nil).Once()
		sysHandler.On("Getenv", "USERPROFILE").Return("").Once()
		sysHandler.On("Getenv", "APPDATA").Return("").Once()

		err := api.Uninstall()

		assert.Equal(t, err, customErrors.Uninstall)
		sysHandler.AssertCalled(t, "MoveFile", "./data/merp.pack", "./IS_Files/merp.pack")
		sysHandler.AssertCalled(t, "Remove", "./IS_Files/")
		sysHandler.AssertCalled(t, "Remove", "./data/campaigns/imperial_splendour/")
		sysHandler.AssertCalled(t, "Remove", "/Desktop/Imperial Splendour.lnk")
		sysHandler.AssertCalled(t, "Remove", "/Microsoft/Windows/Start Menu/Programs/Imperial Splendour")

		testHelpers.After(*api)
	})

	t.Run("Deactivates and uninstalls Imperial Splendour even if there's an error deleting the desktop and startmenu shortcuts", func(t *testing.T) {
		api, _, _, _, sysHandler := testHelpers.VariableBefore("2.0", true, "test")
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte("merp.pack"), nil)
		sysHandler.On("MoveFile", testifyMock.Anything, testifyMock.Anything).Return(nil)
		sysHandler.On("WriteFile", "./IS_Files/IS_info.json", mock.Anything).Return(nil)
		sysHandler.On("Remove", testifyMock.Anything).Return(nil).Twice()
		sysHandler.On("Remove", "/Desktop/Imperial Splendour.lnk").Return(errors.New("Could not delete desktop shortcut")).Once()
		sysHandler.On("Remove", testifyMock.Anything).Return(errors.New("Could not delete startmenu shortcut")).Once()
		sysHandler.On("Getenv", "USERPROFILE").Return("").Once()
		sysHandler.On("Getenv", "APPDATA").Return("").Once()

		err := api.Uninstall()

		assert.Nil(t, err)
		sysHandler.AssertCalled(t, "MoveFile", "./data/merp.pack", "./IS_Files/merp.pack")
		sysHandler.AssertCalled(t, "Remove", "./IS_Files/")
		sysHandler.AssertCalled(t, "Remove", "./data/campaigns/imperial_splendour/")
		sysHandler.AssertCalled(t, "Remove", "/Desktop/Imperial Splendour.lnk")
		sysHandler.AssertCalled(t, "Remove", "/Microsoft/Windows/Start Menu/Programs/Imperial Splendour")

		testHelpers.After(*api)
	})

	t.Run("Doesn't need to deactivate Imperial Splendour but cannot uninstall it (deleting mod files)", func(t *testing.T) {
		api, _, _, _, sysHandler := testHelpers.VariableBefore("2.0", false, "test")
		sysHandler.On("Remove", testifyMock.Anything).Return(errors.New("Could not delete files")).Once()
		sysHandler.On("Remove", testifyMock.Anything).Return(nil)
		sysHandler.On("Getenv", "USERPROFILE").Return("").Once()
		sysHandler.On("Getenv", "APPDATA").Return("").Once()

		err := api.Uninstall()

		assert.Equal(t, err, customErrors.Uninstall)
		sysHandler.AssertNotCalled(t, "MoveFile", testifyMock.Anything, testifyMock.Anything)
		sysHandler.AssertCalled(t, "Remove", "./IS_Files/")
		sysHandler.AssertCalled(t, "Remove", "./data/campaigns/imperial_splendour/")
		sysHandler.AssertCalled(t, "Remove", "/Desktop/Imperial Splendour.lnk")
		sysHandler.AssertCalled(t, "Remove", "/Microsoft/Windows/Start Menu/Programs/Imperial Splendour")

		testHelpers.After(*api)
	})

	t.Run("Doesn't need to deactivate Imperial Splendour but cannot uninstall it (deleting campaign folder)", func(t *testing.T) {
		api, _, _, _, sysHandler := testHelpers.VariableBefore("2.0", false, "test")
		sysHandler.On("Remove", testifyMock.Anything).Return(nil).Once()
		sysHandler.On("Remove", testifyMock.Anything).Return(errors.New("Could not delete campaign folder")).Once()
		sysHandler.On("Remove", testifyMock.Anything).Return(nil)
		sysHandler.On("Getenv", "USERPROFILE").Return("").Once()
		sysHandler.On("Getenv", "APPDATA").Return("").Once()

		err := api.Uninstall()

		assert.Equal(t, err, customErrors.Uninstall)
		sysHandler.AssertNotCalled(t, "MoveFile", testifyMock.Anything, testifyMock.Anything)
		sysHandler.AssertCalled(t, "Remove", "./IS_Files/")
		sysHandler.AssertCalled(t, "Remove", "./data/campaigns/imperial_splendour/")
		sysHandler.AssertCalled(t, "Remove", "/Desktop/Imperial Splendour.lnk")
		sysHandler.AssertCalled(t, "Remove", "/Microsoft/Windows/Start Menu/Programs/Imperial Splendour")

		testHelpers.After(*api)
	})
	t.Run("Doesn't need to deactivate Imperial Splendour but cannot uninstall it (deleting campaign folder and mod files)", func(t *testing.T) {
		api, _, _, _, sysHandler := testHelpers.VariableBefore("2.0", false, "test")
		sysHandler.On("Remove", testifyMock.Anything).Return(errors.New("Could not delete files")).Once()
		sysHandler.On("Remove", testifyMock.Anything).Return(errors.New("Could not delete campaign folder")).Once()
		sysHandler.On("Remove", testifyMock.Anything).Return(nil)
		sysHandler.On("Getenv", "USERPROFILE").Return("").Once()
		sysHandler.On("Getenv", "APPDATA").Return("").Once()

		err := api.Uninstall()

		assert.Equal(t, err, customErrors.Uninstall)
		sysHandler.AssertNotCalled(t, "MoveFile", testifyMock.Anything, testifyMock.Anything)
		sysHandler.AssertCalled(t, "Remove", "./IS_Files/")
		sysHandler.AssertCalled(t, "Remove", "./data/campaigns/imperial_splendour/")
		sysHandler.AssertCalled(t, "Remove", "/Desktop/Imperial Splendour.lnk")
		sysHandler.AssertCalled(t, "Remove", "/Microsoft/Windows/Start Menu/Programs/Imperial Splendour")

		testHelpers.After(*api)
	})

	t.Run("Doesn't need to deactivate Imperial Splendour and uninstalls it", func(t *testing.T) {
		api, _, _, _, sysHandler := testHelpers.VariableBefore("2.0", false, "test")
		sysHandler.On("Remove", testifyMock.Anything).Return(nil)
		sysHandler.On("Getenv", "USERPROFILE").Return("").Once()
		sysHandler.On("Getenv", "APPDATA").Return("").Once()

		err := api.Uninstall()

		assert.Nil(t, err)
		sysHandler.AssertNotCalled(t, "MoveFile", testifyMock.Anything, testifyMock.Anything)
		sysHandler.AssertCalled(t, "Remove", "./IS_Files/")
		sysHandler.AssertCalled(t, "Remove", "./data/campaigns/imperial_splendour/")
		sysHandler.AssertCalled(t, "Remove", "/Desktop/Imperial Splendour.lnk")
		sysHandler.AssertCalled(t, "Remove", "/Microsoft/Windows/Start Menu/Programs/Imperial Splendour")

		testHelpers.After(*api)
	})
}
