package backend_test

import (
	"errors"
	"imperial-splendour-launcher/backend"
	"imperial-splendour-launcher/backend/mock"
	"imperial-splendour-launcher/backend/testHelpers"

	"testing"

	"github.com/stretchr/testify/assert"
	testifyMock "github.com/stretchr/testify/mock"
)

func TestInit(t *testing.T) {
	mockL := &mock.Logger{}
	mockL.On("Infof", testifyMock.Anything, testifyMock.Anything).Return()
	mockL.On("Warnf", testifyMock.Anything, testifyMock.Anything).Return()

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
		mockS.On("Getenv", "APPDATA").Return("APPDATA")
		mockS.On("ReadFile", testifyMock.Anything).Return(nil, errors.New("FileNotFound")).Once()

		err := api.Init(mockB, mockW, mockL, mockS)

		mockL.AssertCalled(t, "Infof", testHelpers.ExpectFmt("ETW/Current directory: %s", "./"))
		mockL.AssertCalled(t, "Infof", testHelpers.ExpectFmt("AppData directory: %s", "APPDATA/The Creative Assembly/Empire/scripts/"))
		assert.Equal(t, "FileNotFound", err.Error())
	}

	var errorsOutWhenThereWereIssuesReadingTheInfoFile = func() {
		mockS := &mock.SystemHandler{}
		mockB := &mock.Browser{}
		mockW := &mock.Window{}
		api := &backend.API{}

		mockS.On("Executable").Return(".", nil).Once()
		mockS.On("Getenv", "APPDATA").Return("APPDATA")
		mockS.On("ReadFile", testifyMock.Anything).Return([]byte{}, nil).Once()

		err := api.Init(mockB, mockW, mockL, mockS)

		mockL.AssertCalled(t, "Infof", testHelpers.ExpectFmt("ETW/Current directory: %s", "./"))
		mockL.AssertCalled(t, "Infof", testHelpers.ExpectFmt("AppData directory: %s", "APPDATA/The Creative Assembly/Empire/scripts/"))
		assert.Equal(t, "unexpected end of JSON input", err.Error())
	}

	var corruptedVersion = func() {
		mockS := &mock.SystemHandler{}
		mockB := &mock.Browser{}
		mockW := &mock.Window{}
		api := &backend.API{}

		mockS.On("Executable").Return(".", nil).Once()
		mockS.On("Getenv", "APPDATA").Return("APPDATA")
		mockS.On("ReadFile", testifyMock.Anything).Return([]byte("{\"isActive\": true, \"version\": \"\", \"usChecksum\": \"test\"}"), nil).Once()

		err := api.Init(mockB, mockW, mockL, mockS)

		mockL.AssertCalled(t, "Infof", testHelpers.ExpectFmt("ETW/Current directory: %s", "./"))
		mockL.AssertCalled(t, "Infof", testHelpers.ExpectFmt("AppData directory: %s", "APPDATA/The Creative Assembly/Empire/scripts/"))
		assert.Equal(t, "Corrupt Info File", err.Error())
	}

	var corruptedChecksum = func() {
		mockS := &mock.SystemHandler{}
		mockB := &mock.Browser{}
		mockW := &mock.Window{}
		api := &backend.API{}

		mockS.On("Executable").Return(".", nil).Once()
		mockS.On("Getenv", "APPDATA").Return("APPDATA")
		mockS.On("ReadFile", testifyMock.Anything).Return([]byte("{\"isActive\": true, \"version\": \"2.0\", \"usChecksum\": \"\"}"), nil).Once()

		err := api.Init(mockB, mockW, mockL, mockS)

		mockL.AssertCalled(t, "Infof", testHelpers.ExpectFmt("ETW/Current directory: %s", "./"))
		mockL.AssertCalled(t, "Infof", testHelpers.ExpectFmt("AppData directory: %s", "APPDATA/The Creative Assembly/Empire/scripts/"))
		assert.Equal(t, "Corrupt Info File", err.Error())
	}

	var everythingWorks = func() {
		mockS := &mock.SystemHandler{}
		mockB := &mock.Browser{}
		mockW := &mock.Window{}
		api := &backend.API{}

		mockS.On("Executable").Return(".", nil).Once()
		mockS.On("Getenv", "APPDATA").Return("APPDATA")
		mockS.On("ReadFile", testifyMock.Anything).Return([]byte("{\"isActive\": true, \"version\": \"2.0\", \"usChecksum\": \"test\"}"), nil).Once()

		err := api.Init(mockB, mockW, mockL, mockS)

		mockL.AssertCalled(t, "Infof", testHelpers.ExpectFmt("ETW/Current directory: %s", "./"))
		mockL.AssertCalled(t, "Infof", testHelpers.ExpectFmt("AppData directory: %s", "APPDATA/The Creative Assembly/Empire/scripts/"))
		assert.Nil(t, err)
	}

	errorsOutWhenExecutableCannotBeRead()
	errorsOutWhenThereWereIssuesLoadingTheInfoFile()
	errorsOutWhenThereWereIssuesReadingTheInfoFile()
	corruptedVersion()
	corruptedChecksum()
	everythingWorks()
}
