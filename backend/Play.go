package backend

import "imperial-splendour-launcher/backend/customErrors"

func (a *API) Play() error {
	a.logger.Infof("Starting Game")

	if err := a.browser.OpenURL(etwSteamURI); err != nil {
		a.logger.Warnf("%v", err)
		return customErrors.Play
	}

	a.window.Close()
	return nil
}
