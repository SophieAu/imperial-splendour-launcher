package backend_test

import (
	"errors"
	"imperial-splendour-bundler/backend"
	"imperial-splendour-bundler/backend/customErrors"
	"imperial-splendour-bundler/backend/mocks"
	"imperial-splendour-bundler/backend/test"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInit(t *testing.T) {

	t.Run("InnoSetup is not present", func(t *testing.T) {
		mockSh := &mocks.MockSystemHandler{}
		mockB := &mocks.MockBrowser{}
		mockW := &mocks.MockWindow{}
		mockL := &mocks.MockLogger{}
		mockD := &mocks.MockDialog{}
		mockSt := &mocks.MockStore{}
		api := &backend.API{}

		mockSh.On("StartCommand", mock.Anything, mock.Anything).Return(errors.New("No Innosetup installed"))
		mockL.On("Warn", mock.Anything).Return()

		err := api.Init(mockB, mockW, mockL, mockSt, mockD, mockSh)

		assert.Equal(t, err, customErrors.InnoSetup)

		test.After(*api)
	})

	t.Run("InnoSetup is present", func(t *testing.T) {
		mockSh := &mocks.MockSystemHandler{}
		mockB := &mocks.MockBrowser{}
		mockW := &mocks.MockWindow{}
		mockL := &mocks.MockLogger{}
		mockD := &mocks.MockDialog{}
		mockSt := &mocks.MockStore{}
		api := &backend.API{}

		mockSh.On("StartCommand", mock.Anything, mock.Anything).Return(nil)
		mockL.On("Warn", mock.Anything).Return()

		err := api.Init(mockB, mockW, mockL, mockSt, mockD, mockSh)

		assert.Nil(t, err)

		test.After(*api)
	})
}
