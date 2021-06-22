package backend

import (
	"imperial-splendour-bundler/backend/customErrors"
	"imperial-splendour-bundler/backend/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testing"
)

func TestUserInputValidation(t *testing.T) {
	mockL := &mocks.MockLogger{}
	mockL.On("Warn", mock.Anything).Return()

	t.Run("Validation of source path fails", func(t *testing.T) {
		api := &API{logger: mockL}
		err := api.validateUserInput("", "", "")

		assert.Equal(t, err, customErrors.EmptySource)
	})

	t.Run("Validation of version number fails", func(t *testing.T) {
		api := &API{logger: mockL}
		err := api.validateUserInput("s", "a", "")

		assert.Equal(t, err, customErrors.InvalidVersion)
	})

	t.Run("Validation of version number fails", func(t *testing.T) {
		api := &API{logger: mockL}
		err := api.validateUserInput("s", "2.0.0.0", "")

		assert.Equal(t, err, customErrors.InvalidVersion)
	})

	t.Run("Validation of version number fails", func(t *testing.T) {
		api := &API{logger: mockL}
		err := api.validateUserInput("s", "2", "")

		assert.Equal(t, err, customErrors.InvalidVersion)
	})

	t.Run("Validation of file list path fails", func(t *testing.T) {
		api := &API{logger: mockL}
		err := api.validateUserInput("s", "2.0.0", "")

		assert.Equal(t, err, customErrors.NoFileList)
	})

	t.Run("Validation of file list path fails", func(t *testing.T) {
		api := &API{logger: mockL}
		err := api.validateUserInput("s", "2.0", "")

		assert.Equal(t, err, customErrors.NoFileList)
	})

	t.Run("Validation succeeds", func(t *testing.T) {
		api := &API{logger: mockL}
		err := api.validateUserInput("s", "2.0", "a")

		assert.Nil(t, err)
	})
}
