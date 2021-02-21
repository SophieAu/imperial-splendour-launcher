package backend_test

import (
	"errors"
	"imperial-splendour-launcher/backend"
	"imperial-splendour-launcher/backend/mock"

	"testing"

	"github.com/stretchr/testify/assert"
	testifyMock "github.com/stretchr/testify/mock"
)

func TestInit(t *testing.T) {
	mockL := &mock.Logger{}
	mockL.On("Infof", testifyMock.Anything, testifyMock.Anything).Return()

	var errorsOutWhenExecutableCannotBeRead = func() {
		mockS := &mock.SystemHandler{}
		mockB := &mock.Browser{}
		mockW := &mock.Window{}
		api := &backend.API{}

		mockS.On("Executable").Return("", errors.New("Could not find executable")).Once()

		err := api.Init(mockB, mockW, mockL, mockS)

		assert.Equal(t, "Could not find executable", err.Error())
		mockS.AssertNotCalled(t, "ReadFile", testifyMock.Anything)
	}

	// TODO: THIS ERROR SHOULD BE SEMI-RECOVERABLE
	var errorsOutWhenThereWereIssuesLoadingTheInfoFile = func() {
		mockS := &mock.SystemHandler{}
		mockB := &mock.Browser{}
		mockW := &mock.Window{}
		api := &backend.API{}

		mockS.On("Executable").Return(".", nil).Once()
		mockS.On("ReadFile", testifyMock.Anything).Return(nil, errors.New("FileNotFound")).Once()

		err := api.Init(mockB, mockW, mockL, mockS)

		mockL.AssertCalled(t, "Infof", expectFmt("ETW/Current directory: %s", "./"))
		mockL.AssertCalled(t, "Infof", expectFmt("AppData directory: %s", "./appDataFolder/The Creative Assembly/Empire/scripts/"))
		assert.Equal(t, "FileNotFound", err.Error())
	}

	var errorsOutWhenThereWereIssuesReadingTheInfoFile = func() {
		mockS := &mock.SystemHandler{}
		mockB := &mock.Browser{}
		mockW := &mock.Window{}
		api := &backend.API{}

		mockS.On("Executable").Return(".", nil).Once()
		mockS.On("ReadFile", testifyMock.Anything).Return([]byte{}, nil).Once()

		err := api.Init(mockB, mockW, mockL, mockS)

		mockL.AssertCalled(t, "Infof", expectFmt("ETW/Current directory: %s", "./"))
		mockL.AssertCalled(t, "Infof", expectFmt("AppData directory: %s", "./appDataFolder/The Creative Assembly/Empire/scripts/"))
		assert.NotNil(t, err)
	}

	var everythingWorks = func() {
		mockS := &mock.SystemHandler{}
		mockB := &mock.Browser{}
		mockW := &mock.Window{}
		api := &backend.API{}

		mockS.On("Executable").Return(".", nil).Once()
		mockS.On("ReadFile", testifyMock.Anything).Return([]byte("{\"isActive\": true, \"version\": \"2.0\"}"), nil).Once()

		err := api.Init(mockB, mockW, mockL, mockS)

		mockL.AssertCalled(t, "Infof", expectFmt("ETW/Current directory: %s", "./"))
		mockL.AssertCalled(t, "Infof", expectFmt("AppData directory: %s", "./appDataFolder/The Creative Assembly/Empire/scripts/"))
		assert.Nil(t, err)
	}

	errorsOutWhenExecutableCannotBeRead()
	errorsOutWhenThereWereIssuesLoadingTheInfoFile()
	errorsOutWhenThereWereIssuesReadingTheInfoFile()
	everythingWorks()
}
