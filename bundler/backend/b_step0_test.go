package backend

import (
	"errors"
	"imperial-splendour-bundler/backend/customErrors"
	"imperial-splendour-bundler/backend/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testing"
)

func TestSetupBaseFolder(t *testing.T) {
	mockL := &mocks.MockLogger{}
	mockL.On("Warn", mock.Anything).Return()
	mockL.On("Info", mock.Anything).Return()

	t.Run("Folder exists already, as well as one fallback", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("DoesFileExist", "testfolder/"+preferredSetupBaseFolder).Return(true, nil).Once()
		mockS.On("DoesFileExist", "testfolder/"+preferredSetupBaseFolder+"_1").Return(true, nil).Once()
		mockS.On("DoesFileExist", mock.Anything).Return(false, nil)

		api := &API{logger: mockL, Sh: mockS}
		folder := api.getSetupBaseFolder("testfolder")

		assert.Equal(t, "testfolder/"+preferredSetupBaseFolder+"_2", folder)

	})

	t.Run("Folder exists already, no fallback though", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("DoesFileExist", "testfolder/"+preferredSetupBaseFolder).Return(true, nil).Once()
		mockS.On("DoesFileExist", mock.Anything).Return(false, nil)

		api := &API{logger: mockL, Sh: mockS}
		folder := api.getSetupBaseFolder("testfolder")

		assert.Equal(t, "testfolder/"+preferredSetupBaseFolder+"_1", folder)
	})

	t.Run("Folder doesn't exist", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("DoesFileExist", mock.Anything).Return(false, nil)

		api := &API{logger: mockL, Sh: mockS}
		folder := api.getSetupBaseFolder("testfolder")

		assert.Equal(t, "testfolder/"+preferredSetupBaseFolder, folder)
	})
}

func TestCreateTempFolder(t *testing.T) {
	mockL := &mocks.MockLogger{}
	mockL.On("Warn", mock.Anything).Return()
	mockL.On("Info", mock.Anything).Return()

	t.Run("Cannot create mod folder", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("MkdirAll", "folder/"+tempPath+modPath).Return(errors.New("nope")).Once()
		api := &API{logger: mockL, Sh: mockS}
		api.setupBaseFolder = "folder"

		err := api.createTempFolder()

		assert.Equal(t, err, customErrors.TempFolderCreation)
	})

	t.Run("Cannot create uninstall folder", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("MkdirAll", "folder/"+tempPath+modPath).Return(nil).Once()
		mockS.On("MkdirAll", "folder/"+tempPath+uninstallPath).Return(errors.New("nope")).Once()
		api := &API{logger: mockL, Sh: mockS}
		api.setupBaseFolder = "folder"

		err := api.createTempFolder()

		assert.Equal(t, err, customErrors.TempFolderCreation)
	})

	t.Run("Successfully creates all temp folders", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("MkdirAll", "folder/"+tempPath+modPath).Return(nil).Once()
		mockS.On("MkdirAll", "folder/"+tempPath+uninstallPath).Return(nil).Once()
		api := &API{logger: mockL, Sh: mockS}
		api.setupBaseFolder = "folder"

		err := api.createTempFolder()

		assert.Nil(t, err)
	})
}
