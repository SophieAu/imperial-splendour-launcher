package main

import (
	"deactivator/app"
	"os"
)

func main() {
	a := &app.API{}

	if err := a.Init(&app.SystemHandler{}); err != nil {
		os.Exit(1)
	}

	if err := a.Deactivate(); err != nil {
		os.Exit(1)
	}
}
