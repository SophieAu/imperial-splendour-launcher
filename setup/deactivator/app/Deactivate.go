package app

func (a *API) Deactivate(isStrict bool) error {
	if !a.info.IsActive {
		return nil
	}

	return a.deactivateImpSplen(isStrict)
}
