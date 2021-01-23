package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

var wailsRuntime *wails.Runtime

func version() string {
	return "2.0"
}

func main() {
	js := mewn.String("./frontend/public/build/bundle.js")
	css := mewn.String("./frontend/public/build/bundle.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "Imperial Splendour Launcher",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})

	app.Bind(version)
	app.Bind(&ExitCall{})
	app.Run()
}
