package test

import (
	"imperial-splendour-bundler/backend"

	"github.com/stretchr/testify/mock"
)

func Before() (*backend.API, *MockBrowser, *MockWindow, *MockLogger, *MockSystemHandler) {
	mockS := &MockSystemHandler{}
	mockB := &MockBrowser{}
	mockW := &MockWindow{}
	mockL := &MockLogger{}

	mockW.On("Close").Return()

	mockL.On("Infof", mock.Anything, mock.Anything).Return()
	mockL.On("Info", mock.Anything).Return()
	mockL.On("Warnf", mock.Anything, mock.Anything).Return()
	mockL.On("Debugf", mock.Anything, mock.Anything).Return()
	mockL.On("Debug", mock.Anything).Return(nil)

	api := &backend.API{}

	if err := api.Init(mockB, mockW, mockL, mockS); err != nil {
		panic(err)
	}
	return api, mockB, mockW, mockL, mockS
}

func After(api backend.API) {
}
