package main

import (
	"os"

	"github.com/wailsapp/wails"
	"github.com/wailsapp/wails/lib/logger"
)

var userScript = "user.empire_script.txt"
var etwDir = "/"
var modPath = "IS_Files/"
var dataPath = "data/"
var campaignPath = dataPath + "campaigns/imperial_splendour/"
var fileListFile = "IS_FileList.txt"

// API .
type API struct {
	runtime *wails.Runtime
	logger  *logger.CustomLogger
}

func (a *API) moveFile() error {
	oldLocation := etwDir + "/test.txt"
	newLocation := etwDir + "/src/test.txt"
	a.logger.Infof("Moving from %s to %s", oldLocation, newLocation)
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		a.logger.Errorf("%v", err)
		return err
	}
	return nil
}
