package backend

import "imperial-splendour-launcher/backend/customErrors"

func (a *API) GoToWebsite() error {
	a.logger.Infof("Navigating to %s", websiteURL)

	if err := a.browser.OpenURL(websiteURL); err != nil {
		a.logger.Warnf("Could not open website: %v", err)
		return customErrors.Website
	}
	return nil
}
