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
