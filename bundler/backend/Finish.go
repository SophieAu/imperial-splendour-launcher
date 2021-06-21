package backend

func (a *API) Exit() {
	a.Sh.Exit(0)
}
