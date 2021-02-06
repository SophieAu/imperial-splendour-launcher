package main

import (
	"imperial-splendour-launcher/backend"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func main() {
	js := mewn.String("./frontend/public/build/bundle.js")
	css := mewn.String("./frontend/public/build/bundle.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:            1280,
		Height:           800,
		Resizable:        false,
		Title:            backend.AppName,
		JS:               js,
		CSS:              css,
		DisableInspector: true,
	})

	app.Bind(&backend.API{})
	_ = app.Run()
}
