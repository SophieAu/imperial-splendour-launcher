package backend

func (a *API) IsActive() bool {
	a.logger.Infof("IsActive: %s", a.info.IsActive)
	return a.info.IsActive
}
