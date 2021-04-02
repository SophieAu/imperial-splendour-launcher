package backend

import "errors"

func (a *API) Version() string {
	a.logger.Infof("Version: %s", a.info.Version)
	return a.info.Version
}

func (a *API) IsActive() bool {
	a.logger.Infof("IsActive: %s", a.info.IsActive)
	return a.info.IsActive
}

func (a *API) Play() error {
	a.logger.Infof("Starting Game")

	if err := a.browser.OpenURL(etwSteamURI); err != nil {
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

	if err := switchFn(); err != nil {
		a.logger.Warnf("%v", err)
		return err
	}

	return nil
}

func (a *API) GoToWebsite() error {
	a.logger.Infof("Navigating to %s", websiteURL)

	if err := a.browser.OpenURL(websiteURL); err != nil {
		a.logger.Warnf("Could not open website: %v", err)
		return err
	}
	return nil
}

func (a *API) Uninstall() error {
	a.logger.Info("Uninstalling")

	if a.info.IsActive {
		if err := a.deactivateImpSplen(); err != nil {
			a.logger.Warnf("%v", err)
			return errors.New("Could not uninstall")
		}
	}

	err := a.deleteAllFiles()

	// TODO: run script to delete self?
	// TODO: delete shortcuts
	return err
}

func (a *API) Exit() {
	a.window.Close()
}
