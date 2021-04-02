package backend_test

import (
	"imperial-splendour-launcher/backend/testHelpers"

	"testing"
)

func TestExit(t *testing.T) {
	t.Run("Close Window", func(t *testing.T) {
		api, _, window, _, _ := testHelpers.Before()

		window.On("Close").Return().Once()
		api.Exit()

		window.AssertCalled(t, "Close")

		testHelpers.After(*api)
	})
}
