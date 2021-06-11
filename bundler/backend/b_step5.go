package backend

import (
	"errors"
	"strings"
)

func (a *API) updateSetupVersion() error {

	setupFileBlob, err := a.Sh.ReadFile(a.setupBaseFolder + "/" + setupFile)
	if err != nil {
		return err
	}
	setupFileContent := string(setupFileBlob)

	tempCount := strings.Count(setupFileContent, VersionPlaceholder)
	if tempCount != 1 {
		return errors.New("pls no")
	}

	newSetupFileContent := strings.Replace(setupFileContent, VersionPlaceholder, a.userSettings.versionNumber, 1)

	err = a.Sh.WriteFile(a.setupBaseFolder+"/"+setupFile, []byte(newSetupFileContent))
	return err
}
