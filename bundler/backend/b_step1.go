package backend

import (
	"errors"
	"io/ioutil"
	"strings"
)

func (a *API) createFileList() ([]string, error) {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		return nil, err
	}

	fileList := []string{}
	for _, file := range files {
		fileName := file.Name()
		if strings.HasSuffix(fileName, ".pack") || strings.HasSuffix(fileName, ".tga") || strings.HasSuffix(fileName, ".esf") || strings.HasSuffix(fileName, ".lua") {
			fileList = append(fileList, fileName)
		}
	}

	return fileList, nil
}

func GetSetDifference(a, b []string) (diff []string) {
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

func (a *API) readComparisonFileList() ([]string, error) {
	expectedListBlob, err := a.Sh.ReadFile(a.userSettings.comparisonFileListPath)
	if err != nil {
		return nil, err
	}
	expectedList := strings.Split(string(expectedListBlob), "\n")
	return expectedList, err
}

func (a *API) compareFileLists(expected, actual []string) error {
	valuesInActualButNotInExpected := GetSetDifference(actual, expected)
	valuesInExpectedButNotInActual := GetSetDifference(expected, actual)

	// TODO: PROPER ERROR HERE
	if len(valuesInActualButNotInExpected) != 0 {
		return errors.New("Woops")
	}

	if len(valuesInExpectedButNotInActual) != 0 {
		return errors.New("Happens")
	}

	return nil
}

func (a *API) moveFilesIntoModFolder(fileList []string) error {
	for _, file := range fileList {
		source := "./" + file
		destination := a.setupBaseFolder + tempPath + modPath + file

		a.logger.Debugf("Moving from %s to %s", source, destination)
		err := a.Sh.MoveFile(source, destination)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *API) saveFileListIntoModFolder(fileList []string) error {
	destination := a.setupBaseFolder + tempPath + modPath + fileListFile

	fileListByteSlice := []byte(strings.Join(fileList, "\n"))

	err := a.Sh.WriteFile(destination, fileListByteSlice)
	return err
}
