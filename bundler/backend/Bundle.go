package backend

func (a *API) Bundle() error {
	if err := a.compileSetup(); err != nil {
		return err
	}

	if !a.userSettings.packageRawFiles {
		return nil
	}

	if err := a.buildZipBundle(a.userSettings.sourcePath, a.userSettings.versionNumber, a.userSettings.fileListPath); err != nil {
		return err
	}

	return nil
}
