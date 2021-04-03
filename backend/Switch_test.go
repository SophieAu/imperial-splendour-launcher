package backend_test

import (
	"errors"
	"imperial-splendour-launcher/backend"
	"imperial-splendour-launcher/backend/mock"
	"imperial-splendour-launcher/backend/testHelpers"
	"strings"

	"testing"

	"github.com/stretchr/testify/assert"
	testifyMock "github.com/stretchr/testify/mock"
)

func activeBefore() (*backend.API, *mock.Browser, *mock.Window, *mock.Logger, *mock.SystemHandler) {
	return testHelpers.VariableBefore("2.0", true, "test")
}

func inactiveBefore() (*backend.API, *mock.Browser, *mock.Window, *mock.Logger, *mock.SystemHandler) {
	return testHelpers.VariableBefore("2.0", false, "test")
}

var list = []string{"dataFile.pack", "dataFile2.pack", "campaignTGA.tga", "campaignESF.esf", "campaignLUA.lua"}
var fileList = strings.Join(list, "\n")
var fileCount = len(list) + 1 // list length + user script

func TestDeactivate(t *testing.T) {
	t.Run("Cannot deactivate because file list file cannot be read", func(t *testing.T) {
		api, _, _, _, sysHandler := activeBefore()
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return(nil, errors.New("FileNotFound")).Once()

		err := api.Switch()

		assert.EqualError(t, err, "FileListError")
		sysHandler.AssertExpectations(t)
		sysHandler.AssertNotCalled(t, "MoveFile", testifyMock.Anything, testifyMock.Anything)
		assert.True(t, api.IsActive())

		testHelpers.After(*api)
	})

	t.Run("Cannot deactivate because file list contains unknown files", func(t *testing.T) {
		api, _, _, _, sysHandler := activeBefore()
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte(fileList+"\na\nb\n"), nil).Once()

		err := api.Switch()

		assert.EqualError(t, err, "FileListError")
		sysHandler.AssertExpectations(t)
		sysHandler.AssertNotCalled(t, "MoveFile", testifyMock.Anything, testifyMock.Anything)
		assert.True(t, api.IsActive())

		testHelpers.After(*api)
	})

	t.Run("Cannot deactivate because new status cannot be saved", func(t *testing.T) {
		api, _, _, _, sysHandler := activeBefore()
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte(fileList), nil).Once()
		sysHandler.On("WriteFile", "./IS_Files/IS_info.json", testifyMock.Anything).Return(errors.New("StatusUpdateError")).Once()
		sysHandler.On("MoveFile", testifyMock.Anything, testifyMock.Anything).Return(nil)

		err := api.Switch()

		assert.EqualError(t, err, "StatusUpdateError")
		sysHandler.AssertExpectations(t)
		sysHandler.AssertNumberOfCalls(t, "MoveFile", fileCount)
		sysHandler.AssertCalled(t, "WriteFile", "./IS_Files/IS_info.json", testHelpers.FmtInfoFile(false, "2.0", "test"))
		assert.True(t, api.IsActive())

		testHelpers.After(*api)
	})

	t.Run("Successfully deactivates Imperial Splendour despite an issue with moving files", func(t *testing.T) {
		api, _, _, logger, sysHandler := activeBefore()
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte(fileList), nil).Once()
		sysHandler.On("WriteFile", "./IS_Files/IS_info.json", testifyMock.Anything).Return(nil).Once()
		sysHandler.On("MoveFile", testifyMock.Anything, testifyMock.Anything).Return(errors.New("Cannot move file"))

		err := api.Switch()

		assert.EqualError(t, err, "DeactivationError")
		sysHandler.AssertExpectations(t)
		sysHandler.AssertCalled(t, "MoveFile", "./data/dataFile.pack", "./IS_Files/dataFile.pack")
		sysHandler.AssertCalled(t, "MoveFile", "./data/dataFile2.pack", "./IS_Files/dataFile2.pack")
		sysHandler.AssertCalled(t, "MoveFile", "./data/campaigns/imperial_splendour/campaignTGA.tga", "./IS_Files/campaignTGA.tga")
		sysHandler.AssertCalled(t, "MoveFile", "./data/campaigns/imperial_splendour/campaignESF.esf", "./IS_Files/campaignESF.esf")
		sysHandler.AssertCalled(t, "MoveFile", "./data/campaigns/imperial_splendour/campaignLUA.lua", "./IS_Files/campaignLUA.lua")
		sysHandler.AssertCalled(t, "MoveFile", "APPDATA/The Creative Assembly/Empire/scripts/user.empire_script.txt", "./IS_Files/user.empire_script.txt")
		logger.AssertNumberOfCalls(t, "Warnf", fileCount+1) // filecount times for moving files, 1x for endpoint return error
		assert.False(t, api.IsActive())

		testHelpers.After(*api)
	})

	t.Run("Successfully deactivates Imperial Splendour", func(t *testing.T) {
		api, _, _, _, sysHandler := activeBefore()
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte(fileList), nil).Once()
		sysHandler.On("WriteFile", "./IS_Files/IS_info.json", testifyMock.Anything).Return(nil).Once()
		sysHandler.On("MoveFile", testifyMock.Anything, testifyMock.Anything).Return(nil)

		err := api.Switch()

		assert.Nil(t, err)
		sysHandler.AssertExpectations(t)
		sysHandler.AssertNumberOfCalls(t, "MoveFile", fileCount)
		sysHandler.AssertCalled(t, "WriteFile", "./IS_Files/IS_info.json", testHelpers.FmtInfoFile(false, "2.0", "test"))
		assert.False(t, api.IsActive())

		testHelpers.After(*api)
	})
}

