package main

import (
	"os"
	"path/filepath"

	"github.com/wailsapp/wails"
)

func getExecDirectory() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}

// WailsInit is the init fuction for the wails runtime
func (a *API) WailsInit(runtime *wails.Runtime) error {
	a.runtime = runtime
	a.logger = runtime.Log.New("API")

	etwDir = getExecDirectory()
	a.logger.Infof("ETW/Current directory: %s", etwDir)

	return nil
}

// WailsShutdown is the shutdown function that is called when wails shuts down
func (a *API) WailsShutdown() {
	// a.store.close()
	// if a.cancelMonitoring != nil {
	// 	a.cancelMonitoring()
	// }

}
