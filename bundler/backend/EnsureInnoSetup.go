package backend

import (
	"imperial-splendour-bundler/backend/customErrors"
)

func (a *API) EnsureInnoSetup() error {

	err := a.Sh.StartCommand("get-command iscc")
	if err != nil {
		return a.error("InnoSetup not installed", customErrors.InnoSetup)
	}

	return nil
}
