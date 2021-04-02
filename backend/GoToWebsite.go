package backend

func (a *API) GoToWebsite() error {
	a.logger.Infof("Navigating to %s", websiteURL)

	if err := a.browser.OpenURL(websiteURL); err != nil {
		a.logger.Warnf("Could not open website: %v", err)
		return err
	}
	return nil
}
