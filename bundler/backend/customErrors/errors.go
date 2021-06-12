package customErrors

import "errors"

var InnoSetup = errors.New("InnoSetupError")
var EmptySource = errors.New("EmptySourceError")
var InvalidVersion = errors.New("InvalidVersionError")

var TempFolderCreation = errors.New("TempFolderCreationError")

var NoFileList = errors.New("NoFileListError")
