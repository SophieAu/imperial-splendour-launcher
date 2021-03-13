package backend_test

import (
	"errors"
	"imperial-splendour-launcher/backend"
	"imperial-splendour-launcher/backend/mock"
	"strconv"

	"github.com/stretchr/testify/assert"
	testifyMock "github.com/stretchr/testify/mock"

	"testing"
)

func before() (*backend.API, *mock.Browser, *mock.Window, *mock.Logger, *mock.SystemHandler) {
	return variableBefore("2.0", true, "test")
}

func variableBefore(version string, isActive bool, usChecksum string) (*backend.API, *mock.Browser, *mock.Window, *mock.Logger, *mock.SystemHandler) {
	mockS := &mock.SystemHandler{}
	mockB := &mock.Browser{}
	mockW := &mock.Window{}
	mockL := &mock.Logger{}

	mockS.On("Executable").Return(".", nil)
	mockW.On("Close").Return()
	mockL.On("Infof", testifyMock.Anything, testifyMock.Anything).Return()
	mockL.On("Errorf", testifyMock.Anything, testifyMock.Anything).Return()
	mockS.On("ReadFile", "IS_Files/IS_info.json").Return([]byte("{\"isActive\": "+strconv.FormatBool(isActive)+", \"version\": \""+version+"\", \"usChecksum\": \""+usChecksum+"\"}"), nil)

	api := &backend.API{}
	err := api.Init(mockB, mockW, mockL, mockS)
	if err != nil {
		panic(err)
	}

	return api, mockB, mockW, mockL, mockS
}

func after(api backend.API) {
}

func TestVersion(t *testing.T) {
	api, _, _, _, _ := variableBefore("2.1", false, "test")

	version := api.Version()
	assert.Equal(t, "2.1", version)

	after(*api)
}

func TestIsActive(t *testing.T) {
	api, _, _, _, _ := variableBefore("2.1", false, "test")

	isActive := api.IsActive()
	assert.Equal(t, false, isActive)

	api, _, _, _, _ = variableBefore("2.1", true, "test")

	isActive = api.IsActive()
	assert.Equal(t, true, isActive)

	after(*api)
}

func TestPlay(t *testing.T) {
	api, browser, window, _, _ := before()

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

	after(*api)
}

func TestSwitch(t *testing.T) {
	assert.True(t, false)

}

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

func TestUninstall(t *testing.T) {
	assert.True(t, false)

}

func TestExit(t *testing.T) {
	api, browser, window, _, _ := before()

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

	after(*api)
}
