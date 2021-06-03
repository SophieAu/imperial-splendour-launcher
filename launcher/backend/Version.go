package backend

func (a *API) Version() string {
	a.logger.Infof("Version: %s", a.info.Version)
	return a.info.Version
}
