package backend

import (
	"errors"
	"imperial-splendour-bundler/backend/customErrors"
	"imperial-splendour-bundler/backend/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testing"
)

func TestCompilation(t *testing.T) {
	mockL := &mocks.MockLogger{}
	mockSt := &mocks.MockStore{}
	mockL.On("Warn", mock.Anything).Return()
	mockL.On("Info", mock.Anything).Return()
	mockL.On("Infof", mock.Anything).Return()

	baseFolder := "testfolder/"

	t.Run("Cannot compile for some reason", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("RunCommand", "iscc", []string{baseFolder + setupFile}).Return(errors.New("no runing InnoSetup")).Once()

		api := &API{logger: mockL, Sh: mockS, logStore: mockSt}
		api.setupBaseFolder = baseFolder
		err := api.compileSetup()

		assert.Equal(t, customErrors.CompileSetup, err)
	})

	t.Run("All Good", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("RunCommand", "iscc", []string{baseFolder + setupFile}).Return(nil).Once()

		api := &API{logger: mockL, Sh: mockS, logStore: mockSt}
		api.setupBaseFolder = baseFolder
		err := api.compileSetup()

		assert.Nil(t, err)
	})
}
