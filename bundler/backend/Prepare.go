package backend

var versionRegex = "^\\d*\\.\\d*$"

func (a *API) Prepare(sourcePath string, versionNumber string, packageRawFiles bool, fileListPath string) error {
	if err := a.validateUserInput(sourcePath, versionNumber, fileListPath); err != nil {
		return err
	}
	a.userSettings = UserSettings{sourcePath, versionNumber, packageRawFiles, fileListPath}
	a.setupBaseFolder = a.getSetupBaseFolder(sourcePath) + "/"

	if err := a.createTempFolder(); err != nil {
		return err
	}

	if err := a.prepareModFiles(sourcePath, fileListPath); err != nil {
		return err
	}

	// Step 2: User Script
	if err := a.prepareUserScript(sourcePath); err != nil {
		return err
	}

	// Step 3: Download setup files
	if err := a.downloadFiles(); err != nil {
		return err
	}

	// Step 4: Create IS_Info.json
	if err := a.createInfoJSON(versionNumber); err != nil {
		return err
	}

	// STEP 5: Update version in Setup
	if err := a.updateSetupVersion(versionNumber); err != nil {
		return err
	}

	return nil
}
