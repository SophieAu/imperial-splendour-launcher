package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

var appName = "Imperial Splendour: Rise of the Republic"

var appDataDir = ""
var appDataPath = "The Creative Assembly/Empire/scripts/"

var etwDir = ""
var modPath = "IS_Files/"
var dataPath = "data/"
var campaignPath = dataPath + "campaigns/imperial_splendour/"

var userScript = "user.empire_script.txt"
var fileListFile = "IS_FileList.txt"
var infoFile = "IS_info.json"

var websiteURL = "https://imperialsplendour.com/"
var etwSteamURI = "steam://rungameid/10500"

func main() {
	js := mewn.String("./frontend/public/build/bundle.js")
	css := mewn.String("./frontend/public/build/bundle.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:            1280,
		Height:           800,
		Resizable:        false,
		Title:            appName,
		JS:               js,
		CSS:              css,
		DisableInspector: true,
	})

	app.Bind(&API{fh: &FileHandler{}})
	app.Run()
}
