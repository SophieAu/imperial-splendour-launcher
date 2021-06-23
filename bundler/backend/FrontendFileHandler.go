package backend

import (
	"imperial-splendour-bundler/backend/customErrors"
	"os"
)

func (a *API) SelectSourceDir() (string, error) {
	dir := a.dialog.SelectDirectory()

	sourceDir, err := os.Stat(dir)
	if err != nil || !sourceDir.IsDir() {
		return "", customErrors.InvalidDir
	}
	return dir + "/", nil
}

func (a *API) SelectFileListLocation() (string, error) {
	file := a.dialog.SelectFile("Select the file list file", "*.txt")

	fileListFile, err := os.Stat(file)
	if err != nil || fileListFile.IsDir() {
		return "", customErrors.InvalidFile
	}
	return file + "/", nil
}
