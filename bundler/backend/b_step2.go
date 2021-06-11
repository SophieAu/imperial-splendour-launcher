package backend

func (a *API) ensureUserScript() bool {
	doesUserScriptExist, _ := a.Sh.DoesFileExist(a.userSettings.sourcePath + "/" + userScript)

	return doesUserScriptExist
}

func (a *API) moveUserScriptIntoModFolder() error {
	source := a.userSettings.sourcePath + "/" + userScript
	destination := a.setupBaseFolder + tempPath + modPath + userScript

	a.logger.Debugf("Moving from %s to %s", source, destination)
	err := a.Sh.MoveFile(source, destination)
	return err
}
