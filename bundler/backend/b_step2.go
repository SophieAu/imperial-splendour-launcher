package backend

import "imperial-splendour-bundler/backend/customErrors"

func (a *API) prepareUserScript(sourcePath string) error {
	source := sourcePath + userScript

	if doesUserScriptExist, _ := a.Sh.DoesFileExist(source); !doesUserScriptExist {
		return a.error("Userscript not found in "+source+".", customErrors.UserScriptMissing)
	}

	destination := a.setupBaseFolder + tempPath + modPath + userScript

	a.logger.Info("Moving from " + source + " to " + destination)
	err := a.Sh.MoveFile(source, destination)
	if err != nil {
		return a.error("Cannot move userScript: "+err.Error(), customErrors.MoveUserScript)
	}
	a.logToFrontend("User script was moved to mod folder")
	return nil
}
