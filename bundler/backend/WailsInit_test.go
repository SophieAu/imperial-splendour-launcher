package backend_test

import (
	"imperial-splendour-bundler/backend"
	"imperial-splendour-bundler/backend/mocks"
	"imperial-splendour-bundler/backend/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {

	t.Run("No issues initializing", func(t *testing.T) {
		mockSh := &mocks.MockSystemHandler{}
		mockB := &mocks.MockBrowser{}
		mockW := &mocks.MockWindow{}
		mockL := &mocks.MockLogger{}
		mockD := &mocks.MockDialog{}
		mockSt := &mocks.MockStore{}
		api := &backend.API{}

		err := api.Init(mockB, mockW, mockL, mockSt, mockD, mockSh)

		assert.Nil(t, err)

		test.After(*api)
	})
}
