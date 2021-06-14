package backend

import (
	"errors"
	"imperial-splendour-bundler/backend/customErrors"
	"imperial-splendour-bundler/backend/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testing"
)

func TestPrepareModFiles(t *testing.T) {
	mockL := &mocks.MockLogger{}
	mockSt := &mocks.MockStore{}
	mockL.On("Warn", mock.Anything).Return()
	mockL.On("Info", mock.Anything).Return()
	mockL.On("Infof", mock.Anything).Return()
	mockSt.On("Update", mock.Anything).Return()

	baseFolder := "testfolder"

	t.Run("Cannot load source dir contents", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("GetDirContentByName", mock.Anything).Return([]string{}, errors.New("dir is evil")).Once()

		api := &API{logger: mockL, Sh: mockS, logStore: mockSt}
		api.setupBaseFolder = baseFolder + "/"
		err := api.prepareModFiles(baseFolder, "fileListHere")

		assert.Equal(t, customErrors.ReadSourceDir, err)
	})

	t.Run("Cannot read file list", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("GetDirContentByName", mock.Anything).Return([]string{"File1.txt", "File2.txt"}, nil).Once()
		mockS.On("ReadFile", "fileListHere").Return(nil, errors.New("Cannot read file"))

		api := &API{logger: mockL, Sh: mockS, logStore: mockSt}
		api.setupBaseFolder = baseFolder + "/"
		err := api.prepareModFiles(baseFolder, "fileListHere")

		assert.Equal(t, customErrors.ReadFileList, err)
	})

	t.Run("Compare is evil", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("GetDirContentByName", mock.Anything).Return([]string{"File1.txt", "File2.txt"}, nil).Once()
		mockS.On("ReadFile", "fileListHere").Return([]byte("File1.txt\nFile4.txt"), nil)

		api := &API{logger: mockL, Sh: mockS, logStore: mockSt}
		api.setupBaseFolder = baseFolder + "/"
		err := api.prepareModFiles(baseFolder, "fileListHere")

		assert.Equal(t, customErrors.FileMissing.Error()+" File4.txt", err.Error())
	})

	t.Run("Cannot move files", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("GetDirContentByName", mock.Anything).Return([]string{"File1.txt", "File2.txt", "F3.txt", "F4.txt"}, nil).Once()
		mockS.On("ReadFile", "fileListHere").Return([]byte("File1.txt\nFile2.txt\nF3.txt"), nil)
		mockS.On("MoveFile", "testfolder/File1.txt", mock.Anything).Return(nil)
		mockS.On("MoveFile", "testfolder/File2.txt", mock.Anything).Return(errors.New("not moving"))

		api := &API{logger: mockL, Sh: mockS, logStore: mockSt}
		api.setupBaseFolder = baseFolder + "/"
		err := api.prepareModFiles(baseFolder, "fileListHere")

		assert.Equal(t, customErrors.MoveFile.Error()+" File2.txt", err.Error())
		mockS.AssertCalled(t, "MoveFile", "testfolder/File1.txt", "testfolder/"+tempPath+modPath+"File1.txt")
		mockS.AssertNotCalled(t, "MoveFile", "testfolder/F3.txt", mock.Anything)
	})

	t.Run("Cannot save file list", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("GetDirContentByName", mock.Anything).Return([]string{"File1.txt", "File2.txt", "F3.txt"}, nil).Once()
		mockS.On("ReadFile", "fileListHere").Return([]byte("File1.txt\nFile2.txt"), nil)
		mockS.On("MoveFile", mock.Anything, mock.Anything).Return(nil)
		mockS.On("WriteFile", mock.Anything, mock.Anything).Return(errors.New("no writing"))

		api := &API{logger: mockL, Sh: mockS, logStore: mockSt}
		api.setupBaseFolder = baseFolder + "/"
		err := api.prepareModFiles(baseFolder, "fileListHere")

		assert.EqualError(t, err, customErrors.SaveFileList.Error())
	})

	t.Run("All good", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("GetDirContentByName", mock.Anything).Return([]string{"File1.txt", "File2.txt", "F3.txt"}, nil).Once()
		mockS.On("ReadFile", "fileListHere").Return([]byte("File1.txt\nFile2.txt"), nil)
		mockS.On("MoveFile", mock.Anything, mock.Anything).Return(nil)
		mockS.On("WriteFile", mock.Anything, []byte("File1.txt\nFile2.txt")).Return(nil)

		api := &API{logger: mockL, Sh: mockS, logStore: mockSt}
		api.setupBaseFolder = baseFolder + "/"
		err := api.prepareModFiles(baseFolder, "fileListHere")

		assert.Nil(t, err)
	})
}
