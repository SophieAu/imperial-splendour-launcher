package backend_test

import (
	"testing"
)

func TestAll(t *testing.T) {
	t.Run("Ensure Init works Properly", TestInit)
	t.Run("Ensure Shutdown works Properly", TestShutdown)

	t.Run("Test Version Endpoint", TestVersion)
	t.Run("Test IsActive Endpoint", TestIsActive)
	t.Run("Test Play Endpoint", TestPlay)
	t.Run("Test Switch Endpoint", TestSwitch)
	t.Run("Test GoToWebsite Endpoint", TestGoToWebsite)
	t.Run("Test Uninstall Endpoint", TestUninstall)
	t.Run("Test Exit Endpoint", TestExit)
}
