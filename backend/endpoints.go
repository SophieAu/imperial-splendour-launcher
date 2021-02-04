package backend

// Version .
func (a *API) Version() string {
	a.logger.Infof("Version: %s", a.info.Version)
	return a.info.Version
}

// IsActive .
func (a *API) IsActive() bool {
	a.logger.Infof("IsActive: %s", a.info.IsActive)
	return a.info.IsActive
}

// Play .
func (a *API) Play() error {
	err := a.browser.OpenURL(etwSteamURI)
	if err != nil {
		return err
	}

	a.Exit()
	return nil
}

// Switch .
func (a *API) Switch() error {
	var switchFn func() error
	if a.info.IsActive {
		switchFn = a.deactivateImpSplen
	} else {
		switchFn = a.activateImpSplen
	}

	err := switchFn()
	if err != nil {
		return err
	}

	return nil
}

// GoToWebsite .
func (a *API) GoToWebsite() error {
	a.logger.Infof("Navigating to %s", websiteURL)

	err := a.browser.OpenURL(websiteURL)
	if err != nil {
		a.logger.Errorf("Could not open website: %v", err)
		return err
	}
	return nil
}

// Uninstall .
func (a *API) Uninstall() error {
	a.logger.Info("Uninstalling")
	return nil
}

// Exit .
func (a *API) Exit() {
	a.window.Close()
}
