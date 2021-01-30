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

	err := a.moveFile()
	if err != nil {
		return err
	}
	return nil
}

// GoToWebsite .
func (a *API) GoToWebsite() {
	err := a.runtime.Browser.OpenURL("kdlfjhttps://imperialsplendour.com/")
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
