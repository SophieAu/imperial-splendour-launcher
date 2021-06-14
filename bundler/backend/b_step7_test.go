package backend

import (
	"errors"
	"imperial-splendour-bundler/backend/customErrors"
	"imperial-splendour-bundler/backend/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testing"
)

func TestZipBundling(t *testing.T) {
	mockL := &mocks.MockLogger{}
	mockSt := &mocks.MockStore{}
	mockL.On("Warn", mock.Anything).Return()
	mockL.On("Info", mock.Anything).Return()
	mockL.On("Infof", mock.Anything).Return()
	mockSt.On("Update", mock.Anything).Return()

	sourcePath := "testfolder/"
	fileList := "path/fileList.txt"

	t.Run("Couldn't read FileList", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("ReadFile", mock.Anything).Return([]byte("File1.txt\nFile2.txt"), errors.New("error")).Once()

		api := &API{logger: mockL, Sh: mockS, logStore: mockSt}
		err := api.buildZipBundle(sourcePath, "2.3", fileList)

		assert.Equal(t, customErrors.ZipFiles, err)
		mockS.AssertCalled(t, "ReadFile", fileList)
	})

	t.Run("Couldn't Zip Files", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("ReadFile", fileList).Return([]byte("File1.txt\nFile2.txt"), nil).Once()
		mockS.On("ZipFiles", mock.Anything, mock.Anything).Return(errors.New("No zippy")).Once()

		api := &API{logger: mockL, Sh: mockS, logStore: mockSt}
		err := api.buildZipBundle(sourcePath, "2.3", fileList)

		assert.Equal(t, customErrors.ZipFiles, err)
		mockS.AssertCalled(t, "ReadFile", fileList)
		mockS.AssertCalled(t, "ZipFiles", sourcePath+"IS-RotRv2.3.zip", []string{"File1.txt", "File2.txt", fileListFile, userScript})
	})

	t.Run("All good", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("ReadFile", fileList).Return([]byte("File1.txt\nFile2.txt"), nil).Once()
		mockS.On("ZipFiles", mock.Anything, mock.Anything).Return(nil).Once()

		api := &API{logger: mockL, Sh: mockS, logStore: mockSt}
		err := api.buildZipBundle(sourcePath, "4.18", fileList)

		assert.Nil(t, err)
		mockS.AssertCalled(t, "ReadFile", fileList)
		mockS.AssertCalled(t, "ZipFiles", sourcePath+"IS-RotRv4.18.zip", []string{"File1.txt", "File2.txt", fileListFile, userScript})
	})
}
