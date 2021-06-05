package main

import (
	"deactivator/app"
	"flag"
	"os"
)

func main() {
	isStrict := flag.Bool("strict", false, "enable strict-mode, i.e. fail on deactivation errors")
	flag.Parse()

	a := &app.API{}
	if err := a.Init(&app.SystemHandler{}); err != nil {
		os.Exit(12)
	}

	// TODO: Add flag for hard failure
	if err := a.Deactivate(*isStrict); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
