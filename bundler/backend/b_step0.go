package backend

import (
	"fmt"
	"imperial-splendour-bundler/backend/customErrors"
)

func (a *API) getSetupBaseFolder(sourcePath string) string {
	preferredBase := sourcePath + "/" + preferredSetupBaseFolder

	hasPreviousSetup, _ := a.Sh.DoesFileExist(preferredBase)
	if !hasPreviousSetup {
		a.logger.Info("Setup Folder: " + preferredBase)
		return preferredBase
	}

	setupFolderBase := preferredBase + "_"
	number := 1
	for {
		currentFolderIteration := setupFolderBase + fmt.Sprint(number)
		doesXist, _ := a.Sh.DoesFileExist(currentFolderIteration)
		if !doesXist {
			a.logger.Info("Setup Folder: " + currentFolderIteration)
			return currentFolderIteration
		}
		number++
	}
}

func (a *API) createTempFolder() error {
	basePath := a.setupBaseFolder + "/"

	if err := a.Sh.MkdirAll(basePath + tempPath + modPath); err != nil {
		return a.error("Couldn't create "+basePath+tempPath+modPath+": "+err.Error(), customErrors.TempFolderCreation)
	}
	a.logToFrontend("Created folder " + basePath + tempPath + modPath)

	if err := a.Sh.MkdirAll(basePath + tempPath + uninstallPath); err != nil {
		return a.error("Couldn't create "+basePath+tempPath+uninstallPath+": "+err.Error(), customErrors.TempFolderCreation)
	}
	a.logToFrontend("Created folder " + basePath + tempPath + uninstallPath)

	return nil
}
