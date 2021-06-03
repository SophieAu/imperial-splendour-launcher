package backend_test

import (
	"imperial-splendour-launcher/backend/test"

	"github.com/stretchr/testify/assert"

	"testing"
)

func TestVersion(t *testing.T) {
	t.Run("Return version", func(t *testing.T) {
		api, _, _, _, _ := test.VariableBefore("2.1", false, "test")

		version := api.Version()
		assert.Equal(t, "2.1", version)

		test.After(*api)
	})
}
