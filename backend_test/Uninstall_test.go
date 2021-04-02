package backend_test

import (
	"github.com/stretchr/testify/assert"
	testifyMock "github.com/stretchr/testify/mock"

	"testing"
)

func TestUninstall(t *testing.T) {
	TestSuccessfullyUninstallImpSplen := func(t *testing.T) {
		fileList := "merp.pack\nz.lua"

		api, _, _, _, sysHandler := variableBefore("2.0", true, "test")
		sysHandler.On("ReadFile", "./IS_Files/IS_FileList.txt").Return([]byte(fileList), nil).Once()
		sysHandler.On("MoveFile", testifyMock.Anything, testifyMock.Anything).Return(nil)
		sysHandler.On("WriteFile", "./IS_Files/IS_info.json", fmtInfoFile(false, "2.0", "test")).Return(nil)
		sysHandler.On("Remove", testifyMock.Anything, testifyMock.Anything).Return(nil)

		err := api.Uninstall()

		assert.Nil(t, err)
		// deactivating
		sysHandler.AssertCalled(t, "MoveFile", "./data/merp.pack", "./IS_Files/merp.pack")
		sysHandler.AssertCalled(t, "MoveFile", "./data/campaigns/imperial_splendour/z.lua", "./IS_Files/z.lua")
		sysHandler.AssertCalled(t, "MoveFile", "APPDATA/The Creative Assembly/Empire/scripts/user.empire_script.txt", "./IS_Files/user.empire_script.txt")
		sysHandler.AssertCalled(t, "WriteFile", "./IS_Files/IS_info.json", fmtInfoFile(false, "2.0", "test"))
		//deleting
		sysHandler.AssertCalled(t, "Remove", "./IS_Files/")

		after(*api)
	}
	TestSuccessfullyUninstallImpSplen(t)
}
