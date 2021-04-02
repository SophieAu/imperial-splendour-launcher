package backend

import "errors"

func (a *API) Uninstall() error {
	a.logger.Info("Uninstalling")

	if a.info.IsActive {
		if err := a.deactivateImpSplen(); err != nil {
			a.logger.Warnf("%v", err)
			return errors.New("Could not uninstall")
		}
	}

	err := a.deleteAllFiles()

	// TODO: run script to delete self?
	// TODO: delete shortcuts
	return err
}
