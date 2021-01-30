package main

// Play .
func (a *API) Play() error {
	err := a.runtime.Browser.OpenURL("steam://rungameid/10500")
	if err != nil {
		return err
		// return thrown promise here asking the user if they have steam installed
	}
	a.Exit()
	return nil
}

// Switch .
func (a *API) Switch() error {
	a.logger.Info("Switching")
	var err error

	if a.info.IsActive {
		err = a.deactivateImpSplen()
	} else {
		err = a.activateImpSplen()
	}
	if err != nil {
		return err
	}
	return nil
}

// GoToWebsite .
func (a *API) GoToWebsite() {
	err := a.runtime.Browser.OpenURL("https://imperialsplendour.com/")
	if err != nil {
		a.logger.Errorf("%v", err)
	}
}

// Uninstall .
func (a *API) Uninstall() {
	a.logger.Info("Uninstalling")
}

// Exit .
func (a *API) Exit() {
	a.runtime.Window.Close()
}

// Version .
func (a *API) Version() string {
	return a.info.Version
}

// IsActive .
func (a *API) IsActive() bool {
	a.logger.Infof("Version: %s", a.info.Version)
	return a.info.IsActive
}

// STARTUP CHECKS
// * is mod active
// * is in correct folder

// WHERE TO MOVE WHAT
/*
reading from file list

data files
from mod-diretory




*/
