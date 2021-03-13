package backend_test

import (
	"testing"
)

func TestAll(t *testing.T) {
	t.Run("Ensure Setup works Properly", TestInit)

	t.Run("Test Version Endpoint", TestVersion)
	t.Run("Test IsActive Endpoint", TestIsActive)

	t.Run("Test Play Endpoint", TestPlay)
	t.Run("Test Switch Endpoint", TestSwitch)
	t.Run("Test GoToWebsite Endpoint", TestGoToWebsite)
	t.Run("Test Uninstall Endpoint", TestUninstall)
	t.Run("Test Exit Endpoint", TestExit)
}

// HELPERS

func expectFmt(message string, args ...interface{}) []interface{} {
	return append([]interface{}{message}, args...)
}
