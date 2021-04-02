package backend

func (a *API) Play() error {
	a.logger.Infof("Starting Game")

	if err := a.browser.OpenURL(etwSteamURI); err != nil {
		a.logger.Warnf("%v", err)
		return err
	}

	a.window.Close()
	return nil
}
