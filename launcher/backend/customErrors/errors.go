package customErrors

import "errors"

// Switch Errors
var FileList = errors.New("FileListError")
var Activation = errors.New("ActivationError")
var Rollback = errors.New("RollbackError")
var Deactivation = errors.New("DeactivationError")
var StatusUpdate = errors.New("StatusUpdateError")

var Uninstall = errors.New("UninstallError")
var Website = errors.New("WebsiteError")
var Play = errors.New("PlayError")



