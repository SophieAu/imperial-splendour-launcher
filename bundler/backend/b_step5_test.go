package backend

import (
	"errors"
	"imperial-splendour-bundler/backend/customErrors"
	"imperial-splendour-bundler/backend/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testing"
)

func TestUpdatingVersionInSetupFile(t *testing.T) {
	mockL := &mocks.MockLogger{}
	mockSt := &mocks.MockStore{}
	mockL.On("Warn", mock.Anything).Return()
	mockL.On("Info", mock.Anything).Return()
	mockL.On("Infof", mock.Anything).Return()
	mockSt.On("Update", mock.Anything).Return()

	baseFolder := "testfolder"

	t.Run("Cannot read Script File", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("ReadFile", baseFolder+"/"+setupFile).Return([]byte("stringy string !!!VERSION HERE!!!"), errors.New("Cannot read")).Once()

		api := &API{logger: mockL, Sh: mockS, logStore: mockSt}
		api.setupBaseFolder = baseFolder
		err := api.updateSetupVersion("2.0")

		assert.Equal(t, customErrors.VersionUpdate, err)
		mockS.AssertCalled(t, "ReadFile", baseFolder+"/"+setupFile)
	})

	t.Run("Found no version strings ", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("ReadFile", baseFolder+"/"+setupFile).Return([]byte("stringy string derpderp"), nil).Once()

		api := &API{logger: mockL, Sh: mockS, logStore: mockSt}
		api.setupBaseFolder = baseFolder
		err := api.updateSetupVersion("2.0")

		assert.Equal(t, customErrors.VersionUpdate, err)
		mockS.AssertCalled(t, "ReadFile", baseFolder+"/"+setupFile)
	})

	t.Run("Replaced all version strings but cannot write", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("ReadFile", baseFolder+"/"+setupFile).Return([]byte("stringy string !!!VERSION HERE!!! merp !!!VERSION HERE!!! derpderp"), nil).Once()
		mockS.On("WriteFile", mock.Anything, mock.Anything).Return(errors.New("cannot write"))

		api := &API{logger: mockL, Sh: mockS, logStore: mockSt}
		api.setupBaseFolder = baseFolder
		err := api.updateSetupVersion("2.0")

		assert.Equal(t, customErrors.VersionUpdate, err)
		mockS.AssertCalled(t, "ReadFile", baseFolder+"/"+setupFile)
		mockS.AssertCalled(t, "WriteFile", baseFolder+"/"+setupFile, []byte("stringy string 2.0 merp 2.0 derpderp"))
	})

	t.Run("Replace the only version string and save file", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("ReadFile", baseFolder+"/"+setupFile).Return([]byte("stringy string !!!VERSION HERE!!!\nmerp derpderp"), nil).Once()
		mockS.On("WriteFile", mock.Anything, mock.Anything).Return(nil)

		api := &API{logger: mockL, Sh: mockS, logStore: mockSt}
		api.setupBaseFolder = baseFolder
		err := api.updateSetupVersion("3.7")

		assert.Nil(t, err)
		mockS.AssertCalled(t, "ReadFile", baseFolder+"/"+setupFile)
		mockS.AssertCalled(t, "WriteFile", baseFolder+"/"+setupFile, []byte("stringy string 3.7\nmerp derpderp"))
	})
}
