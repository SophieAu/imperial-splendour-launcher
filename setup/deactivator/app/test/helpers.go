package test

import (
	"deactivator/app"
	"strconv"
)

func ExpectFmt(message string, args ...interface{}) []interface{} {
	return append([]interface{}{message}, args...)
}

func FmtInfoFile(isActive bool, version, usChecksum string) []byte {
	return []byte("{\n\t\"isActive\": " + strconv.FormatBool(isActive) + ",\n\t\"version\": \"" + version + "\",\n\t\"usChecksum\": \"" + usChecksum + "\"\n}")
}

func Before() (*app.API, *MockSystemHandler) {
	return VariableBefore("2.0", true, "test")
}

func VariableBefore(version string, isActive bool, usChecksum string) (*app.API, *MockSystemHandler) {
	mockS := &MockSystemHandler{}

	mockS.On("Executable").Return(".", nil).Once()
	mockS.On("Getenv", "APPDATA").Return("APPDATA").Once()
	mockS.On("ReadFile", "./IS_Files/IS_Info.json").Return(FmtInfoFile(isActive, version, usChecksum), nil)

	api := &app.API{}

	if err := api.Init(mockS); err != nil {
		panic(err)
	}

	return api, mockS
}

func After(api app.API) {
}