func TestActivate(t *testing.T) {
	t.Run("Cannot activate because file list file cannot be read", func(t *testing.T) {
		api, _, _, _, sysHandler := inactiveBefore()
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return(nil, errors.New("FileNotFound")).Once()

		err := api.Switch()

		assert.EqualError(t, err, "FileListError")
		sysHandler.AssertExpectations(t)
		sysHandler.AssertNotCalled(t, "MoveFile", testifyMock.Anything, testifyMock.Anything)
		assert.False(t, api.IsActive())

		testHelpers.After(*api)
	})

	t.Run("Cannot activate because file list contains unknown files", func(t *testing.T) {
		api, _, _, _, sysHandler := inactiveBefore()
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte(fileList+"\na\nb\n"), nil).Once()

		err := api.Switch()

		assert.EqualError(t, err, "FileListError")
		sysHandler.AssertExpectations(t)
		sysHandler.AssertNotCalled(t, "MoveFile", testifyMock.Anything, testifyMock.Anything)
		assert.False(t, api.IsActive())

		testHelpers.After(*api)
	})

	t.Run("Cancel and rollback without success if a data file cannot be moved", func(t *testing.T) {
		api, _, _, _, sysHandler := inactiveBefore()
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte(fileList), nil)
		sysHandler.On("MoveFile", "./IS_Files/dataFile2.pack", testifyMock.Anything).Return(errors.New("Couldn't move file")).Once()
		sysHandler.On("MoveFile", testifyMock.Anything, "./IS_Files/dataFile.pack").Return(errors.New("Random Error"))
		sysHandler.On("MoveFile", testifyMock.Anything, testifyMock.Anything).Return(nil)

		err := api.Switch()

		assert.EqualError(t, err, "RollbackError")
		sysHandler.AssertExpectations(t)
		sysHandler.AssertNumberOfCalls(t, "MoveFile", 3) // twice for activation move, once for rollback
		sysHandler.AssertCalled(t, "MoveFile", "./IS_Files/dataFile.pack", "./data/dataFile.pack")
		sysHandler.AssertCalled(t, "MoveFile", "./IS_Files/dataFile2.pack", "./data/dataFile2.pack")
		sysHandler.AssertCalled(t, "MoveFile", "./data/dataFile.pack", "./IS_Files/dataFile.pack")
		assert.False(t, api.IsActive())

		testHelpers.After(*api)
	})

	t.Run("Cancel and rollback successfully if a data file cannot be moved", func(t *testing.T) {
		api, _, _, _, sysHandler := inactiveBefore()
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte(fileList), nil)
		sysHandler.On("MoveFile", "./IS_Files/dataFile2.pack", testifyMock.Anything).Return(errors.New("Couldn't move file")).Once()
		sysHandler.On("MoveFile", testifyMock.Anything, testifyMock.Anything).Return(nil)

		err := api.Switch()

		assert.EqualError(t, err, "ActivationError")
		sysHandler.AssertExpectations(t)
		sysHandler.AssertNumberOfCalls(t, "MoveFile", 3) // twice for activation move, once for rollback
		sysHandler.AssertCalled(t, "MoveFile", "./IS_Files/dataFile.pack", "./data/dataFile.pack")
		sysHandler.AssertCalled(t, "MoveFile", "./IS_Files/dataFile2.pack", "./data/dataFile2.pack")
		sysHandler.AssertCalled(t, "MoveFile", "./data/dataFile.pack", "./IS_Files/dataFile.pack")
		assert.False(t, api.IsActive())

		testHelpers.After(*api)
	})

	t.Run("Cancel and rollback successfully if a campaign file cannot be moved", func(t *testing.T) {
		api, _, _, _, sysHandler := inactiveBefore()
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte(fileList), nil)
		sysHandler.On("MoveFile", "./IS_Files/campaignTGA.tga", testifyMock.Anything).Return(errors.New("Couldn't move file")).Once()
		sysHandler.On("MoveFile", testifyMock.Anything, testifyMock.Anything).Return(nil)

		err := api.Switch()

		assert.EqualError(t, err, "ActivationError")
		sysHandler.AssertExpectations(t)
		sysHandler.AssertNumberOfCalls(t, "MoveFile", 5) // 3x for activation move, 2x for undo
		sysHandler.AssertCalled(t, "MoveFile", "./IS_Files/dataFile.pack", "./data/dataFile.pack")
		sysHandler.AssertCalled(t, "MoveFile", "./IS_Files/dataFile2.pack", "./data/dataFile2.pack")
		sysHandler.AssertCalled(t, "MoveFile", "./IS_Files/campaignTGA.tga", "./data/campaigns/imperial_splendour/campaignTGA.tga")
		sysHandler.AssertCalled(t, "MoveFile", testifyMock.Anything, "./IS_Files/dataFile.pack")
		sysHandler.AssertCalled(t, "MoveFile", testifyMock.Anything, "./IS_Files/dataFile2.pack")
		assert.False(t, api.IsActive())

		testHelpers.After(*api)
	})

	t.Run("Cancel and rollback successfully if the user script cannot be moved", func(t *testing.T) {
		api, _, _, _, sysHandler := inactiveBefore()
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte(fileList), nil)
		sysHandler.On("MoveFile", "./IS_Files/user.empire_script.txt", testifyMock.Anything).Return(errors.New("Couldn't move file"))
		sysHandler.On("MoveFile", testifyMock.Anything, testifyMock.Anything).Return(nil)

		err := api.Switch()

		assert.EqualError(t, err, "ActivationError")
		sysHandler.AssertExpectations(t)
		sysHandler.AssertNumberOfCalls(t, "MoveFile", fileCount+fileCount-1) // 1x activation move, fileCount-1 for undo
		assert.False(t, api.IsActive())

		testHelpers.After(*api)
	})

	t.Run("Cancel and rollback successfully if status file cannot be updated", func(t *testing.T) {
		api, _, _, _, sysHandler := inactiveBefore()
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte(fileList), nil)
		sysHandler.On("WriteFile", "./IS_Files/IS_info.json", testifyMock.Anything).Return(errors.New("StatusUpdateError")).Once()
		sysHandler.On("MoveFile", testifyMock.Anything, testifyMock.Anything).Return(nil)

		err := api.Switch()

		assert.EqualError(t, err, "StatusUpdateError")
		sysHandler.AssertExpectations(t)
		sysHandler.AssertNumberOfCalls(t, "MoveFile", 2*fileCount) // twice for 1x activation and 1x undo
		assert.False(t, api.IsActive())

		testHelpers.After(*api)
	})

	t.Run("Cancel and error out on rollback if status file cannot be updated", func(t *testing.T) {
		api, _, _, _, sysHandler := inactiveBefore()
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte(fileList), nil)
		sysHandler.On("WriteFile", "./IS_Files/IS_info.json", testifyMock.Anything).Return(errors.New("StatusUpdateError")).Once()
		sysHandler.On("MoveFile", testifyMock.Anything, "./IS_Files/user.empire_script.txt").Return(errors.New("Couldn't move file")).Once()
		sysHandler.On("MoveFile", testifyMock.Anything, testifyMock.Anything).Return(nil)

		err := api.Switch()

		assert.EqualError(t, err, "RollbackError")
		sysHandler.AssertExpectations(t)
		sysHandler.AssertNumberOfCalls(t, "MoveFile", 2*fileCount) // twice for 1x activation and 1x undo
		assert.False(t, api.IsActive())

		testHelpers.After(*api)
	})

	t.Run("Successfully activate Imperial Splendour", func(t *testing.T) {
		api, _, _, _, sysHandler := inactiveBefore()
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte(fileList), nil)
		sysHandler.On("WriteFile", testifyMock.Anything, testifyMock.Anything).Return(nil).Once()
		sysHandler.On("MoveFile", testifyMock.Anything, testifyMock.Anything).Return(nil)

		err := api.Switch()

		assert.Nil(t, err)
		sysHandler.AssertExpectations(t)
		sysHandler.AssertNumberOfCalls(t, "MoveFile", fileCount) // twice for activation move, fileCount times for undo
		sysHandler.AssertCalled(t, "WriteFile", "./IS_Files/IS_info.json", testHelpers.FmtInfoFile(true, "2.0", "test"))
		assert.True(t, api.IsActive())

		testHelpers.After(*api)
	})
}
