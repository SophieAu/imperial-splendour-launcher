package backend_test

import (
	"errors"

	"github.com/stretchr/testify/assert"
	testifyMock "github.com/stretchr/testify/mock"

	"testing"
)

func TestGoToWebsite(t *testing.T) {
	api, browser, window, _, _ := before()

	// error in opening URL
	browser.On("OpenURL", testifyMock.Anything).Return(errors.New("error")).Once()
	err := api.GoToWebsite()

	assert.NotNil(t, err)
	browser.AssertCalled(t, "OpenURL", "https://imperialsplendour.com/")
	window.AssertNotCalled(t, "Close")

	// working properly
	browser.On("OpenURL", testifyMock.Anything).Return(nil).Once()
	err = api.Play()

	assert.Nil(t, err)
	browser.AssertCalled(t, "OpenURL", "https://imperialsplendour.com/")
	window.AssertCalled(t, "Close")

	after(*api)
}
