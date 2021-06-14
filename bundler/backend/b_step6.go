package backend

import "imperial-splendour-bundler/backend/customErrors"

func (a *API) compileSetup() error {
	err := a.Sh.RunCommand("iscc", a.setupBaseFolder+setupFile)
	if err != nil {
		return a.error("Error compiling the setup", customErrors.CompileSetup)
	}
	return nil
}
