package backend

import (
	"errors"
	"imperial-splendour-bundler/backend/customErrors"
	"imperial-splendour-bundler/backend/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testing"
)

func TestBundle(t *testing.T) {
	mockL := &mocks.MockLogger{}
	mockSt := &mocks.MockStore{}
	mockL.On("Warn", mock.Anything).Return()
	mockL.On("Info", mock.Anything).Return()
	mockL.On("Infof", mock.Anything).Return()
	mockSt.On("Update", mock.Anything).Return()

	sourcePath := "testfolder/"
	versionNumber := "12.0"
	fileListPath := "path/fileList.txt"

	prepApi := func(packageRawFiles bool, mockS Handler) *API {
		api := &API{logger: mockL, Sh: mockS, logStore: mockSt}
		api.setupBaseFolder = sourcePath + "test/"
		api.userSettings = UserSettings{sourcePath, versionNumber, packageRawFiles, fileListPath}
		return api
	}

	t.Run("Couldn't compile", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("RunCommand", "iscc", []string{sourcePath + "test/" + setupFile}).Return(errors.New("no runing InnoSetup")).Once()

		api := prepApi(false, mockS)
		err := api.Bundle()

		assert.Equal(t, customErrors.CompileSetup, err)
	})

	t.Run("No need to zip", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("RunCommand", "iscc", []string{sourcePath + "test/" + setupFile}).Return(nil).Once()

		api := prepApi(false, mockS)
		err := api.Bundle()

		assert.Nil(t, err)
		mockS.AssertNotCalled(t, "ReadFile", mock.Anything)
	})

	t.Run("Couldn't Zip Files", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("RunCommand", "iscc", []string{sourcePath + "test/" + setupFile}).Return(nil).Once()

		mockS.On("ReadFile", fileListPath).Return([]byte("File1.txt\nFile2.txt"), nil).Once()
		mockS.On("ZipFiles", mock.Anything, mock.Anything).Return(errors.New("No zippy")).Once()

		api := prepApi(true, mockS)
		err := api.Bundle()

		assert.Equal(t, customErrors.ZipFiles, err)
		mockS.AssertCalled(t, "ReadFile", fileListPath)
		mockS.AssertCalled(t, "ZipFiles", sourcePath+"IS-RotRv12.0.zip", []string{"File1.txt", "File2.txt", fileListFile, userScript})
	})

	t.Run("All good", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("RunCommand", "iscc", []string{sourcePath + "test/" + setupFile}).Return(nil).Once()

		mockS.On("ReadFile", fileListPath).Return([]byte("File1.txt\nFile2.txt"), nil).Once()
		mockS.On("ZipFiles", mock.Anything, mock.Anything).Return(nil).Once()

		api := prepApi(true, mockS)
		err := api.Bundle()

		assert.Nil(t, err)
		mockS.AssertCalled(t, "ReadFile", fileListPath)
		mockS.AssertCalled(t, "ZipFiles", sourcePath+"IS-RotRv12.0.zip", []string{"File1.txt", "File2.txt", fileListFile, userScript})
	})
}
