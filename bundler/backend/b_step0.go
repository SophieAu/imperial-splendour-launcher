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
	targetModPath := a.setupBaseFolder + tempPath + modPath
	if err := a.Sh.MkdirAll(targetModPath); err != nil {
		return a.error("Couldn't create "+targetModPath+": "+err.Error(), customErrors.TempFolderCreation)
	}
	a.logToFrontend("Created folder " + targetModPath)

	targetUninstallPath := a.setupBaseFolder + tempPath + uninstallPath
	if err := a.Sh.MkdirAll(targetUninstallPath); err != nil {
		return a.error("Couldn't create "+targetUninstallPath+": "+err.Error(), customErrors.TempFolderCreation)
	}
	a.logToFrontend("Created folder " + targetUninstallPath)

	return nil
}
