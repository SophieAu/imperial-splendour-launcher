package backend_test

import (
	"imperial-splendour-launcher/backend/testHelpers"

	"github.com/stretchr/testify/assert"

	"testing"
)

func TestIsActive(t *testing.T) {
	api, _, _, _, _ := testHelpers.VariableBefore("2.1", false, "test")

	isActive := api.IsActive()
	assert.Equal(t, false, isActive)

	api, _, _, _, _ = testHelpers.VariableBefore("2.1", true, "test")

	isActive = api.IsActive()
	assert.Equal(t, true, isActive)

	testHelpers.After(*api)
}
