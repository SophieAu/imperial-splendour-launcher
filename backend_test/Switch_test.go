package backend_test

import (
	"errors"

	"github.com/stretchr/testify/assert"
	testifyMock "github.com/stretchr/testify/mock"

	"testing"
)

/*
ADDITIONAL TEST CASES ON SWITCH

* error cases: cannot rename file (-> copy paste, then delete)
* file doesn't exist (-> check for it in IS_Files, error maybe)
* userscript wrong checksum (-> panic)
* is info file doesn't exist (-> panic)
* couldn't write to is info file (-> panic)

*/

func TestSwitch(t *testing.T) {

	TestFileListNotFound := func(t *testing.T) {
		api, _, _, _, sysHandler := variableBefore("2.0", true, "test")
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return(nil, errors.New("FileNotFound")).Once()

		err := api.Switch()

		assert.EqualError(t, err, "FileNotFound")
		sysHandler.AssertNotCalled(t, "MoveFile", testifyMock.Anything, testifyMock.Anything)
		assert.True(t, api.IsActive())

		after(*api)
	}
	TestFileListNotFound(t)

	TestUnknownFileInFileList := func(t *testing.T) {
		api, _, _, _, sysHandler := variableBefore("2.0", true, "test")
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte("a\nb\n"), nil).Once()

		err := api.Switch()

		assert.EqualError(t, err, "Unknown file 'a' found in file list\nUnknown file 'b' found in file list\n")
		sysHandler.AssertNotCalled(t, "MoveFile", testifyMock.Anything, testifyMock.Anything)
		assert.True(t, api.IsActive())

		after(*api)
	}
	TestUnknownFileInFileList(t)

	TestSwallowErrorsWhenDeactivating := func(t *testing.T) {
		fileList := "merp.pack\nderp.pack\nx.tga\ny.esf\nz.lua"

		api, _, _, _, sysHandler := variableBefore("2.0", true, "test")
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte(fileList), nil).Once()
		sysHandler.On("MoveFile", "./data/campaigns/imperial_splendour/y.esf", "./IS_Files/y.esf").Return(errors.New("FileNotFound"))
		sysHandler.On("MoveFile", testifyMock.Anything, testifyMock.Anything).Return(nil)
		sysHandler.On("WriteFile", "./IS_Files/IS_info.json", fmtInfoFile(false, "2.0", "test")).Return(nil)

		err := api.Switch()

		assert.Nil(t, err)
		sysHandler.AssertCalled(t, "MoveFile", "./data/merp.pack", "./IS_Files/merp.pack")
		sysHandler.AssertCalled(t, "MoveFile", "./data/derp.pack", "./IS_Files/derp.pack")
		sysHandler.AssertCalled(t, "MoveFile", "./data/campaigns/imperial_splendour/x.tga", "./IS_Files/x.tga")
		sysHandler.AssertCalled(t, "MoveFile", "./data/campaigns/imperial_splendour/y.esf", "./IS_Files/y.esf")
		sysHandler.AssertCalled(t, "MoveFile", "./data/campaigns/imperial_splendour/z.lua", "./IS_Files/z.lua")
		sysHandler.AssertCalled(t, "MoveFile", "APPDATA/The Creative Assembly/Empire/scripts/user.empire_script.txt", "./IS_Files/user.empire_script.txt")
		sysHandler.AssertCalled(t, "WriteFile", "./IS_Files/IS_info.json", fmtInfoFile(false, "2.0", "test"))
		assert.False(t, api.IsActive())

		after(*api)
	}
	TestSwallowErrorsWhenDeactivating(t)

	TestSuccessfullyDeactivateImpSplen := func(t *testing.T) {
		fileList := "merp.pack\nderp.pack\nx.tga\ny.esf\nz.lua"

		api, _, _, _, sysHandler := variableBefore("2.0", true, "test")
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte(fileList), nil).Once()
		sysHandler.On("MoveFile", testifyMock.Anything, testifyMock.Anything).Return(nil)
		sysHandler.On("WriteFile", "./IS_Files/IS_info.json", fmtInfoFile(false, "2.0", "test")).Return(nil)

		err := api.Switch()

		assert.Nil(t, err)
		sysHandler.AssertCalled(t, "MoveFile", "./data/merp.pack", "./IS_Files/merp.pack")
		sysHandler.AssertCalled(t, "MoveFile", "./data/derp.pack", "./IS_Files/derp.pack")
		sysHandler.AssertCalled(t, "MoveFile", "./data/campaigns/imperial_splendour/x.tga", "./IS_Files/x.tga")
		sysHandler.AssertCalled(t, "MoveFile", "./data/campaigns/imperial_splendour/y.esf", "./IS_Files/y.esf")
		sysHandler.AssertCalled(t, "MoveFile", "./data/campaigns/imperial_splendour/z.lua", "./IS_Files/z.lua")
		sysHandler.AssertCalled(t, "MoveFile", "APPDATA/The Creative Assembly/Empire/scripts/user.empire_script.txt", "./IS_Files/user.empire_script.txt")
		sysHandler.AssertCalled(t, "WriteFile", "./IS_Files/IS_info.json", fmtInfoFile(false, "2.0", "test"))
		assert.False(t, api.IsActive())

		after(*api)
	}
	TestSuccessfullyDeactivateImpSplen(t)

	TestRollbackOnErrorWhileActivating := func(t *testing.T) {
		fileList := "merp.pack\nderp.pack\nx.tga\ny.esf\nz.lua"

		api, _, _, _, sysHandler := variableBefore("2.0", false, "test")
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte(fileList), nil).Once()
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte(fileList), nil).Once()
		sysHandler.On("MoveFile", "./IS_Files/y.esf", "./data/campaigns/imperial_splendour/y.esf").Return(errors.New("FileNotFound"))
		sysHandler.On("MoveFile", testifyMock.Anything, testifyMock.Anything).Return(nil)
		sysHandler.On("WriteFile", testifyMock.Anything, testifyMock.Anything).Return(nil)

		err := api.Switch()

		assert.NotNil(t, err)
		sysHandler.AssertCalled(t, "MoveFile", "./IS_Files/merp.pack", "./data/merp.pack")
		sysHandler.AssertCalled(t, "MoveFile", "./IS_Files/derp.pack", "./data/derp.pack")
		sysHandler.AssertCalled(t, "MoveFile", "./IS_Files/x.tga", "./data/campaigns/imperial_splendour/x.tga")
		sysHandler.AssertCalled(t, "MoveFile", "./IS_Files/y.esf", "./data/campaigns/imperial_splendour/y.esf")

		sysHandler.AssertCalled(t, "MoveFile", "./data/merp.pack", "./IS_Files/merp.pack")
		sysHandler.AssertCalled(t, "MoveFile", "./data/derp.pack", "./IS_Files/derp.pack")
		sysHandler.AssertCalled(t, "MoveFile", "./data/campaigns/imperial_splendour/x.tga", "./IS_Files/x.tga")
		sysHandler.AssertCalled(t, "MoveFile", "./data/campaigns/imperial_splendour/y.esf", "./IS_Files/y.esf")
		sysHandler.AssertCalled(t, "MoveFile", "./data/campaigns/imperial_splendour/z.lua", "./IS_Files/z.lua")
		sysHandler.AssertCalled(t, "MoveFile", "APPDATA/The Creative Assembly/Empire/scripts/user.empire_script.txt", "./IS_Files/user.empire_script.txt")
		sysHandler.AssertCalled(t, "WriteFile", "./IS_Files/IS_info.json", fmtInfoFile(false, "2.0", "test"))

		sysHandler.AssertNotCalled(t, "MoveFile", "./IS_Files/z.lua", "./data/campaigns/imperial_splendour/z.lua")
		sysHandler.AssertNotCalled(t, "MoveFile", "./IS_Files/user.empire_script.txt", "APPDATA/The Creative Assembly/Empire/scripts/user.empire_script.txt")
		sysHandler.AssertNotCalled(t, "WriteFile", "./IS_Files/IS_info.json", fmtInfoFile(true, "2.0", "test"))

		assert.False(t, api.IsActive())

		after(*api)
	}
	TestRollbackOnErrorWhileActivating(t)

	TestSuccessfullyActivateImpSplen := func(t *testing.T) {
		fileList := "merp.pack\nderp.pack\nx.tga\ny.esf\nz.lua"

		api, _, _, _, sysHandler := variableBefore("2.0", false, "test")
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte(fileList), nil).Once()
		sysHandler.On("MoveFile", testifyMock.Anything, testifyMock.Anything).Return(nil)
		sysHandler.On("WriteFile", testifyMock.Anything, testifyMock.Anything).Return(nil)

		err := api.Switch()

		/*
			error cases: cannot rename file (-> copy paste, then delete)
			file doesn't exist (-> check for it in IS_Files, error maybe)
			userscript wrong checksum (-> panic)
			is info file doesn't exist (-> panic)
			couldn't write to is info file (-> panic)
		*/

		assert.Nil(t, err)
		sysHandler.AssertCalled(t, "MoveFile", "./IS_Files/merp.pack", "./data/merp.pack")
		sysHandler.AssertCalled(t, "MoveFile", "./IS_Files/derp.pack", "./data/derp.pack")
		sysHandler.AssertCalled(t, "MoveFile", "./IS_Files/x.tga", "./data/campaigns/imperial_splendour/x.tga")
		sysHandler.AssertCalled(t, "MoveFile", "./IS_Files/y.esf", "./data/campaigns/imperial_splendour/y.esf")
		sysHandler.AssertCalled(t, "MoveFile", "./IS_Files/z.lua", "./data/campaigns/imperial_splendour/z.lua")
		sysHandler.AssertCalled(t, "MoveFile", "./IS_Files/user.empire_script.txt", "APPDATA/The Creative Assembly/Empire/scripts/user.empire_script.txt")
		sysHandler.AssertCalled(t, "WriteFile", "./IS_Files/IS_info.json", fmtInfoFile(true, "2.0", "test"))
		assert.True(t, api.IsActive())

		after(*api)
	}
	TestSuccessfullyActivateImpSplen(t)
}
