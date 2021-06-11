package backend

import "errors"

func (a *API) Bundle(sourcePath string, versionNumber string, packageRawFiles bool, comparisonFileListPath string) error {
	a.userSettings = UserSettings{
		sourcePath,
		versionNumber,
		packageRawFiles,
		comparisonFileListPath,
	}

	// Step 0: Init
	a.setupBaseFolder = a.getSetupBaseFolder()
	if err := a.createTempFolder(); err != nil {
		return err
	}

	// Step 1: File List
	actualFileList, err := a.createFileList()
	if err != nil {
		return err
	}
	expectedFileList, err := a.readComparisonFileList()
	if err != nil {
		return err
	}

	if err := a.compareFileLists(expectedFileList, actualFileList); err != nil {
		return err
	}
	if err := a.moveFilesIntoModFolder(expectedFileList); err != nil {
		return err
	}
	if err := a.saveFileListIntoModFolder(expectedFileList); err != nil {
		return err
	}

	// Step 2: User Script
	if doesUserScriptExist := a.ensureUserScript(); !doesUserScriptExist {
		return errors.New("oh noes")
	}
	if err := a.moveUserScriptIntoModFolder(); err != nil {
		return err
	}

	// Step 3: Download setup files
	if err := a.downloadFiles(); err != nil {
		return err
	}

	// Step 4: Create IS_Info.json
	if err := a.createInfoJSON(); err != nil {
		return err
	}

	// STEP 5: Update version in Setup
	if err := a.updateSetupVersion(); err != nil {
		return err
	}
	// Step 6: Trigger InnoSetup CLI
	// !!!! WAIT HERE!
	if err := a.compileSetup(); err != nil {
		return err
	}

	// Step 7: optional: build zip
	if err := a.buildZipBundle(); err != nil {
		return err
	}

	return nil
}
