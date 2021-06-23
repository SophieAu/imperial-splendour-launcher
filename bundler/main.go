package main

import (
	_ "embed"
	"imperial-splendour-bundler/backend"

	"github.com/wailsapp/wails"
)

//go:embed frontend/public/build/bundle.js
var js string

//go:embed frontend/public/build/bundle.css
var css string

func main() {
	app := wails.CreateApp(&wails.AppConfig{
		Width:            540,
		Height:           420,
		Resizable:        false,
		Title:            backend.AppName,
		JS:               js,
		CSS:              css,
		DisableInspector: true,
	})

	app.Bind(&backend.API{})
	_ = app.Run()
}
