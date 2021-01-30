package main

import (
	"encoding/json"
	"io/ioutil"
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

	etwDir = getExecDirectory() + "/"
	a.logger.Infof("ETW/Current directory: %s", etwDir)
	// appDataDir = os.Getenv("APPDATA") + "appDataPath
	appDataDir = etwDir + "appDataFolder/" + appDataPath
	a.logger.Infof("AppData directory: %s", appDataDir)

	jsonFile, err := os.Open(etwDir + modPath + "IS_info.json")
	if err != nil {
		a.logger.Warnf("%v", err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	a.logger.Warnf("Info content %s", string(byteValue))

	var info Info

	json.Unmarshal(byteValue, &info)
	a.logger.Warnf("Info loaded ext %v", info)
	a.info = info
	a.logger.Warnf("Info loaded %v", a.info)

	return nil
}

// WailsShutdown is the shutdown function that is called when wails shuts down
func (a *API) WailsShutdown() {
	// a.store.close()
	// if a.cancelMonitoring != nil {
	// 	a.cancelMonitoring()
	// }

}
