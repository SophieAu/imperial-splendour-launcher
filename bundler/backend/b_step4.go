package backend

import "encoding/json"

func (a *API) createInfoJSON() error {
	newInfo := Info{
		IsActive:           false,
		Version:            a.userSettings.versionNumber,
		UserScriptChecksum: "test",
	}

	newInfoJSON, err := json.MarshalIndent(newInfo, "", "\t")
	if err != nil {
		return err
	}

	targetFilePath := a.setupBaseFolder + "/" + tempPath + modPath + infoFile
	err = a.Sh.WriteFile(targetFilePath, newInfoJSON)
	return err
}
