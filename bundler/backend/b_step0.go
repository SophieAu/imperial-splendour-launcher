package backend

import (
	"fmt"
)

func (a *API) getSetupBaseFolder() string {
	hasPreviousSetup, _ := a.Sh.DoesFileExist(a.userSettings.sourcePath + "/" + preferredSetupBaseFolder)

	if !hasPreviousSetup {
		return a.userSettings.sourcePath + "/" + preferredSetupBaseFolder
	}

	setupFolderBase := preferredSetupBaseFolder + "_"
	number := 1
	for {
		currentFolderIteration := setupFolderBase + fmt.Sprint(number)
		doesXist, _ := a.Sh.DoesFileExist(a.userSettings.sourcePath + "/" + currentFolderIteration)
		if !doesXist {
			return currentFolderIteration
		}
		number++
	}
}

func (a *API) createTempFolder() error {
	basePath := a.setupBaseFolder + "/"

	if err := a.Sh.MkdirAll(basePath + tempPath + modPath); err != nil {
		return err
	}
	if err := a.Sh.MkdirAll(basePath + tempPath + uninstallPath); err != nil {
		return err
	}

	return nil
}
