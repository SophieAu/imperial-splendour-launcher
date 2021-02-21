package backend_test

import (
	"imperial-splendour-launcher/backend"
	"imperial-splendour-launcher/backend/mock"

	testifyMock "github.com/stretchr/testify/mock"

	"testing"
)

func before() (*backend.API, *mock.Browser, *mock.Window, *mock.Logger, *mock.SystemHandler) {
	mockS := &mock.SystemHandler{}
	mockB := &mock.Browser{}
	mockW := &mock.Window{}
	mockL := &mock.Logger{}

	mockS.On("Executable").Return(".", nil)
	mockL.On("Infof", testifyMock.Anything, testifyMock.Anything).Return()

	api := &backend.API{}
	err := api.Init(mockB, mockW, mockL, mockS)
	if err != nil {
		panic(err)
	}

	return api, mockB, mockW, mockL, mockS
}

func after(api backend.API) {
}

func testPlay(t *testing.T) {
	// api, browser, window, logger, systemHandler := before()
	// _ = browser
	// _ = window
	// _ = logger
	// _ = systemHandler

	// api.Play()
}
