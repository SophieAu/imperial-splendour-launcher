package backend_test

import (
	"imperial-splendour-bundler/backend/test"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testing"
)

func TestPrepare(t *testing.T) {
	t.Run("Happy Path works great", func(t *testing.T) {
		sourcePath := "sourcePath"
		version := "3.7"
		fileListPath := "here/fileList.txt"

		expectedModPath := sourcePath + "/IS_Setup_Builder/ImperialSplendour/IS_Files/"

		api, _, _, _, _, mockS := test.Before()

		// Setup
		mockS.On("DoesFileExist", sourcePath+"/IS_Setup_Builder").Return(false, nil)
		mockS.On("MkdirAll", expectedModPath).Return(nil).Once()
		mockS.On("MkdirAll", sourcePath+"/IS_Setup_Builder/ImperialSplendour/IS_Uninstall/").Return(nil).Once()

		//File List
		mockS.On("GetDirContentByName", sourcePath).Return([]string{"File1.txt", "File2.txt", "F3.txt"}, nil).Once()
		mockS.On("ReadFile", "here/fileList.txt").Return([]byte("File1.txt\nFile2.txt"), nil).Once()
		mockS.On("MoveFile", sourcePath+"/File1.txt", expectedModPath+"File1.txt").Return(nil).Times(2)
		mockS.On("MoveFile", sourcePath+"/File2.txt", expectedModPath+"File2.txt").Return(nil).Times(2)
		mockS.On("WriteFile", expectedModPath+"IS_FileList.txt", []byte("File1.txt\nFile2.txt")).Return(nil).Once()

		// User Script
		mockS.On("DoesFileExist", sourcePath+"/user.empire_script.txt").Return(true, nil).Once()
		mockS.On("MoveFile", sourcePath+"/user.empire_script.txt", expectedModPath+"user.empire_script.txt").Return(nil).Once()

		// Download
		mockS.On("DownloadFile", mock.Anything, mock.Anything).Return(nil).Times(4)

		// Info file
		mockS.On("WriteFile", expectedModPath+"IS_Info.json", []byte("{\n\t\"isActive\": false,\n\t\"version\": \""+version+"\",\n\t\"usChecksum\": \"test\"\n}")).Return(nil).Once()

		// Version
		mockS.On("ReadFile", sourcePath+"/IS_Setup_Builder/setupBundled.iss").Return([]byte("stringy string !!!VERSION HERE!!!\nmerp derpderp"), nil).Once()
		mockS.On("WriteFile", sourcePath+"/IS_Setup_Builder/setupBundled.iss", []byte("stringy string 3.7\nmerp derpderp")).Return(nil)

		err := api.Prepare(sourcePath, version, false, fileListPath)
		assert.Nil(t, err)

		test.After(*api)
	})
}
