package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails"
	"github.com/wailsapp/wails/lib/logger"
)

var userScript = "user.empire_script.txt"
var etwDir = "/"
var modPath = "IS_Files/"
var dataPath = "data/"
var campaignPath = dataPath + "campaigns/imperial_splendour/"

// API .
type API struct {
	runtime *wails.Runtime
	logger  *logger.CustomLogger
}

// WailsInit is the init fuction for the wails runtime
func (a *API) WailsInit(runtime *wails.Runtime) error {
	a.runtime = runtime
	a.logger = runtime.Log.New("API")

	a.logger.Info("This is fine")
	feeling := "okay"
	a.logger.Infof("I'm %s with the events that are currently unfolding", feeling)

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)

	etwDir = exPath

	return nil
}

// GoToWebsite .
func (a *API) GoToWebsite() {
	a.runtime.Browser.OpenURL("https://imperialsplendour.com/")
}

// Exit .
func (a *API) Exit() {
	a.runtime.Window.Close()

}

// Play .
func (a *API) Play() error {
	err := a.runtime.Browser.OpenURL("steam://rungameid/10500")
	if err != nil {
		return err
		// return thrown promise here asking the user if they have steam installed
	}
	a.Exit()
	return nil
}

// WailsShutdown is the shutdown function that is called when wails shuts down
func (a *API) WailsShutdown() {
	// a.store.close()
	// if a.cancelMonitoring != nil {
	// 	a.cancelMonitoring()
	// }

}
