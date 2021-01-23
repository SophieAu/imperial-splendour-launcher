package main

import (
	"github.com/wailsapp/wails"
)

// ExitCall .
type ExitCall struct {
	runtime *wails.Runtime
	log     *wails.CustomLogger
}

// Exit .
func (t *ExitCall) Exit() {
	t.runtime.Window.Close()
}

// WailsInit .
func (t *ExitCall) WailsInit(runtime *wails.Runtime) error {
	t.runtime = runtime
	return nil
}

// WailsShutdown .
func (t *ExitCall) WailsShutdown() {
	t.log.Info("WailsShutdown() Called!")
}
