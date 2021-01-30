package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/wailsapp/wails"
	"github.com/wailsapp/wails/lib/logger"
)

// API .
type API struct {
	runtime *wails.Runtime
	logger  *logger.CustomLogger
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
	file, err := os.Open(etwDir + modPath + fileListFile)
	if err != nil {
		a.logger.Fatalf("%v", err)
	}
	defer file.Close()

	modFiles := ModFiles{
		dataFiles:     []string{},
		campaignFiles: []string{},
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		item := scanner.Text()

		if strings.HasSuffix(item, ".pack") {
			modFiles.dataFiles = append(modFiles.dataFiles, item)
		} else if strings.HasSuffix(item, ".tga") || strings.HasSuffix(item, ".esf") || strings.HasSuffix(item, ".lua") {
			modFiles.campaignFiles = append(modFiles.campaignFiles, item)
		} else {
			a.logger.Warnf("Unknown file '%s' found in file list", item)
		}

		fmt.Println(item)
	}

	if err := scanner.Err(); err != nil {
		a.logger.Fatalf("%v", err)
	}

	return &modFiles, nil
}

func (a *API) moveFile(source, destination string) error {
	a.logger.Infof("Moving from %s to %s", source, destination)
	err := os.Rename(source, destination)
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

	err = ioutil.WriteFile(etwDir+modPath+infoFile, newInfoJSON, 0644)
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
