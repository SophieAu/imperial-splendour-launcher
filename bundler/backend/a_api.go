package backend

type API struct {
	logger          Logger
	browser         Browser
	window          Window
	Sh              Handler
	userSettings    UserSettings
	setupBaseFolder string
}

type UserSettings struct {
	sourcePath             string
	versionNumber          string
	packageRawFiles        bool
	comparisonFileListPath string
}

type Info struct {
	IsActive           bool   `json:"isActive"`
	Version            string `json:"version"`
	UserScriptChecksum string `json:"usChecksum"`
}
