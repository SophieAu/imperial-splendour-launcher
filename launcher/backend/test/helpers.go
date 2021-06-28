package test

import (
	"imperial-splendour-launcher/backend"
	"imperial-splendour-launcher/backend/mocks"
	"strconv"

	"github.com/stretchr/testify/mock"
)

func ExpectFmt(message string, args ...interface{}) []interface{} {
	return append([]interface{}{message}, args...)
}

func FmtInfoFile(isActive bool, version, usChecksum string) []byte {
	return []byte("{\n\t\"isActive\": " + strconv.FormatBool(isActive) + ",\n\t\"version\": \"" + version + "\",\n\t\"usChecksum\": \"" + usChecksum + "\"\n}")
}

func Before() (*backend.API, *mocks.MockBrowser, *mocks.MockWindow, *mocks.MockLogger, *mocks.MockSystemHandler) {
	return VariableBefore("2.0", true, "test")
}

func VariableBefore(version string, isActive bool, usChecksum string) (*backend.API, *mocks.MockBrowser, *mocks.MockWindow, *mocks.MockLogger, *mocks.MockSystemHandler) {
	mockS := &mocks.MockSystemHandler{}
	mockB := &mocks.MockBrowser{}
	mockW := &mocks.MockWindow{}
	mockL := &mocks.MockLogger{}

	mockW.On("Close").Return()

	mockL.On("Infof", mock.Anything, mock.Anything).Return()
	mockL.On("Info", mock.Anything).Return()
	mockL.On("Warnf", mock.Anything, mock.Anything).Return()
	mockL.On("Debugf", mock.Anything, mock.Anything).Return()
	mockL.On("Debug", mock.Anything).Return(nil)

	mockS.On("Executable").Return(".", nil).Once()
	mockS.On("Getenv", "APPDATA").Return("APPDATA").Once()
	mockS.On("ReadFile", "./IS_Files/IS_Info.json").Return(FmtInfoFile(isActive, version, usChecksum), nil)

	api := &backend.API{}

	if err := api.Init(mockB, mockW, mockL, mockS); err != nil {
		panic(err)
	}

	return api, mockB, mockW, mockL, mockS
}

func After(api backend.API) {
}
