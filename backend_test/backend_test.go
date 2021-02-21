package backend_test

import (
	"testing"
)

func TestAll(t *testing.T) {
	t.Run("Ensure Setup and Teardown Work Properly", TestInit)
	t.Run("Ensure adding projects works", testPlay)
}

// HELPERS

func expectFmt(message string, args ...interface{}) []interface{} {
	return append([]interface{}{message}, args...)

}
