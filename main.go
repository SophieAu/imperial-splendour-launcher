package main

import (
	"github.com/leaanthony/mewn"
	"github.com/pkg/browser"
	"github.com/wailsapp/wails"
)

func play()       {}
func switchMode() {}
func uninstall()  {}

func openWebsite() {
	err := browser.OpenURL("https://imperialsplendour.com/")
	print(err)
}

func version() string {
	return "2.0"
}

func main() {
	js := mewn.String("./frontend/public/build/bundle.js")
	css := mewn.String("./frontend/public/build/bundle.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1920,
		Height: 1200,
		Title:  "Imperial Splendour: Rise of the Republic",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})

	app.Bind(play)
	app.Bind(switchMode)
	app.Bind(uninstall)
	app.Bind(openWebsite)
	app.Bind(&ExitCall{})
	app.Bind(version)
	app.Run()
}
