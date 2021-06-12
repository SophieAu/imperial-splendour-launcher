package backend

import "errors"

type API struct {
	logger          Logger
	browser         Browser
	window          Window
	Sh              Handler
	userSettings    UserSettings
	logString       []string
	setupBaseFolder string
}

type UserSettings struct {
	sourcePath      string
	versionNumber   string
	packageRawFiles bool
	fileListPath    string
}

type Info struct {
	IsActive           bool   `json:"isActive"`
	Version            string `json:"version"`
	UserScriptChecksum string `json:"usChecksum"`
}

func (a *API) logToFrontend(s string) {
	a.logger.Info(s)
	a.logString = append(a.logString, s)
}

func (a *API) error(warning string, err error) error {
	a.logger.Warn(warning)

	if err == nil {
		err = errors.New(warning)
	}
	a.logger.Warn(warning)
	return err
}
