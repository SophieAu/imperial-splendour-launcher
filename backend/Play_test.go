package backend_test

import (
	"errors"
	"imperial-splendour-launcher/backend/test"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testing"
)

func TestPlay(t *testing.T) {
	t.Run("Error launching game", func(t *testing.T) {
		api, browser, window, _, _ := test.Before()

		// error in opening URL
		browser.On("OpenURL", mock.Anything).Return(errors.New("error")).Once()
		err := api.Play()

		assert.NotNil(t, err)
		browser.AssertCalled(t, "OpenURL", "steam://rungameid/10500")
		window.AssertNotCalled(t, "Close")

		test.After(*api)
	})

	t.Run("Successfully launch game", func(t *testing.T) {
		api, browser, window, _, _ := test.Before()

		browser.On("OpenURL", mock.Anything).Return(nil).Once()
		err := api.Play()

		assert.Nil(t, err)
		browser.AssertCalled(t, "OpenURL", "steam://rungameid/10500")
		window.AssertCalled(t, "Close")

		test.After(*api)
	})
}
