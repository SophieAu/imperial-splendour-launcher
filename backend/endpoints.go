package backend

func (a *API) Version() string {
	a.logger.Infof("Version: %s", a.info.Version)
	return a.info.Version
}

func (a *API) IsActive() bool {
	a.logger.Infof("IsActive: %s", a.info.IsActive)
	return a.info.IsActive
}

func (a *API) Play() error {
	err := a.browser.OpenURL(etwSteamURI)
	if err != nil {
		return err
	}

	a.Exit()
	return nil
}

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

func (a *API) GoToWebsite() error {
	a.logger.Infof("Navigating to %s", websiteURL)

	err := a.browser.OpenURL(websiteURL)
	if err != nil {
		a.logger.Errorf("Could not open website: %v", err)
		return err
	}
	return nil
}

func (a *API) Uninstall() error {
	a.logger.Info("Uninstalling")
	return nil
}

func (a *API) Exit() {
	a.window.Close()
}
