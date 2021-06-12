package backend

import (
	"imperial-splendour-bundler/backend/customErrors"
	"regexp"
)

func (a *API) validateUserInput(sourcePath string, versionNumber string, fileListPath string) error {
	if sourcePath == "" {
		return a.error("Empty source path", customErrors.EmptySource)
	}
	doesMatchRegex, err := regexp.MatchString(versionRegex, versionNumber)
	if !doesMatchRegex || err != nil {
		return a.error("Invalid version number", customErrors.InvalidVersion)
	}
	if fileListPath == "" {
		return a.error("No file list provided", customErrors.NoFileList)
	}

	return nil
}
