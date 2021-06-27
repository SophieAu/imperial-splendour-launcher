package backend

import (
	"github.com/wailsapp/wails"
)

func (a *API) Init(browser Browser, window Window, logger Logger, store Store, dialog Dialog, systemHandler Handler) error {
	a.browser = browser
	a.window = window
	a.logger = logger
	a.dialog = dialog
	a.Sh = systemHandler
	a.logStore = store

	return nil
}

func (a *API) WailsInit(runtime *wails.Runtime) error {
	return a.Init(runtime.Browser, runtime.Window, runtime.Log.New("API"), runtime.Store.New("Log", []string{}), runtime.Dialog, &SystemHandler{})
}
