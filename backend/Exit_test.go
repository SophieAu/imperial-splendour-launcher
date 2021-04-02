package backend_test

import (
	"errors"
	"imperial-splendour-launcher/backend/testHelpers"

	"github.com/stretchr/testify/assert"
	testifyMock "github.com/stretchr/testify/mock"

	"testing"
)

func TestExit(t *testing.T) {
	api, browser, window, _, _ := testHelpers.Before()

	// error in opening URL
	browser.On("OpenURL", testifyMock.Anything).Return(errors.New("error")).Once()
	err := api.Play()

	assert.NotNil(t, err)
	browser.AssertCalled(t, "OpenURL", "steam://rungameid/10500")
	window.AssertNotCalled(t, "Close")

	// working properly
	browser.On("OpenURL", testifyMock.Anything).Return(nil).Once()
	err = api.Play()

	assert.Nil(t, err)
	browser.AssertCalled(t, "OpenURL", "steam://rungameid/10500")
	window.AssertCalled(t, "Close")

	testHelpers.After(*api)
}
