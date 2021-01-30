package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func main() {
	js := mewn.String("./frontend/public/build/bundle.js")
	css := mewn.String("./frontend/public/build/bundle.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:     1280,
		Height:    800,
		Resizable: false,
		Title:     "Imperial Splendour: Rise of the Republic",
		JS:        js,
		CSS:       css,
		Colour:    "#131313",
	})

	app.Bind(&API{})
	app.Run()
}
