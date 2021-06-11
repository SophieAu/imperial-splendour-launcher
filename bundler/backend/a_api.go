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

func (a *API) error(warning string, err error) error {
	a.logger.Warn(warning)

	if err == nil {
		err = errors.New(warning)
	}
	return err
}
