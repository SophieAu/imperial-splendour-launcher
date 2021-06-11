package backend_test

import (
	"errors"
	"imperial-splendour-bundler/backend"
	"imperial-splendour-bundler/backend/test"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInit(t *testing.T) {

	t.Run("InnoSetup is not present", func(t *testing.T) {
		mockS := &test.MockSystemHandler{}
		mockB := &test.MockBrowser{}
		mockW := &test.MockWindow{}
		mockL := &test.MockLogger{}
		api := &backend.API{}

		mockS.On("RunCommand", mock.Anything, mock.Anything).Return(errors.New("No Innosetup installed"))

		err := api.Init(mockB, mockW, mockL, mockS)

		assert.EqualError(t, err, "No Innosetup installed")

		test.After(*api)
	})

	t.Run("InnoSetup is present", func(t *testing.T) {
		mockS := &test.MockSystemHandler{}
		mockB := &test.MockBrowser{}
		mockW := &test.MockWindow{}
		mockL := &test.MockLogger{}
		api := &backend.API{}

		mockS.On("RunCommand", mock.Anything, mock.Anything).Return(nil)

		err := api.Init(mockB, mockW, mockL, mockS)

		assert.Nil(t, err)

		test.After(*api)
	})
}
