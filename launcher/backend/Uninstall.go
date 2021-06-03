package backend

import (
	"imperial-splendour-launcher/backend/customErrors"
)

func (a *API) Uninstall() error {
	a.logger.Info("Uninstalling")

	if err := a.Sh.StartCommand(a.dirs.etw + uninstallerFile); err != nil {
		return customErrors.Uninstall
	}

	a.window.Close()
	return nil
}
