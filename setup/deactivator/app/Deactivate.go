package app

func (a *API) Deactivate() error {
	if !a.info.IsActive {
		return nil
	}

	return a.deactivateImpSplen()
}
