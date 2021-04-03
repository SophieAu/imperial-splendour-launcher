package backend

import (
	"imperial-splendour-launcher/backend/customErrors"
)

func (a *API) Uninstall() error {
	a.logger.Info("Uninstalling")

	if a.info.IsActive {
		if err := a.deactivateImpSplen(); err != nil {
			a.logger.Warnf("%v", err)
			return customErrors.Deactivation
		}
	}

	if err := a.deleteAllFiles(); err != nil {
		a.logger.Warnf("%v", err)
		return customErrors.Uninstall
	}

	// TODO: run script to delete self?
	// TODO: delete shortcuts
	return nil
}
