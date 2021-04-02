package backend_test

import (
	"errors"
	"imperial-splendour-launcher/backend/testHelpers"

	"github.com/stretchr/testify/assert"
	testifyMock "github.com/stretchr/testify/mock"

	"testing"
)

func TestGoToWebsite(t *testing.T) {
	t.Run("Error opening website", func(t *testing.T) {
		api, browser, _, _, _ := testHelpers.Before()

		// error in opening URL
		browser.On("OpenURL", testifyMock.Anything).Return(errors.New("error")).Once()
		err := api.GoToWebsite()

		assert.NotNil(t, err)
		browser.AssertCalled(t, "OpenURL", "https://imperialsplendour.com/")

		testHelpers.After(*api)
	})

	t.Run("Successfully open website", func(t *testing.T) {
		api, browser, _, _, _ := testHelpers.Before()

		browser.On("OpenURL", testifyMock.Anything).Return(nil).Once()
		err := api.GoToWebsite()

		assert.Nil(t, err)
		browser.AssertCalled(t, "OpenURL", "https://imperialsplendour.com/")

		testHelpers.After(*api)
	})
}
