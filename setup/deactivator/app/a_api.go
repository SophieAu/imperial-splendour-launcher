package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
)

type API struct {
	Sh   Handler
	dirs struct {
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
	log.Printf("Moving from %s to %s", source, destination)
	return a.Sh.MoveFile(source, destination)
}

func (a *API) deactivateImpSplen(isStrict bool) error {
	log.Print("Deactivating ImpSplen")

	files, err := a.readFileList()
	if err != nil {
		log.Printf("%v", err)
		return errors.New("FileListError")
	}

	hasError := false
	log.Print("Moving data files")
	for _, file := range files.dataFiles {
		if err := a.moveFile(a.dirs.etw+dataPath+file, a.dirs.etw+modPath+file); err != nil {
			log.Printf("%v", err)
			hasError = true
		}
	}

	log.Print("Moving campaign files")
	for _, file := range files.campaignFiles {
		if err := a.moveFile(a.dirs.etw+campaignPath+file, a.dirs.etw+modPath+file); err != nil {
			log.Printf("%v", err)
			hasError = true
		}
	}

	log.Print("Moving User Script")
	if err = a.moveFile(a.dirs.appData+userScript, a.dirs.etw+modPath+userScript); err != nil {
		log.Printf("%v", err)
		hasError = true
	}

	if err := a.setStatus(false); err != nil && isStrict {
		log.Printf("%v", err)
		return errors.New("StatusUpdateError")
	}

	log.Print("ImpSplen deactivated")
	if hasError && isStrict {
		return errors.New("DeactivationError")
	}
	return nil
}
