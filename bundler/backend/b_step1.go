package backend

import (
	"errors"
	"imperial-splendour-bundler/backend/customErrors"
	"strings"
)

func (a *API) prepareModFiles(sourcePath, fileListPath string) error {
	actualFileList, err := a.Sh.GetDirContentByName(sourcePath)
	if err != nil {
		return a.error("Cannot read dir under"+sourcePath+": "+err.Error(), customErrors.ReadSourceDir)
	}
	expectedFileList, err := a.readFileList(fileListPath)
	if err != nil {
		return err
	}

	if err := a.compareFileLists(expectedFileList, actualFileList); err != nil {
		return err
	}
	a.logToFrontend("All mod files are accounted for.")
	if err := a.moveFilesIntoModFolder(sourcePath, expectedFileList); err != nil {
		return err
	}
	a.logToFrontend("Mod files were moved to mod folder.")
	if err := a.saveFileListIntoModFolder(expectedFileList); err != nil {
		return err
	}
	a.logToFrontend("File list was added to mod folder.")
	return nil
}

func (a *API) readFileList(fileListPath string) ([]string, error) {
	expectedListBlob, err := a.Sh.ReadFile(fileListPath)
	if err != nil {
		return nil, a.error("Cannot read file list: "+err.Error(), customErrors.ReadFileList)
	}
	expectedList := strings.Split(string(expectedListBlob), "\n")
	return expectedList, nil
}

func GetDifference(a, b []string) (diff []string) {
	m := make(map[string]bool)

	// set values in b as keys in the map
	for _, item := range b {
		m[item] = true
	}

	// check that values in a are keys in the map
	for _, itemInA := range a {
		if _, isKey := m[itemInA]; !isKey {
			diff = append(diff, itemInA)
		}
	}
	return diff
}

func (a *API) compareFileLists(expected, actual []string) error {
	valuesInExpectedButNotInActual := GetDifference(expected, actual)
	if len(valuesInExpectedButNotInActual) != 0 {
		missingFiles := strings.Join(valuesInExpectedButNotInActual, ", ")
		return a.error("Couldn't find the following files: "+missingFiles, errors.New(customErrors.FileMissing.Error()+" "+missingFiles))
	}
	return nil
}

func (a *API) moveFilesIntoModFolder(sourcePath string, fileList []string) error {
	for _, file := range fileList {
		source := sourcePath + "/" + file
		destination := a.setupBaseFolder + tempPath + modPath + file

		a.logToFrontend("Moving from " + source + " to " + destination)
		err := a.Sh.MoveFile(source, destination)
		if err != nil {
			return a.error("Couldn't move "+file, errors.New(customErrors.MoveFile.Error()+" "+file))
		}
	}

	return nil
}

func (a *API) saveFileListIntoModFolder(fileList []string) error {
	destination := a.setupBaseFolder + tempPath + modPath + fileListFile

	fileListByteSlice := []byte(strings.Join(fileList, "\n"))

	err := a.Sh.WriteFile(destination, fileListByteSlice)
	if err != nil {
		return a.error("Couldn't save file list: "+err.Error(), customErrors.SaveFileList)
	}
	return nil
}
