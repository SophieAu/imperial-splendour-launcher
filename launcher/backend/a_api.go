package backend

import (
	"encoding/json"
	"errors"
	"fmt"
	"imperial-splendour-launcher/backend/customErrors"
	"strings"
)

const (
	AppName = "Imperial Splendour: Rise of the Republic"

	appDataPath   = "The Creative Assembly/Empire/scripts/"
	modPath       = "IS_Files/"
	dataPath      = "data/"
	campaignPath  = dataPath + "campaigns/imperial_splendour/"
	uninstallPath = "IS_Uninstall/"

	userScript      = "user.empire_script.txt"
	fileListFile    = modPath + "IS_FileList.txt"
	infoFile        = modPath + "IS_Info.json"
	uninstallerFile = uninstallPath + "unins000.exe"

	websiteURL  = "https://imperialsplendour.com/"
	etwSteamURI = "steam://rungameid/10500"
)

type API struct {
	logger  Logger
	browser Browser
	window  Window
	Sh      Handler
	dirs    struct {
		etw     string
		appData string
	}
	info info
}

type modFiles struct {
	dataFiles     []string
	campaignFiles []string
}

type info struct {
	IsActive           bool   `json:"isActive"`
	Version            string `json:"version"`
	UserScriptChecksum string `json:"usChecksum"`
}

type fileTuple struct {
	file string
	dir  string
}

func (a *API) setStatus(isActive bool) error {
	newInfo := a.info
	newInfo.IsActive = isActive

	newInfoJSON, err := json.MarshalIndent(newInfo, "", "\t")
	if err != nil {
		return err
	}

	if err = a.Sh.WriteFile(a.dirs.etw+infoFile, newInfoJSON); err != nil {
		return err
	}

	a.info.IsActive = isActive
	return nil
}

func (a *API) readFileList() (*modFiles, error) {
	modFiles := modFiles{
		dataFiles:     []string{},
		campaignFiles: []string{},
	}

	fileBlob, err := a.Sh.ReadFile(a.dirs.etw + fileListFile)
	if err != nil {
		return nil, err
	}

	var errMsg string
	for _, file := range strings.Split(string(fileBlob), "\n") {
		if strings.HasSuffix(file, ".pack") {
			modFiles.dataFiles = append(modFiles.dataFiles, file)
		} else if strings.HasSuffix(file, ".tga") || strings.HasSuffix(file, ".esf") || strings.HasSuffix(file, ".lua") {
			modFiles.campaignFiles = append(modFiles.campaignFiles, file)
		} else if file != "" {
			errMsg = errMsg + fmt.Sprintf("Unknown file '%s' found in file list\n", file)
		}
	}
	if errMsg != "" {
		return nil, errors.New(errMsg)
	}
	return &modFiles, nil
}

func (a *API) moveFile(source, destination string) error {
	a.logger.Debugf("Moving from %s to %s", source, destination)
	return a.Sh.MoveFile(source, destination)
}

func (a *API) rollbackActivation(fileList []fileTuple) error {
	hasError := false
	for _, tuple := range fileList {
		if err := a.moveFile(tuple.dir+tuple.file, a.dirs.etw+modPath+tuple.file); err != nil {
			a.logger.Warnf("%v", err)
			hasError = true
		}
	}

	if hasError {
		return customErrors.Rollback
	}
	return nil
}

func (a *API) activateFile(fileName string, targetDir string, activeFiles *[]fileTuple) error {
	if err := a.moveFile(a.dirs.etw+modPath+fileName, targetDir+fileName); err != nil {
		a.logger.Warnf("%v", err)
		if err = a.rollbackActivation(*activeFiles); err != nil {
			return err
		}
		return customErrors.Activation
	}

	*activeFiles = append(*activeFiles, fileTuple{fileName, targetDir})
	return nil
}

func (a *API) activateImpSplen() error {
	a.logger.Debug("Activating ImpSplen")
	var filesMoved []fileTuple

	files, err := a.readFileList()
	if err != nil {
		a.logger.Warnf("%v", err)
		return customErrors.FileList
	}

	a.logger.Debug("Moving data files")
	for _, file := range (*files).dataFiles {
		if err := a.activateFile(file, a.dirs.etw+dataPath, &filesMoved); err != nil {
			return err
		}
	}

	a.logger.Debug("Moving campaign files")
	for _, file := range (*files).campaignFiles {
		if err := a.activateFile(file, a.dirs.etw+campaignPath, &filesMoved); err != nil {
			return err
		}
	}

	a.logger.Debug("Moving User Script")
	if err = a.activateFile(userScript, a.dirs.appData, &filesMoved); err != nil {
		return err
	}

	if err = a.setStatus(true); err != nil {
		a.logger.Warnf("%v", err)
		if err = a.rollbackActivation(filesMoved); err != nil {
			return err
		}
		return customErrors.StatusUpdate
	}
	a.logger.Debug("ImpSplen activated")
	return nil
}

func (a *API) deactivateImpSplen() error {
	a.logger.Debug("Deactivating ImpSplen")

	files, err := a.readFileList()
	if err != nil {
		a.logger.Warnf("%v", err)
		return customErrors.FileList
	}

	hasError := false
	a.logger.Debug("Moving data files")
	for _, file := range files.dataFiles {
		if err := a.moveFile(a.dirs.etw+dataPath+file, a.dirs.etw+modPath+file); err != nil {
			a.logger.Warnf("%v", err)
			hasError = true
		}
	}

	a.logger.Debug("Moving campaign files")
	for _, file := range files.campaignFiles {
		if err := a.moveFile(a.dirs.etw+campaignPath+file, a.dirs.etw+modPath+file); err != nil {
			a.logger.Warnf("%v", err)
			hasError = true
		}
	}

	a.logger.Debug("Moving User Script")
	if err = a.moveFile(a.dirs.appData+userScript, a.dirs.etw+modPath+userScript); err != nil {
		a.logger.Warnf("%v", err)
		hasError = true
	}

	if err := a.setStatus(false); err != nil {
		a.logger.Warnf("%v", err)
		return customErrors.StatusUpdate
	}

	a.logger.Debug("ImpSplen deactivated")
	if hasError {
		return customErrors.Deactivation
	}
	return nil
}
