package backend

import "imperial-splendour-launcher/backend/customErrors"

func (a *API) GoToDownloadPage() error {
	a.logger.Infof("Navigating to %s", downloadPageUrl)

	if err := a.browser.OpenURL(downloadPageUrl); err != nil {
		a.logger.Warnf("Could not open website: %v", err)
		return customErrors.Website
	}
	return nil
}
