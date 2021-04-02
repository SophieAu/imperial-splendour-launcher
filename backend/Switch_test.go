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

/*
ADDITIONAL TEST CASES ON SWITCH

* error cases: cannot rename file (-> copy paste, then delete)
* file doesn't exist (-> check for it in IS_Files, error maybe)
* userscript wrong checksum (-> panic)
* is info file doesn't exist (-> panic)
* couldn't write to is info file (-> panic)

*/

func activeBefore() (*backend.API, *mock.Browser, *mock.Window, *mock.Logger, *mock.SystemHandler) {
	return testHelpers.VariableBefore("2.0", true, "test")
}

func inactiveBefore() (*backend.API, *mock.Browser, *mock.Window, *mock.Logger, *mock.SystemHandler) {
	return testHelpers.VariableBefore("2.0", false, "test")
}

const fileCount = 6
const fileList = "dataFile.pack\ndataFile2.pack\ncampaignTGA.tga\ncampaignESF.esf\ncampaignLUA.lua"

func TestDeactivate(t *testing.T) {
	t.Run("Cannot deactivate because file list file cannot be read", func(t *testing.T) {
		api, _, _, _, sysHandler := activeBefore()
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return(nil, errors.New("FileNotFound")).Once()

		err := api.Switch()

		assert.EqualError(t, err, "FileNotFound")
		sysHandler.AssertNotCalled(t, "MoveFile", testifyMock.Anything, testifyMock.Anything)
		assert.True(t, api.IsActive())

		testHelpers.After(*api)
	})

	t.Run("Cannot deactivate because file list contains unknown files", func(t *testing.T) {
		api, _, _, _, sysHandler := activeBefore()
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte(fileList+"\na\nb\n"), nil).Once()

		err := api.Switch()

		assert.EqualError(t, err, "Unknown file 'a' found in file list\nUnknown file 'b' found in file list\n")
		sysHandler.AssertNotCalled(t, "MoveFile", testifyMock.Anything, testifyMock.Anything)
		assert.True(t, api.IsActive())

		testHelpers.After(*api)
	})

	t.Run("Cannot deactivate because new status cannot be saved", func(t *testing.T) {
		api, _, _, _, sysHandler := activeBefore()
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte(fileList), nil).Once()
		sysHandler.On("WriteFile", "./IS_Files/IS_info.json", testifyMock.Anything).Return(errors.New("Cannot update Status")).Once()
		sysHandler.On("MoveFile", testifyMock.Anything, testifyMock.Anything).Return(nil)

		err := api.Switch()

		assert.EqualError(t, err, "Cannot update Status")
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

		assert.Nil(t, err)
		sysHandler.AssertCalled(t, "MoveFile", "./data/dataFile.pack", "./IS_Files/dataFile.pack")
		sysHandler.AssertCalled(t, "MoveFile", "./data/dataFile2.pack", "./IS_Files/dataFile2.pack")
		sysHandler.AssertCalled(t, "MoveFile", "./data/campaigns/imperial_splendour/campaignTGA.tga", "./IS_Files/campaignTGA.tga")
		sysHandler.AssertCalled(t, "MoveFile", "./data/campaigns/imperial_splendour/campaignESF.esf", "./IS_Files/campaignESF.esf")
		sysHandler.AssertCalled(t, "MoveFile", "./data/campaigns/imperial_splendour/campaignLUA.lua", "./IS_Files/campaignLUA.lua")
		sysHandler.AssertCalled(t, "MoveFile", "APPDATA/The Creative Assembly/Empire/scripts/user.empire_script.txt", "./IS_Files/user.empire_script.txt")
		logger.AssertNumberOfCalls(t, "Warnf", fileCount)
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

		assert.EqualError(t, err, "FileNotFound")
		sysHandler.AssertNotCalled(t, "MoveFile", testifyMock.Anything, testifyMock.Anything)
		assert.False(t, api.IsActive())

		testHelpers.After(*api)
	})

	t.Run("Cannot activate because file list contains unknown files", func(t *testing.T) {
		api, _, _, _, sysHandler := inactiveBefore()
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte(fileList+"\na\nb\n"), nil).Once()

		err := api.Switch()

		assert.EqualError(t, err, "Unknown file 'a' found in file list\nUnknown file 'b' found in file list\n")
		sysHandler.AssertNotCalled(t, "MoveFile", testifyMock.Anything, testifyMock.Anything)
		assert.False(t, api.IsActive())

		testHelpers.After(*api)
	})

	// TestRollbackOnErrorWhileActivating := func(t *testing.T) {
	// 	fileList := "merp.pack\nderp.pack\nx.tga\ny.esf\nz.lua"

	// 	api, _, _, _, sysHandler := testHelpers.VariableBefore("2.0", false, "test")
	// 	sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte(fileList), nil).Once()
	// 	sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte(fileList), nil).Once()
	// 	sysHandler.On("MoveFile", "./IS_Files/y.esf", "./data/campaigns/imperial_splendour/y.esf").Return(errors.New("FileNotFound"))
	// 	sysHandler.On("MoveFile", testifyMock.Anything, testifyMock.Anything).Return(nil)
	// 	sysHandler.On("WriteFile", testifyMock.Anything, testifyMock.Anything).Return(nil)

	// 	err := api.Switch()

	// 	assert.NotNil(t, err)
	// 	sysHandler.AssertCalled(t, "MoveFile", "./IS_Files/merp.pack", "./data/merp.pack")
	// 	sysHandler.AssertCalled(t, "MoveFile", "./IS_Files/derp.pack", "./data/derp.pack")
	// 	sysHandler.AssertCalled(t, "MoveFile", "./IS_Files/x.tga", "./data/campaigns/imperial_splendour/x.tga")
	// 	sysHandler.AssertCalled(t, "MoveFile", "./IS_Files/y.esf", "./data/campaigns/imperial_splendour/y.esf")

	// 	sysHandler.AssertCalled(t, "MoveFile", "./data/merp.pack", "./IS_Files/merp.pack")
	// 	sysHandler.AssertCalled(t, "MoveFile", "./data/derp.pack", "./IS_Files/derp.pack")
	// 	sysHandler.AssertCalled(t, "MoveFile", "./data/campaigns/imperial_splendour/x.tga", "./IS_Files/x.tga")
	// 	sysHandler.AssertCalled(t, "MoveFile", "./data/campaigns/imperial_splendour/y.esf", "./IS_Files/y.esf")
	// 	sysHandler.AssertCalled(t, "MoveFile", "./data/campaigns/imperial_splendour/z.lua", "./IS_Files/z.lua")
	// 	sysHandler.AssertCalled(t, "MoveFile", "APPDATA/The Creative Assembly/Empire/scripts/user.empire_script.txt", "./IS_Files/user.empire_script.txt")
	// 	sysHandler.AssertCalled(t, "WriteFile", "./IS_Files/IS_info.json", testHelpers.FmtInfoFile(false, "2.0", "test"))

	// 	sysHandler.AssertNotCalled(t, "MoveFile", "./IS_Files/z.lua", "./data/campaigns/imperial_splendour/z.lua")
	// 	sysHandler.AssertNotCalled(t, "MoveFile", "./IS_Files/user.empire_script.txt", "APPDATA/The Creative Assembly/Empire/scripts/user.empire_script.txt")
	// 	sysHandler.AssertNotCalled(t, "WriteFile", "./IS_Files/IS_info.json", testHelpers.FmtInfoFile(true, "2.0", "test"))

	// 	assert.False(t, api.IsActive())

	// 	testHelpers.After(*api)
	// }
	// TestRollbackOnErrorWhileActivating(t)

	// TestSuccessfullyActivateImpSplen := func(t *testing.T) {
	// 	fileList := "merp.pack\nderp.pack\nx.tga\ny.esf\nz.lua"

	// 	api, _, _, _, sysHandler := testHelpers.VariableBefore("2.0", false, "test")
	// 	sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte(fileList), nil).Once()
	// 	sysHandler.On("MoveFile", testifyMock.Anything, testifyMock.Anything).Return(nil)
	// 	sysHandler.On("WriteFile", testifyMock.Anything, testifyMock.Anything).Return(nil)

	// 	err := api.Switch()

	// 	/*
	// 		error cases: cannot rename file (-> copy paste, then delete)
	// 		file doesn't exist (-> check for it in IS_Files, error maybe)
	// 		userscript wrong checksum (-> panic)
	// 		is info file doesn't exist (-> panic)
	// 		couldn't write to is info file (-> panic)
	// 	*/

	// 	assert.Nil(t, err)
	// 	sysHandler.AssertCalled(t, "MoveFile", "./IS_Files/merp.pack", "./data/merp.pack")
	// 	sysHandler.AssertCalled(t, "MoveFile", "./IS_Files/derp.pack", "./data/derp.pack")
	// 	sysHandler.AssertCalled(t, "MoveFile", "./IS_Files/x.tga", "./data/campaigns/imperial_splendour/x.tga")
	// 	sysHandler.AssertCalled(t, "MoveFile", "./IS_Files/y.esf", "./data/campaigns/imperial_splendour/y.esf")
	// 	sysHandler.AssertCalled(t, "MoveFile", "./IS_Files/z.lua", "./data/campaigns/imperial_splendour/z.lua")
	// 	sysHandler.AssertCalled(t, "MoveFile", "./IS_Files/user.empire_script.txt", "APPDATA/The Creative Assembly/Empire/scripts/user.empire_script.txt")
	// 	sysHandler.AssertCalled(t, "WriteFile", "./IS_Files/IS_info.json", testHelpers.FmtInfoFile(true, "2.0", "test"))
	// 	assert.True(t, api.IsActive())

	// 	testHelpers.After(*api)
	// }
	// TestSuccessfullyActivateImpSplen(t)
}
