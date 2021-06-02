package backend_test

import (
	"errors"
	"imperial-splendour-launcher/backend/customErrors"
	"imperial-splendour-launcher/backend/test"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testing"
)

func TestUninstall(t *testing.T) {
	t.Run("Cannot launch uninstaller", func(t *testing.T) {
		api, _, window, _, sysHandler := test.Before()
		sysHandler.On("StartCommand", mock.Anything).Return(errors.New("Random Error launching the uninstaller"))

		err := api.Uninstall()

		assert.Equal(t, err, customErrors.Uninstall)
		sysHandler.AssertCalled(t, "StartCommand", "./IS_Uninstall/unins000.exe")
		window.AssertNotCalled(t, "Close")

		test.After(*api)
	})

	t.Run("Launches uninstaller", func(t *testing.T) {
		api, _, window, _, sysHandler := test.Before()
		sysHandler.On("StartCommand", mock.Anything).Return(nil)

		err := api.Uninstall()

		assert.Nil(t, err)
		sysHandler.AssertCalled(t, "StartCommand", "./IS_Uninstall/unins000.exe")
		window.AssertCalled(t, "Close")

		test.After(*api)
	})
}
