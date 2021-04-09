package backend_test

import (
	"imperial-splendour-launcher/backend/test"

	"testing"
)

func TestExit(t *testing.T) {
	t.Run("Close Window", func(t *testing.T) {
		api, _, window, _, _ := test.Before()

		window.On("Close").Return().Once()
		api.Exit()

		window.AssertCalled(t, "Close")

		test.After(*api)
	})
}
