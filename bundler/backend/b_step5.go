package backend

import (
	"imperial-splendour-bundler/backend/customErrors"
	"strings"
)

func (a *API) updateSetupVersion(versionNumber string) error {

	setupFileBlob, err := a.Sh.ReadFile(a.setupBaseFolder + setupFile)
	if err != nil {
		return a.error("Cannot read setup file: "+err.Error(), customErrors.VersionUpdate)
	}
	setupFileContent := string(setupFileBlob)

	tempCount := strings.Count(setupFileContent, VersionPlaceholder)
	if tempCount == 0 {
		return a.error("Couldn't find any version strings to replace", customErrors.VersionUpdate)
	}

	newSetupFileContent := strings.ReplaceAll(setupFileContent, VersionPlaceholder, versionNumber)

	err = a.Sh.WriteFile(a.setupBaseFolder+setupFile, []byte(newSetupFileContent))
	if err != nil {
		return a.error("Cannot update setup file: "+err.Error(), customErrors.VersionUpdate)
	}
	a.logToFrontend("Version was set in setup script")
	return nil
}

func (a *API) updateSetupTempFolder() error {
	tempFolderName := strings.TrimSuffix(tempPath, "/")

	setupFileBlob, err := a.Sh.ReadFile(a.setupBaseFolder + setupFile)
	if err != nil {
		return a.error("Cannot read setup file: "+err.Error(), customErrors.TempFolderUpdate)
	}
	setupFileContent := string(setupFileBlob)

	tempCount := strings.Count(setupFileContent, TempFolderPlaceholder)
	if tempCount == 0 {
		return a.error("Couldn't find any temp folder strings to replace", customErrors.TempFolderUpdate)
	}

	newSetupFileContent := strings.ReplaceAll(setupFileContent, TempFolderPlaceholder, tempFolderName)

	err = a.Sh.WriteFile(a.setupBaseFolder+setupFile, []byte(newSetupFileContent))
	if err != nil {
		return a.error("Cannot update setup file: "+err.Error(), customErrors.TempFolderUpdate)
	}
	a.logToFrontend("TempFolder was set in setup script")
	return nil
}
