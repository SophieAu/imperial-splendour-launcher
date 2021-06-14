package backend

import "imperial-splendour-bundler/backend/customErrors"

func (a *API) prepareUserScript(sourcePath string) error {
	if doesUserScriptExist, _ := a.Sh.DoesFileExist(sourcePath + "/" + userScript); !doesUserScriptExist {
		return a.error("Userscript not found.", customErrors.UserScriptMissing)
	}

	source := sourcePath + "/" + userScript
	destination := a.setupBaseFolder + tempPath + modPath + userScript

	a.logger.Info("Moving from " + source + " to " + destination)
	err := a.Sh.MoveFile(source, destination)
	if err != nil {
		return a.error("Cannot Move userScript: "+err.Error(), customErrors.MoveUserScript)
	}
	return nil
}
