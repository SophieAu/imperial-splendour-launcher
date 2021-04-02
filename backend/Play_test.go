package backend_test

import (
	"errors"
	"imperial-splendour-launcher/backend/testHelpers"

	"github.com/stretchr/testify/assert"
	testifyMock "github.com/stretchr/testify/mock"

	"testing"
)

func TestPlay(t *testing.T) {
	t.Run("Error launching game", func(t *testing.T) {
		api, browser, window, _, _ := testHelpers.Before()

		// error in opening URL
		browser.On("OpenURL", testifyMock.Anything).Return(errors.New("error")).Once()
		err := api.Play()

		assert.NotNil(t, err)
		browser.AssertCalled(t, "OpenURL", "steam://rungameid/10500")
		window.AssertNotCalled(t, "Close")

		testHelpers.After(*api)
	})

	t.Run("Successfully launch game", func(t *testing.T) {
		api, browser, window, _, _ := testHelpers.Before()

		browser.On("OpenURL", testifyMock.Anything).Return(nil).Once()
		err := api.Play()

		assert.Nil(t, err)
		browser.AssertCalled(t, "OpenURL", "steam://rungameid/10500")
		window.AssertCalled(t, "Close")

		testHelpers.After(*api)
	})
}
