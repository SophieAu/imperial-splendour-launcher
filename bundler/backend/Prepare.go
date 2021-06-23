package backend

import "strings"

var versionRegex = "^\\d+\\.\\d+(.\\d+)?$"

func (a *API) Prepare(sourcePath string, versionNumber string, packageRawFiles bool, fileListPath string) error {
	if err := a.validateUserInput(sourcePath, versionNumber, fileListPath); err != nil {
		return err
	}
	if !strings.HasSuffix(sourcePath, "/") {
		sourcePath = sourcePath + "/"
	}

	a.userSettings = UserSettings{sourcePath, versionNumber, packageRawFiles, fileListPath}
	a.setupBaseFolder = a.getSetupBaseFolder(sourcePath)

	if err := a.createTempFolder(); err != nil {
		return err
	}

	if err := a.prepareModFiles(sourcePath, fileListPath); err != nil {
		return err
	}

	if err := a.prepareUserScript(sourcePath); err != nil {
		return err
	}

	if err := a.downloadFiles(); err != nil {
		return err
	}

	if err := a.createInfoJSON(versionNumber); err != nil {
		return err
	}

	if err := a.updateSetupVersion(versionNumber); err != nil {
		return err
	}

	if err := a.updateSetupTempFolder(); err != nil {
		return err
	}

	return nil
}
