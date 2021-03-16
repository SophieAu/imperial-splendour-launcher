package backend_test

import (
	"errors"
	"imperial-splendour-launcher/backend"
	"imperial-splendour-launcher/backend/mock"
	"strconv"

	"github.com/stretchr/testify/assert"
	testifyMock "github.com/stretchr/testify/mock"

	"testing"
)

func before() (*backend.API, *mock.Browser, *mock.Window, *mock.Logger, *mock.SystemHandler) {
	return variableBefore("2.0", true, "test")
}

func variableBefore(version string, isActive bool, usChecksum string) (*backend.API, *mock.Browser, *mock.Window, *mock.Logger, *mock.SystemHandler) {
	mockS := &mock.SystemHandler{}
	mockB := &mock.Browser{}
	mockW := &mock.Window{}
	mockL := &mock.Logger{}

	mockS.On("Executable").Return(".", nil)
	mockW.On("Close").Return()
	mockL.On("Infof", testifyMock.Anything, testifyMock.Anything).Return()
	mockL.On("Info", testifyMock.Anything).Return()
	mockL.On("Errorf", testifyMock.Anything, testifyMock.Anything).Return()
	mockL.On("Warnf", testifyMock.Anything, testifyMock.Anything).Return()
	mockL.On("Debugf", testifyMock.Anything, testifyMock.Anything).Return()
	mockL.On("Debug", testifyMock.Anything).Return(nil)
	mockS.On("ReadFile", "IS_Files/IS_info.json").Return([]byte("{\"isActive\": "+strconv.FormatBool(isActive)+", \"version\": \""+version+"\", \"usChecksum\": \""+usChecksum+"\"}"), nil)

	api := &backend.API{}
	err := api.Init(mockB, mockW, mockL, mockS)
	if err != nil {
		panic(err)
	}

	return api, mockB, mockW, mockL, mockS
}

func after(api backend.API) {
}

func TestVersion(t *testing.T) {
	api, _, _, _, _ := variableBefore("2.1", false, "test")

	version := api.Version()
	assert.Equal(t, "2.1", version)

	after(*api)
}

func TestIsActive(t *testing.T) {
	api, _, _, _, _ := variableBefore("2.1", false, "test")

	isActive := api.IsActive()
	assert.Equal(t, false, isActive)

	api, _, _, _, _ = variableBefore("2.1", true, "test")

	isActive = api.IsActive()
	assert.Equal(t, true, isActive)

	after(*api)
}

