package backend

import (
	"encoding/json"
	"path/filepath"

	"github.com/wailsapp/wails"
)

func (a *API) getExecDirectory() string {
	ex, err := a.Sh.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}

func (a *API) loadInfoFromFile() {
	byteValue, err := a.Sh.ReadFile(etwDir + modPath + infoFile)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(byteValue, &a.info)
	if err != nil {
		panic(err)
	}
}

// WailsInit is the init fuction for the wails runtime
func (a *API) WailsInit(runtime *wails.Runtime) error {
	a.browser = runtime.Browser
	a.window = runtime.Window
	a.logger = runtime.Log.New("API")

	etwDir = a.getExecDirectory() + "/"
	a.logger.Infof("ETW/Current directory: %s", etwDir)

	// appDataDir = Sh.Getenv("APPDATA") + "appDataPath"
	appDataDir = etwDir + "appDataFolder/" + appDataPath
	a.logger.Infof("AppData directory: %s", appDataDir)

	a.loadInfoFromFile()
	a.logger.Infof("Info loaded %v", a.info)

	return nil
}

// WailsShutdown is the shutdown function that is called when wails shuts down
func (a *API) WailsShutdown() {
}
