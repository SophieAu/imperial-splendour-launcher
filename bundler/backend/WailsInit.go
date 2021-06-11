package backend

import (
	"imperial-splendour-bundler/backend/customErrors"

	"github.com/wailsapp/wails"
)

func (a *API) ensureInnoSetup() error {
	err := a.Sh.RunCommand("/bin/sh", "-c", "command -v iscc")
	if err != nil {
		return a.error("InnoSetup not installed", customErrors.InnoSetup)
	}
	return nil
}

func (a *API) Init(browser Browser, window Window, logger Logger, systemHandler Handler) error {
	a.browser = browser
	a.window = window
	a.logger = logger
	a.Sh = systemHandler

	return a.ensureInnoSetup()
}

func (a *API) WailsInit(runtime *wails.Runtime) error {
	return a.Init(runtime.Browser, runtime.Window, runtime.Log.New("API"), &SystemHandler{})
}
