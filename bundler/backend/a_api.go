package backend

type API struct {
	logger          Logger
	browser         Browser
	window          Window
	Sh              Handler
	logStore        Store
	userSettings    UserSettings
	setupBaseFolder string
}

type UserSettings struct {
	sourcePath      string
	versionNumber   string
	packageRawFiles bool
	fileListPath    string
}

type Info struct {
	IsActive           bool   `json:"isActive"`
	Version            string `json:"version"`
	UserScriptChecksum string `json:"usChecksum"`
}

func (a *API) logToFrontend(s string) {
	a.logger.Info(s)

	a.logStore.Update(func(currentValue []string) []string {
		return append(currentValue, s)
	})
}

func (a *API) error(warning string, err error) error {
	a.logger.Warn(warning)
	return err
}