func TestPlay(t *testing.T) {
	api, browser, window, _, _ := before()

	// error in opening URL
	browser.On("OpenURL", testifyMock.Anything).Return(errors.New("error")).Once()
	err := api.Play()

	assert.NotNil(t, err)
	browser.AssertCalled(t, "OpenURL", "steam://rungameid/10500")
	window.AssertNotCalled(t, "Close")

	// working properly
	browser.On("OpenURL", testifyMock.Anything).Return(nil).Once()
	err = api.Play()

	assert.Nil(t, err)
	browser.AssertCalled(t, "OpenURL", "steam://rungameid/10500")
	window.AssertCalled(t, "Close")

	after(*api)
}

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
		sysHandler.On("ReadFile", "IS_Files/IS_FileList.txt").Return(nil, errors.New("FileNotFound")).Once()

		err := api.Switch()

		assert.EqualError(t, err, "FileNotFound")
		sysHandler.AssertNotCalled(t, "MoveFile", testifyMock.Anything, testifyMock.Anything)
		assert.True(t, api.IsActive())

		after(*api)
	}
	TestFileListNotFound(t)

	TestUnknownFileInFileList := func(t *testing.T) {
		api, _, _, _, sysHandler := variableBefore("2.0", true, "test")
		sysHandler.On("ReadFile", "IS_Files/IS_FileList.txt").Return([]byte("a\nb\n"), nil).Once()

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
		sysHandler.On("ReadFile", "IS_Files/IS_FileList.txt").Return([]byte(fileList), nil).Once()
		sysHandler.On("MoveFile", "data/campaigns/imperial_splendour/y.esf", "IS_Files/y.esf").Return(errors.New("FileNotFound"))
		sysHandler.On("MoveFile", testifyMock.Anything, testifyMock.Anything).Return(nil)
		sysHandler.On("WriteFile", "IS_Files/IS_info.json", fmtInfoFile(false, "2.0", "test")).Return(nil)

		err := api.Switch()

		assert.Nil(t, err)
		sysHandler.AssertCalled(t, "MoveFile", "data/merp.pack", "IS_Files/merp.pack")
		sysHandler.AssertCalled(t, "MoveFile", "data/derp.pack", "IS_Files/derp.pack")
		sysHandler.AssertCalled(t, "MoveFile", "data/campaigns/imperial_splendour/x.tga", "IS_Files/x.tga")
		sysHandler.AssertCalled(t, "MoveFile", "data/campaigns/imperial_splendour/y.esf", "IS_Files/y.esf")
		sysHandler.AssertCalled(t, "MoveFile", "data/campaigns/imperial_splendour/z.lua", "IS_Files/z.lua")
		sysHandler.AssertCalled(t, "MoveFile", "./appDataFolder/The Creative Assembly/Empire/scripts/user.empire_script.txt", "IS_Files/user.empire_script.txt")
		sysHandler.AssertCalled(t, "WriteFile", "IS_Files/IS_info.json", fmtInfoFile(false, "2.0", "test"))
		assert.False(t, api.IsActive())

		after(*api)
	}
	TestSwallowErrorsWhenDeactivating(t)

	TestSuccessfullyDeactivateImpSplen := func(t *testing.T) {
		fileList := "merp.pack\nderp.pack\nx.tga\ny.esf\nz.lua"

		api, _, _, _, sysHandler := variableBefore("2.0", true, "test")
		sysHandler.On("ReadFile", "IS_Files/IS_FileList.txt").Return([]byte(fileList), nil).Once()
		sysHandler.On("MoveFile", testifyMock.Anything, testifyMock.Anything).Return(nil)
		sysHandler.On("WriteFile", "IS_Files/IS_info.json", fmtInfoFile(false, "2.0", "test")).Return(nil)

		err := api.Switch()

		assert.Nil(t, err)
		sysHandler.AssertCalled(t, "MoveFile", "data/merp.pack", "IS_Files/merp.pack")
		sysHandler.AssertCalled(t, "MoveFile", "data/derp.pack", "IS_Files/derp.pack")
		sysHandler.AssertCalled(t, "MoveFile", "data/campaigns/imperial_splendour/x.tga", "IS_Files/x.tga")
		sysHandler.AssertCalled(t, "MoveFile", "data/campaigns/imperial_splendour/y.esf", "IS_Files/y.esf")
		sysHandler.AssertCalled(t, "MoveFile", "data/campaigns/imperial_splendour/z.lua", "IS_Files/z.lua")
		sysHandler.AssertCalled(t, "MoveFile", "./appDataFolder/The Creative Assembly/Empire/scripts/user.empire_script.txt", "IS_Files/user.empire_script.txt")
		sysHandler.AssertCalled(t, "WriteFile", "IS_Files/IS_info.json", fmtInfoFile(false, "2.0", "test"))
		assert.False(t, api.IsActive())

		after(*api)
	}
	TestSuccessfullyDeactivateImpSplen(t)

	TestRollbackOnErrorWhileActivating := func(t *testing.T) {
		fileList := "merp.pack\nderp.pack\nx.tga\ny.esf\nz.lua"

		api, _, _, _, sysHandler := variableBefore("2.0", false, "test")
		sysHandler.On("ReadFile", "IS_Files/IS_FileList.txt").Return([]byte(fileList), nil).Once()
		sysHandler.On("ReadFile", "IS_Files/IS_FileList.txt").Return([]byte(fileList), nil).Once()
		sysHandler.On("MoveFile", "IS_Files/y.esf", "data/campaigns/imperial_splendour/y.esf").Return(errors.New("FileNotFound"))
		sysHandler.On("MoveFile", testifyMock.Anything, testifyMock.Anything).Return(nil)
		sysHandler.On("WriteFile", testifyMock.Anything, testifyMock.Anything).Return(nil)

		err := api.Switch()

		assert.NotNil(t, err)
		sysHandler.AssertCalled(t, "MoveFile", "IS_Files/merp.pack", "data/merp.pack")
		sysHandler.AssertCalled(t, "MoveFile", "IS_Files/derp.pack", "data/derp.pack")
		sysHandler.AssertCalled(t, "MoveFile", "IS_Files/x.tga", "data/campaigns/imperial_splendour/x.tga")
		sysHandler.AssertCalled(t, "MoveFile", "IS_Files/y.esf", "data/campaigns/imperial_splendour/y.esf")

		sysHandler.AssertCalled(t, "MoveFile", "data/merp.pack", "IS_Files/merp.pack")
		sysHandler.AssertCalled(t, "MoveFile", "data/derp.pack", "IS_Files/derp.pack")
		sysHandler.AssertCalled(t, "MoveFile", "data/campaigns/imperial_splendour/x.tga", "IS_Files/x.tga")
		sysHandler.AssertCalled(t, "MoveFile", "data/campaigns/imperial_splendour/y.esf", "IS_Files/y.esf")
		sysHandler.AssertCalled(t, "MoveFile", "data/campaigns/imperial_splendour/z.lua", "IS_Files/z.lua")
		sysHandler.AssertCalled(t, "MoveFile", "./appDataFolder/The Creative Assembly/Empire/scripts/user.empire_script.txt", "IS_Files/user.empire_script.txt")
		sysHandler.AssertCalled(t, "WriteFile", "IS_Files/IS_info.json", fmtInfoFile(false, "2.0", "test"))

		sysHandler.AssertNotCalled(t, "MoveFile", "IS_Files/z.lua", "data/campaigns/imperial_splendour/z.lua")
		sysHandler.AssertNotCalled(t, "MoveFile", "IS_Files/user.empire_script.txt", "./appDataFolder/The Creative Assembly/Empire/scripts/user.empire_script.txt")
		sysHandler.AssertNotCalled(t, "WriteFile", "IS_Files/IS_info.json", fmtInfoFile(true, "2.0", "test"))

		assert.False(t, api.IsActive())

		after(*api)
	}
	TestRollbackOnErrorWhileActivating(t)

	TestSuccessfullyActivateImpSplen := func(t *testing.T) {
		fileList := "merp.pack\nderp.pack\nx.tga\ny.esf\nz.lua"

		api, _, _, _, sysHandler := variableBefore("2.0", false, "test")
		sysHandler.On("ReadFile", "IS_Files/IS_FileList.txt").Return([]byte(fileList), nil).Once()
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
		sysHandler.AssertCalled(t, "MoveFile", "IS_Files/merp.pack", "data/merp.pack")
		sysHandler.AssertCalled(t, "MoveFile", "IS_Files/derp.pack", "data/derp.pack")
		sysHandler.AssertCalled(t, "MoveFile", "IS_Files/x.tga", "data/campaigns/imperial_splendour/x.tga")
		sysHandler.AssertCalled(t, "MoveFile", "IS_Files/y.esf", "data/campaigns/imperial_splendour/y.esf")
		sysHandler.AssertCalled(t, "MoveFile", "IS_Files/z.lua", "data/campaigns/imperial_splendour/z.lua")
		sysHandler.AssertCalled(t, "MoveFile", "IS_Files/user.empire_script.txt", "./appDataFolder/The Creative Assembly/Empire/scripts/user.empire_script.txt")
		sysHandler.AssertCalled(t, "WriteFile", "IS_Files/IS_info.json", fmtInfoFile(true, "2.0", "test"))
		assert.True(t, api.IsActive())

		after(*api)
	}
	TestSuccessfullyActivateImpSplen(t)
}

