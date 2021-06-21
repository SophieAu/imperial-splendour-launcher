package backend

func (a *API) Bundle() error {
	// Step 6: Compile InnoSetup
	if err := a.compileSetup(); err != nil {
		return err
	}

	if !a.userSettings.packageRawFiles {
		return nil
	}

	// Step 7: optional: build zip
	if err := a.buildZipBundle(a.userSettings.sourcePath, a.userSettings.versionNumber, a.userSettings.fileListPath); err != nil {
		return err
	}

	return nil
}
