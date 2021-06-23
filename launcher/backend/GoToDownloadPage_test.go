package backend_test

import (
	"errors"
	"imperial-splendour-launcher/backend/test"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testing"
)

func TestGoToDownloadPage(t *testing.T) {
	t.Run("Error opening website", func(t *testing.T) {
		api, browser, _, _, _ := test.Before()

		// error in opening URL
		browser.On("OpenURL", mock.Anything).Return(errors.New("error")).Once()
		err := api.GoToDownloadPage()

		assert.NotNil(t, err)
		browser.AssertCalled(t, "OpenURL", "https://imperialsplendour.com/download")

		test.After(*api)
	})

	t.Run("Successfully open website", func(t *testing.T) {
		api, browser, _, _, _ := test.Before()

		browser.On("OpenURL", mock.Anything).Return(nil).Once()
		err := api.GoToDownloadPage()

		assert.Nil(t, err)
		browser.AssertCalled(t, "OpenURL", "https://imperialsplendour.com/download")

		test.After(*api)
	})
}
