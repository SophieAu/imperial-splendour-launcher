package backend_test

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestVersion(t *testing.T) {
	api, _, _, _, _ := variableBefore("2.1", false, "test")

	version := api.Version()
	assert.Equal(t, "2.1", version)

	after(*api)
}
