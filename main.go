package main

import (
	"github.com/leaanthony/mewn"
	"github.com/pkg/browser"
	"github.com/wailsapp/wails"
)

func play(exitCall *ExitCall) func() error {
	playFunc := func() error {
		err := browser.OpenURL("steam://rungameid/10500")
		if err != nil {
			return err
		}
		exitCall.Exit()
		return nil
	}
	return playFunc
}

func switchMode() {}
func uninstall()  {}

func openWebsite() {
	browser.OpenURL("https://imperialsplendour.com/")
}

func exit(exitCall *ExitCall) func() {
	exitFunc := func() {
		exitCall.Exit()
	}
	return exitFunc
}

func version() string {
	return "2.0"
}

func main() {
	js := mewn.String("./frontend/public/build/bundle.js")
	css := mewn.String("./frontend/public/build/bundle.css")

	exitCall := &ExitCall{}

	app := wails.CreateApp(&wails.AppConfig{
		Width:     1280,
		Height:    800,
		Resizable: false,
		Title:     "Imperial Splendour: Rise of the Republic",
		JS:        js,
		CSS:       css,
		Colour:    "#131313",
	})

	app.Bind(play(exitCall))
	app.Bind(switchMode)
	app.Bind(uninstall)
	app.Bind(openWebsite)
	app.Bind(exit(exitCall))
	app.Bind(version)
	app.Run()
}
