package backend

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

const (
	AppName = "Imperial Splendour: Rise of the Republic"

	appDataPath  = "The Creative Assembly/Empire/scripts/"
	modPath      = "IS_Files/"
	dataPath     = "data/"
	campaignPath = dataPath + "campaigns/imperial_splendour/"

	userScript   = "user.empire_script.txt"
	fileListFile = "IS_FileList.txt"
	infoFile     = "IS_info.json"

	websiteURL  = "https://imperialsplendour.com/"
	etwSteamURI = "steam://rungameid/10500"
)

var (
	etwDir     = ""
	appDataDir = ""
)

type API struct {
	logger  Logger
	browser Browser
	window  Window
	Sh      Handler
	info    info
}

type info struct {
	IsActive           bool   `json:"isActive"`
	Version            string `json:"version"`
	UserScriptChecksum string `json:"usChecksum"`
}

type modFiles struct {
	dataFiles     []string
	campaignFiles []string
}

func (a *API) readFileList() (*modFiles, error) {
	modFiles := modFiles{
		dataFiles:     []string{},
		campaignFiles: []string{},
	}

	fileBlob, err := a.Sh.ReadFile(etwDir + modPath + fileListFile)
	if err != nil {
		a.logger.Errorf("%v", err)
		return nil, err
	}

	var errMsg string
	for _, file := range strings.Split(string(fileBlob), "\n") {
		if strings.HasSuffix(file, ".pack") {
			modFiles.dataFiles = append(modFiles.dataFiles, file)
		} else if strings.HasSuffix(file, ".tga") || strings.HasSuffix(file, ".esf") || strings.HasSuffix(file, ".lua") {
			modFiles.campaignFiles = append(modFiles.campaignFiles, file)
		} else if file != "" {
			errMsg = errMsg + fmt.Sprintf("Unknown file '%s' found in file list\n", file)
		}
	}
	if errMsg != "" {
		return nil, errors.New(errMsg)
	}
	return &modFiles, nil
}

func (a *API) moveFile(source, destination string) error {
	a.logger.Debugf("Moving from %s to %s", source, destination)

	err := a.Sh.MoveFile(source, destination)
	if err != nil {
		a.logger.Errorf("%v", err)
		return err
	}
	return nil
}

func (a *API) setStatus(isActive bool) error {
	newInfo := a.info
	newInfo.IsActive = isActive

	newInfoJSON, err := json.MarshalIndent(newInfo, "", "\t")
	if err != nil {
		a.logger.Errorf("%v", err)
		return err
	}

	err = a.Sh.WriteFile(etwDir+modPath+infoFile, newInfoJSON)
	if err != nil {
		a.logger.Errorf("%v", err)
		return err
	}

	a.info.IsActive = isActive
	return nil
}

func (a *API) activateImpSplen() error {
	a.logger.Debug("Activating ImpSplen")

	files, err := a.readFileList()
	if err != nil {
		a.logger.Errorf("%v", err)
		return err
	}

	a.logger.Debug("Moving data files")
	for _, v := range (*files).dataFiles {
		err := a.moveFile(etwDir+modPath+v, etwDir+dataPath+v)
		if err != nil {
			_ = a.deactivateImpSplen()
			return err
		}
	}

	a.logger.Debug("Moving campaign files")
	for _, v := range (*files).campaignFiles {
		err := a.moveFile(etwDir+modPath+v, etwDir+campaignPath+v)
		if err != nil {
			_ = a.deactivateImpSplen()
			return err
		}
	}

	a.logger.Debug("Moving User Script")
	err = a.moveFile(etwDir+modPath+userScript, appDataDir+userScript)
	if err != nil {
		_ = a.deactivateImpSplen()
		return err
	}

	err = a.setStatus(true)
	if err != nil {
		_ = a.deactivateImpSplen()
		return err
	}
	a.logger.Debug("ImpSplen activated")
	return nil
}

func (a *API) deactivateImpSplen() error {
	a.logger.Debug("Deactivating ImpSplen")

	files, err := a.readFileList()
	if err != nil {
		a.logger.Errorf("%v", err)
		return err
	}

	a.logger.Debug("Moving data files")
	for _, v := range files.dataFiles {
		err := a.moveFile(etwDir+dataPath+v, etwDir+modPath+v)
		if err != nil {
			a.logger.Errorf("%v", err)
		}
	}

	a.logger.Debug("Moving campaign files")
	for _, v := range files.campaignFiles {
		err := a.moveFile(etwDir+campaignPath+v, etwDir+modPath+v)
		if err != nil {
			a.logger.Errorf("%v", err)
		}
	}

	a.logger.Debug("Moving User Script")
	err = a.moveFile(appDataDir+userScript, etwDir+modPath+userScript)
	if err != nil {
		a.logger.Errorf("%v", err)
	}

	err = a.setStatus(false)
	a.logger.Debug("ImpSplen deactivated")
	return err
}
