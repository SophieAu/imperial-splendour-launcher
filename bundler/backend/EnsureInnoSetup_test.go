package backend_test

import (
	"errors"
	"imperial-splendour-bundler/backend/customErrors"
	"imperial-splendour-bundler/backend/test"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestEnsureInnoSetup(t *testing.T) {

	t.Run("InnoSetup is not present", func(t *testing.T) {
		api, _, _, _, _, mockSh := test.Before()
		mockSh.On("StartCommand", mock.Anything, mock.Anything).Return(errors.New("No Innosetup installed"))

		err := api.EnsureInnoSetup()

		assert.Equal(t, customErrors.InnoSetup, err)

		test.After(*api)
	})

	t.Run("InnoSetup is present", func(t *testing.T) {
		api, _, _, _, _, mockSh := test.Before()

		mockSh.On("StartCommand", mock.Anything, mock.Anything).Return(nil)

		err := api.EnsureInnoSetup()

		assert.Nil(t, err)

		test.After(*api)
	})
}
