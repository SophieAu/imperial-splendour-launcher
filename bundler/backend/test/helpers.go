package test

import (
	"imperial-splendour-bundler/backend"
	"imperial-splendour-bundler/backend/mocks"

	"github.com/stretchr/testify/mock"
)

func Before() (*backend.API, *mocks.MockBrowser, *mocks.MockWindow, *mocks.MockLogger, *mocks.MockStore, *mocks.MockSystemHandler) {
	mockSh := &mocks.MockSystemHandler{}
	mockB := &mocks.MockBrowser{}
	mockW := &mocks.MockWindow{}
	mockL := &mocks.MockLogger{}
	mockD := &mocks.MockDialog{}
	mockSt := &mocks.MockStore{}

	mockW.On("Close").Return()

	mockL.On("Info", mock.Anything).Return()
	mockL.On("Infof", mock.Anything, mock.Anything).Return()
	mockL.On("Warn", mock.Anything).Return()
	mockL.On("Warnf", mock.Anything, mock.Anything).Return()
	mockL.On("Debug", mock.Anything).Return(nil)
	mockL.On("Debugf", mock.Anything, mock.Anything).Return()

	mockSt.On("Update", mock.Anything).Return()

	api := &backend.API{}

	mockSh.On("StartCommand", "iscc /?", mock.Anything).Return(nil)
	if err := api.Init(mockB, mockW, mockL, mockSt, mockD, mockSh); err != nil {
		panic(err)
	}
	return api, mockB, mockW, mockL, mockSt, mockSh
}

func After(api backend.API) {
}
