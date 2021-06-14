package backend

import (
	"encoding/json"
	"imperial-splendour-bundler/backend/customErrors"
)

func (a *API) createInfoJSON(versionNumber string) error {
	newInfo := Info{
		IsActive:           false,
		Version:            versionNumber,
		UserScriptChecksum: "test",
	}

	newInfoJSON, err := json.MarshalIndent(newInfo, "", "\t")
	if err != nil {
		return a.error("Cannot serialize Info file: "+err.Error(), customErrors.InfoFile)
	}

	targetFilePath := a.setupBaseFolder + tempPath + modPath + infoFile
	err = a.Sh.WriteFile(targetFilePath, newInfoJSON)
	if err != nil {
		return a.error("Cannot save Info file: "+err.Error(), customErrors.InfoFile)
	}
	return nil
}
