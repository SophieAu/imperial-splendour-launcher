package backend

import (
	"imperial-splendour-bundler/backend/customErrors"
	"os/exec"
)

func (a *API) EnsureInnoSetup() error {

	path, err := exec.LookPath("iscc")
	if err != nil {
		a.logger.Warn("installing fortune is in your future")
	}
	a.logger.Infof("fortune is available at %s\n", path)

	err = a.Sh.StartCommand("iscc /?")
	if err != nil {
		return a.error("InnoSetup not installed", customErrors.InnoSetup)
	}

	return nil
}
