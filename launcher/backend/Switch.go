package backend

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
