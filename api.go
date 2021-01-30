package main

import (
	"encoding/json"
	"strings"

	"github.com/wailsapp/wails"
	"github.com/wailsapp/wails/lib/logger"
)

// API .
type API struct {
	runtime *wails.Runtime
	logger  *logger.CustomLogger
	fh      Handler
	info    Info
}

// Info .
type Info struct {
	IsActive bool   `json:"isActive"`
	Version  string `json:"version"`
}

// ModFiles .
type ModFiles struct {
	dataFiles     []string
	campaignFiles []string
}

func (a *API) readFileList() (*ModFiles, error) {
	modFiles := ModFiles{
		dataFiles:     []string{},
		campaignFiles: []string{},
	}

	fileBlob, err := a.fh.ReadFile(etwDir + modPath + fileListFile)
	if err != nil {
		a.logger.Fatalf("%v", err)
	}

	for _, file := range strings.Split(string(fileBlob), "\n") {
		if strings.HasSuffix(file, ".pack") {
			modFiles.dataFiles = append(modFiles.dataFiles, file)
		} else if strings.HasSuffix(file, ".tga") || strings.HasSuffix(file, ".esf") || strings.HasSuffix(file, ".lua") {
			modFiles.campaignFiles = append(modFiles.campaignFiles, file)
		} else if file != "" {
			a.logger.Warnf("Unknown file '%s' found in file list", file)
		}
	}

	return &modFiles, nil
}

func (a *API) moveFile(source, destination string) error {
	a.logger.Debugf("Moving from %s to %s", source, destination)

	err := a.fh.MoveFile(source, destination)
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

	err = a.fh.WriteFile(etwDir+modPath+infoFile, newInfoJSON, 0644)
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
		a.moveFile(etwDir+modPath+v, etwDir+dataPath+v)
	}

	a.logger.Debug("Moving campaign files")
	for _, v := range (*files).campaignFiles {
		a.moveFile(etwDir+modPath+v, etwDir+campaignPath+v)
	}

	a.logger.Debug("Moving User Script")
	a.moveFile(etwDir+modPath+userScript, appDataDir+userScript)

	a.setStatus(true)
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
		a.moveFile(etwDir+dataPath+v, etwDir+modPath+v)
	}

	a.logger.Debug("Moving campaign files")
	for _, v := range files.campaignFiles {
		a.moveFile(etwDir+campaignPath+v, etwDir+modPath+v)
	}

	a.logger.Debug("Moving User Script")
	a.moveFile(appDataDir+userScript, etwDir+modPath+userScript)

	a.setStatus(false)
	a.logger.Debug("ImpSplen deactivated")
	return nil
}
