package backend

import (
	"encoding/json"
	"errors"
	"path/filepath"

	"github.com/wailsapp/wails"
)

func (a *API) getExecDirectory() (string, error) {
	ex, err := a.Sh.Executable()
	return filepath.Dir(ex) + "/", err
}

func (a *API) loadInfoFromFile() error {
	byteValue, err := a.Sh.ReadFile(etwDir + modPath + infoFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteValue, &a.info)
	if err != nil {
		return err
	}

	if a.info.UserScriptChecksum == "" || a.info.Version == "" {
		return errors.New("Corrupt Info File")
	}

	return nil
}

func (a *API) Init(browser Browser, window Window, logger Logger, systemHandler Handler) error {
	a.browser = browser
	a.window = window
	a.logger = logger
	a.Sh = systemHandler

	etwDir, err := a.getExecDirectory()
	if err != nil {
		return err
	}
	a.logger.Infof("ETW/Current directory: %s", etwDir)

	// appDataDir = Sh.Getenv("APPDATA") + "appDataPath"
	appDataDir = etwDir + "appDataFolder/" + appDataPath
	a.logger.Infof("AppData directory: %s", appDataDir)

	err = a.loadInfoFromFile()
	if err != nil {
		return err
	}
	a.logger.Infof("Info loaded %v", a.info)

	return nil
}

func (a *API) WailsInit(runtime *wails.Runtime) error {
	return a.Init(runtime.Browser, runtime.Window, runtime.Log.New("API"), &SystemHandler{})
}

func (a *API) WailsShutdown() {
}
