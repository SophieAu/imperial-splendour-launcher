package backend

import "imperial-splendour-bundler/backend/customErrors"

const (
	deactivatorUrl = "https://github.com/SophieAu/imperial-splendour-launcher/raw/master/artifacts/deactivator.exe"
	launcherUrl    = "https://github.com/SophieAu/imperial-splendour-launcher/raw/master/artifacts/ImperialSplendour.exe"
	setupUrl       = "https://github.com/SophieAu/imperial-splendour-launcher/raw/master/setup/setupBundled.iss"
	appiconUrl     = "https://github.com/SophieAu/imperial-splendour-launcher/raw/master/setup/appicon.ico"
)

func (a *API) downloadFiles() error {
	appiconTarget := a.setupBaseFolder + appicon
	setupTarget := a.setupBaseFolder + setupFile
	launcherTarget := a.setupBaseFolder + tempPath + launcherFile
	deactivatorTarget := a.setupBaseFolder + tempPath + uninstallPath + deactivatorFile

	if err := a.Sh.DownloadFile(appiconUrl, appiconTarget); err != nil {
		return a.error("Cannot download Appicon", customErrors.Download)
	}
	a.logToFrontend("App icon was downloaded")
	if err := a.Sh.DownloadFile(setupUrl, setupTarget); err != nil {
		return a.error("Cannot download setup script", customErrors.Download)
	}
	a.logToFrontend("Setup script was downloaded")
	if err := a.Sh.DownloadFile(launcherUrl, launcherTarget); err != nil {
		return a.error("Cannot download launcher", customErrors.Download)
	}
	a.logToFrontend("Launcher was downloaded")
	if err := a.Sh.DownloadFile(deactivatorUrl, deactivatorTarget); err != nil {
		return a.error("Cannot download deactivator", customErrors.Download)
	}
	a.logToFrontend("Deactivator was downloaded")
	return nil
}
