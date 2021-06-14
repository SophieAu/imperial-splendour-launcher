package backend

import (
	"errors"
	"imperial-splendour-bundler/backend/customErrors"
	"imperial-splendour-bundler/backend/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testing"
)

func TestDownloadFiles(t *testing.T) {
	mockL := &mocks.MockLogger{}
	mockSt := &mocks.MockStore{}
	mockL.On("Warn", mock.Anything).Return()
	mockL.On("Info", mock.Anything).Return()
	mockL.On("Infof", mock.Anything).Return()
	mockSt.On("Update", mock.Anything).Return()

	baseFolder := "testfolder/"

	t.Run("Errors on the 2nd download", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("DownloadFile", setupUrl, mock.Anything).Return(errors.New("no download")).Once()
		mockS.On("DownloadFile", mock.Anything, mock.Anything).Return(nil)

		api := &API{logger: mockL, Sh: mockS, logStore: mockSt}
		api.setupBaseFolder = baseFolder + "/"
		err := api.downloadFiles()

		assert.Equal(t, customErrors.Download, err)
		mockS.AssertNumberOfCalls(t, "DownloadFile", 2)
	})

	t.Run("Downloads Everything", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("DownloadFile", mock.Anything, mock.Anything).Return(nil)

		api := &API{logger: mockL, Sh: mockS, logStore: mockSt}
		api.setupBaseFolder = baseFolder + "/"
		err := api.downloadFiles()

		assert.Nil(t, err)
		mockS.AssertNumberOfCalls(t, "DownloadFile", 4)
	})
}
