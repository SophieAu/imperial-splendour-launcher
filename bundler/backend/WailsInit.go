package backend

import (
	"github.com/wailsapp/wails"
)

func (a *API) ensureInnoSetup() error {
	return a.Sh.RunCommand("/bin/sh", "-c", "command -v iscc")
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
