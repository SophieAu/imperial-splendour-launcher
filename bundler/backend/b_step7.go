package backend

import "imperial-splendour-bundler/backend/customErrors"

func (a *API) buildZipBundle(sourcePath string, versionNumber string, fileListPath string) error {
	fileListFiles, err := a.readFileList(fileListPath)
	if err != nil {
		return a.error("Cannot read file list", customErrors.ZipFiles)
	}

	files := append(fileListFiles, fileListFile, userScript)
	outputFolderName := sourcePath + "IS-RotRv" + versionNumber + ".zip"

	if err := a.Sh.ZipFiles(outputFolderName, files); err != nil {
		return a.error("Couldn't zip files", customErrors.ZipFiles)
	}
	return nil
}
