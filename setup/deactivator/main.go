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

	// TODO: Add flag for hard failure
	if err := a.Deactivate(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
