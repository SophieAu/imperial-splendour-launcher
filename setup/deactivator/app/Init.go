package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

func (a *API) getExecDirectory() (string, error) {
	ex, err := a.Sh.Executable()
	return strings.TrimSuffix(filepath.ToSlash(filepath.Dir(ex)), "/"+strings.TrimSuffix(uninstallPath, "/")) + "/", err
}

func (a *API) loadInfoFromFile() error {
	byteValue, err := a.Sh.ReadFile(a.dirs.etw + infoFile)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(byteValue, &a.info); err != nil {
		return err
	}

	if a.info.UserScriptChecksum == "" || a.info.Version == "" {
		return errors.New("Corrupt Info File")
	}

	return nil
}

func (a *API) Init(sytemHandler Handler) error {
	a.Sh = sytemHandler

	etwDir, err := a.getExecDirectory()
	if err != nil {
		return err
	}
	a.dirs.etw = etwDir
	log.Printf("ETW/Current directory: %s", a.dirs.etw)

	appDataDir := a.Sh.Getenv("APPDATA")
	if appDataDir == "" {
		return errors.New("Couldn't get user's APPDATA dir")
	}
	a.dirs.appData = appDataDir + "/" + appDataPath
	log.Printf("AppData directory: %s", a.dirs.appData)

	if err = a.loadInfoFromFile(); err != nil {
		log.Printf("%v", err)
		return err
	}
	log.Printf("Info loaded %v", fmt.Sprintf("%v", a.info))

	return nil
}