func TestGoToWebsite(t *testing.T) {
	api, browser, window, _, _ := before()

	// error in opening URL
	browser.On("OpenURL", testifyMock.Anything).Return(errors.New("error")).Once()
	err := api.GoToWebsite()

	assert.NotNil(t, err)
	browser.AssertCalled(t, "OpenURL", "https://imperialsplendour.com/")
	window.AssertNotCalled(t, "Close")

	// working properly
	browser.On("OpenURL", testifyMock.Anything).Return(nil).Once()
	err = api.Play()

	assert.Nil(t, err)
	browser.AssertCalled(t, "OpenURL", "https://imperialsplendour.com/")
	window.AssertCalled(t, "Close")

	after(*api)
}

func TestUninstall(t *testing.T) {
	TestSuccessfullyUninstallImpSplen := func(t *testing.T) {
		fileList := "merp.pack\nz.lua"

		api, _, _, _, sysHandler := variableBefore("2.0", true, "test")
		sysHandler.On("ReadFile", "IS_Files/IS_FileList.txt").Return([]byte(fileList), nil).Once()
		sysHandler.On("MoveFile", testifyMock.Anything, testifyMock.Anything).Return(nil)
		sysHandler.On("WriteFile", "IS_Files/IS_info.json", fmtInfoFile(false, "2.0", "test")).Return(nil)
		sysHandler.On("Remove", testifyMock.Anything, testifyMock.Anything).Return(nil)

		err := api.Uninstall()

		assert.Nil(t, err)
		// deactivating
		sysHandler.AssertCalled(t, "MoveFile", "data/merp.pack", "IS_Files/merp.pack")
		sysHandler.AssertCalled(t, "MoveFile", "data/campaigns/imperial_splendour/z.lua", "IS_Files/z.lua")
		sysHandler.AssertCalled(t, "MoveFile", "./appDataFolder/The Creative Assembly/Empire/scripts/user.empire_script.txt", "IS_Files/user.empire_script.txt")
		sysHandler.AssertCalled(t, "WriteFile", "IS_Files/IS_info.json", fmtInfoFile(false, "2.0", "test"))
		//deleting
		sysHandler.AssertCalled(t, "Remove", "IS_Files/")

		after(*api)
	}
	TestSuccessfullyUninstallImpSplen(t)
}

func TestExit(t *testing.T) {
	api, browser, window, _, _ := before()

	// error in opening URL
	browser.On("OpenURL", testifyMock.Anything).Return(errors.New("error")).Once()
	err := api.Play()

	assert.NotNil(t, err)
	browser.AssertCalled(t, "OpenURL", "steam://rungameid/10500")
	window.AssertNotCalled(t, "Close")

	// working properly
	browser.On("OpenURL", testifyMock.Anything).Return(nil).Once()
	err = api.Play()

	assert.Nil(t, err)
	browser.AssertCalled(t, "OpenURL", "steam://rungameid/10500")
	window.AssertCalled(t, "Close")

	after(*api)
}
