package backend

import (
	"errors"
	"imperial-splendour-bundler/backend/customErrors"
	"imperial-splendour-bundler/backend/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testing"
)

func TestPrepareUserScript(t *testing.T) {
	mockL := &mocks.MockLogger{}
	mockSt := &mocks.MockStore{}
	mockL.On("Warn", mock.Anything).Return()
	mockL.On("Info", mock.Anything).Return()
	mockL.On("Infof", mock.Anything).Return()
	mockSt.On("Update", mock.Anything).Return()

	baseFolder := "testfolder"

	t.Run("User Script cannot be found", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("DoesFileExist", baseFolder+"/"+userScript).Return(false, errors.New("dir is evil")).Once()

		api := &API{logger: mockL, Sh: mockS, logStore: mockSt}
		api.setupBaseFolder = baseFolder + "/"
		err := api.prepareUserScript(baseFolder)

		assert.Equal(t, customErrors.UserScriptMissing, err)
	})

	t.Run("Cannot move userscript", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("DoesFileExist", baseFolder+"/"+userScript).Return(true, nil).Once()
		mockS.On("MoveFile", baseFolder+"/"+userScript, mock.Anything).Return(errors.New("Cannot move file"))

		api := &API{logger: mockL, Sh: mockS, logStore: mockSt}
		api.setupBaseFolder = baseFolder + "/"
		err := api.prepareUserScript(baseFolder)

		assert.Equal(t, customErrors.MoveUserScript, err)
	})

	t.Run("All good", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("DoesFileExist", baseFolder+"/"+userScript).Return(true, nil).Once()
		mockS.On("MoveFile", baseFolder+"/"+userScript, mock.Anything).Return(nil)

		api := &API{logger: mockL, Sh: mockS, logStore: mockSt}
		api.setupBaseFolder = baseFolder + "/"
		err := api.prepareUserScript(baseFolder)

		assert.Nil(t, err)
		mockS.AssertCalled(t, "MoveFile", baseFolder+"/"+userScript, baseFolder+"/"+tempPath+modPath+userScript)
	})
}
