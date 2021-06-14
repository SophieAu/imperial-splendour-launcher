package backend

import (
	"errors"
	"imperial-splendour-bundler/backend/customErrors"
	"imperial-splendour-bundler/backend/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testing"
)

func TestCreateInfoJSON(t *testing.T) {
	mockL := &mocks.MockLogger{}
	mockSt := &mocks.MockStore{}
	mockL.On("Warn", mock.Anything).Return()
	mockL.On("Info", mock.Anything).Return()
	mockL.On("Infof", mock.Anything).Return()
	mockSt.On("Update", mock.Anything).Return()

	baseFolder := "testfolder"

	t.Run("Errors on the 2nd download", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("WriteFile", mock.Anything, mock.Anything).Return(errors.New("no write to json")).Once()

		api := &API{logger: mockL, Sh: mockS, logStore: mockSt}
		api.setupBaseFolder = baseFolder
		err := api.createInfoJSON("2.0")

		assert.Equal(t, customErrors.InfoFile, err)
		mockS.AssertCalled(t, "WriteFile", baseFolder+"/"+tempPath+modPath+infoFile, []byte("{\n\t\"isActive\": false,\n\t\"version\": \"2.0\",\n\t\"usChecksum\": \"test\"\n}"))
	})

	t.Run("Downloads Everything", func(t *testing.T) {
		mockS := &mocks.MockSystemHandler{}
		mockS.On("WriteFile", mock.Anything, mock.Anything).Return(nil)

		api := &API{logger: mockL, Sh: mockS, logStore: mockSt}
		api.setupBaseFolder = baseFolder
		err := api.createInfoJSON("3.7")

		assert.Nil(t, err)
		mockS.AssertCalled(t, "WriteFile", baseFolder+"/"+tempPath+modPath+infoFile, []byte("{\n\t\"isActive\": false,\n\t\"version\": \"3.7\",\n\t\"usChecksum\": \"test\"\n}"))
	})
}
