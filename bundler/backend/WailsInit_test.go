package backend_test

import (
	"errors"
	"imperial-splendour-bundler/backend"
	"imperial-splendour-bundler/backend/customErrors"
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
		mockL.On("Warn", mock.Anything).Return()

		err := api.Init(mockB, mockW, mockL, mockS)

		assert.Equal(t, err, customErrors.InnoSetup)

		test.After(*api)
	})

	t.Run("InnoSetup is present", func(t *testing.T) {
		mockS := &test.MockSystemHandler{}
		mockB := &test.MockBrowser{}
		mockW := &test.MockWindow{}
		mockL := &test.MockLogger{}
		api := &backend.API{}

		mockS.On("RunCommand", mock.Anything, mock.Anything).Return(nil)
		mockL.On("Warn", mock.Anything).Return()

		err := api.Init(mockB, mockW, mockL, mockS)

		assert.Nil(t, err)

		test.After(*api)
	})
}
