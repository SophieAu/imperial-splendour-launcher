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

type API struct {
	logger  Logger
	browser Browser
	window  Window
	Sh      Handler
	dirs    struct {
		etw     string
		appData string
	}
	info info
}

type modFiles struct {
	dataFiles     []string
	campaignFiles []string
}

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

func (a *API) readFileList() (*modFiles, error) {
	modFiles := modFiles{
		dataFiles:     []string{},
		campaignFiles: []string{},
	}

	fileBlob, err := a.Sh.ReadFile(a.dirs.etw + modPath + fileListFile)
	if err != nil {
		a.logger.Warnf("%v", err)
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

	if err := a.Sh.MoveFile(source, destination); err != nil {
		a.logger.Warnf("%v", err)
		return err
	}
	return nil
}

func (a *API) activateImpSplen() error {
	a.logger.Debug("Activating ImpSplen")

	files, err := a.readFileList()
	if err != nil {
		a.logger.Warnf("%v", err)
		return err
	}

	a.logger.Debug("Moving data files")
	for _, v := range (*files).dataFiles {
		if err := a.moveFile(a.dirs.etw+modPath+v, a.dirs.etw+dataPath+v); err != nil {
			_ = a.deactivateImpSplen()
			return err
		}
	}

	a.logger.Debug("Moving campaign files")
	for _, v := range (*files).campaignFiles {
		if err := a.moveFile(a.dirs.etw+modPath+v, a.dirs.etw+campaignPath+v); err != nil {
			_ = a.deactivateImpSplen()
			return err
		}
	}

	a.logger.Debug("Moving User Script")
	if err = a.moveFile(a.dirs.etw+modPath+userScript, a.dirs.appData+userScript); err != nil {
		_ = a.deactivateImpSplen()
		return err
	}

	if err = a.setStatus(true); err != nil {
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
		a.logger.Warnf("%v", err)
		return err
	}

	a.logger.Debug("Moving data files")
	for _, v := range files.dataFiles {
		err := a.moveFile(a.dirs.etw+dataPath+v, a.dirs.etw+modPath+v)
		if err != nil {
			a.logger.Warnf("%v", err)
		}
	}

	a.logger.Debug("Moving campaign files")
	for _, v := range files.campaignFiles {
		err := a.moveFile(a.dirs.etw+campaignPath+v, a.dirs.etw+modPath+v)
		if err != nil {
			a.logger.Warnf("%v", err)
		}
	}

	a.logger.Debug("Moving User Script")
	err = a.moveFile(a.dirs.appData+userScript, a.dirs.etw+modPath+userScript)
	if err != nil {
		a.logger.Warnf("%v", err)
	}

	err = a.setStatus(false)
	a.logger.Debug("ImpSplen deactivated")
	return err
}

func (a *API) deleteAllFiles() error {
	return a.Sh.Remove(a.dirs.etw + modPath)
}
