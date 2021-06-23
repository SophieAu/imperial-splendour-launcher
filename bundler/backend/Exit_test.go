package backend_test

import (
	"imperial-splendour-bundler/backend/test"

	"testing"
)

func TestExit(t *testing.T) {
	t.Run("Close Window", func(t *testing.T) {
		api, _, mockW, _, _, _ := test.Before()

		mockW.On("Close").Return().Once()
		api.Exit()

		mockW.AssertCalled(t, "Close")

		test.After(*api)
	})
}
