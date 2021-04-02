package backend_test

import (
	"imperial-splendour-launcher/backend"
	"imperial-splendour-launcher/backend/mock"
	"strconv"

	testifyMock "github.com/stretchr/testify/mock"
)

func expectFmt(message string, args ...interface{}) []interface{} {
	return append([]interface{}{message}, args...)
}

func fmtInfoFile(isActive bool, version, usChecksum string) []byte {
	return []byte("{\n\t\"isActive\": " + strconv.FormatBool(isActive) + ",\n\t\"version\": \"" + version + "\",\n\t\"usChecksum\": \"" + usChecksum + "\"\n}")
}

func before() (*backend.API, *mock.Browser, *mock.Window, *mock.Logger, *mock.SystemHandler) {
	return variableBefore("2.0", true, "test")
}

func variableBefore(version string, isActive bool, usChecksum string) (*backend.API, *mock.Browser, *mock.Window, *mock.Logger, *mock.SystemHandler) {
	mockS := &mock.SystemHandler{}
	mockB := &mock.Browser{}
	mockW := &mock.Window{}
	mockL := &mock.Logger{}

	mockW.On("Close").Return()

	mockL.On("Infof", testifyMock.Anything, testifyMock.Anything).Return()
	mockL.On("Info", testifyMock.Anything).Return()
	mockL.On("Warnf", testifyMock.Anything, testifyMock.Anything).Return()
	mockL.On("Debugf", testifyMock.Anything, testifyMock.Anything).Return()
	mockL.On("Debug", testifyMock.Anything).Return(nil)

	mockS.On("Executable").Return(".", nil)
	mockS.On("Getenv", "APPDATA").Return("APPDATA")
	mockS.On("ReadFile", "./IS_Files/IS_info.json").Return(fmtInfoFile(isActive, version, usChecksum), nil)

	api := &backend.API{}

	if err := api.Init(mockB, mockW, mockL, mockS); err != nil {
		panic(err)
	}

	return api, mockB, mockW, mockL, mockS
}

func after(api backend.API) {
}
