package backend

import (
	"encoding/json"
)

type info struct {
	IsActive           bool   `json:"isActive"`
	Version            string `json:"version"`
	UserScriptChecksum string `json:"usChecksum"`
}

func (a *API) setStatus(isActive bool) error {
	newInfo := a.info
	newInfo.IsActive = isActive

	newInfoJSON, err := json.MarshalIndent(newInfo, "", "\t")
	if err != nil {
		a.logger.Warnf("%v", err)
		return err
	}

	err = a.Sh.WriteFile(a.dirs.etw+modPath+infoFile, newInfoJSON)
	if err != nil {
		a.logger.Warnf("%v", err)
		return err
	}

	a.info.IsActive = isActive
	return nil
}
