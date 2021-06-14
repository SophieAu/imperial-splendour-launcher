package customErrors

import "errors"

var InnoSetup = errors.New("InnoSetupError")
var EmptySource = errors.New("EmptySourceError")
var InvalidVersion = errors.New("InvalidVersionError")

var TempFolderCreation = errors.New("TempFolderCreationError")
var ReadSourceDir = errors.New("ReadSourceDirError")
var ReadFileList = errors.New("ReadFileListError")
var FileMissing = errors.New("FileMissingError")
var MoveFile = errors.New("MoveFileError")
var SaveFileList = errors.New("SaveFileListError")

var UserScriptMissing = errors.New("UserScriptMissingError")
var MoveUserScript = errors.New("MoveUserScriptError")

var Download = errors.New("DownloadError")
var InfoFile = errors.New("InfoFileError")
var VersionUpdate = errors.New("VersionUpdateError")

var NoFileList = errors.New("NoFileListError")

var CompileSetup = errors.New("CompileSetupError")
var ZipFiles = errors.New("ZipFilesError")
