package backend

import (
	"io"
	"net/http"
	"os"
)

const (
	deactivatorUrl = "https://github.com/SophieAu/imperial-splendour-launcher/raw/master/artifacts/deactivator.exe"
	launcherUrl    = "https://github.com/SophieAu/imperial-splendour-launcher/raw/master/artifacts/ImperialSplendour.exe"
	setupUrl       = "https://github.com/SophieAu/imperial-splendour-launcher/raw/master/setup/setupBundled.iss"
	appiconUrl     = "https://github.com/SophieAu/imperial-splendour-launcher/raw/master/setup/appicon.ico"
)

func (a *API) downloadFiles() error {
	appiconTarget := a.setupBaseFolder + "/" + appicon
	setupTarget := a.setupBaseFolder + "/" + setupFile
	launcherTarget := a.setupBaseFolder + "/" + tempPath + launcherFile
	deactivatorTarget := a.setupBaseFolder + "/" + tempPath + uninstallPath + deactivatorFile

	if err := DownloadFile(appiconTarget, appiconUrl); err != nil {
		return err
	}
	if err := DownloadFile(setupTarget, setupUrl); err != nil {
		return err
	}
	if err := DownloadFile(launcherTarget, launcherUrl); err != nil {
		return err
	}
	if err := DownloadFile(deactivatorTarget, deactivatorUrl); err != nil {
		return err
	}

	return nil
}

func DownloadFile(url string, targetFilePath string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(targetFilePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
