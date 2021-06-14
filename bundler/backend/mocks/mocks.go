package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/wailsapp/wails/lib/logger"
)

// System Handler
type MockSystemHandler struct {
	mock.Mock
}

func (sh *MockSystemHandler) MkdirAll(filePath string) error {
	args := sh.Called(filePath)
	return args.Error(0)
}

func (sh *MockSystemHandler) WriteFile(filePath string, data []byte) error {
	args := sh.Called(filePath, data)
	return args.Error(0)
}

func (sh *MockSystemHandler) ReadFile(filePath string) ([]byte, error) {
	args := sh.Called(filePath)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]byte), args.Error(1)
}

func (sh *MockSystemHandler) MoveFile(source, destination string) error {
	args := sh.Called(source, destination)
	return args.Error(0)
}

func (sh *MockSystemHandler) RunCommand(name string, arg ...string) error {
	args := sh.Called(name, arg)
	return args.Error(0)
}

func (sh *MockSystemHandler) DoesFileExist(path string) (bool, error) {
	args := sh.Called(path)
	return args.Get(0).(bool), args.Error(1)
}

func (sh *MockSystemHandler) GetDirContentByName(dirname string) ([]string, error) {
	args := sh.Called(dirname)
	return args.Get(0).([]string), args.Error(1)
}

// Logger
type MockLogger struct {
	mock.Mock
}

func (l *MockLogger) Infof(message string, args ...interface{}) {
	x := append([]interface{}{message}, args...)
	_ = l.Called(x)
}
func (l *MockLogger) Debug(message string) {
	_ = l.Called(message)
}
func (l *MockLogger) Info(message string) {
	_ = l.Called(message)
}
func (l *MockLogger) InfoFields(message string, fields logger.Fields) {
	_ = l.Called(message, fields)
}
func (l *MockLogger) Debugf(message string, args ...interface{}) {
	_ = l.Called(message, args)
}
func (l *MockLogger) DebugFields(message string, fields logger.Fields) {
	_ = l.Called(message, fields)
}
func (l *MockLogger) Warn(message string) {
	_ = l.Called(message)
}
func (l *MockLogger) Warnf(message string, args ...interface{}) {
	_ = l.Called(message, args)
}
func (l *MockLogger) WarnFields(message string, fields logger.Fields) {
	_ = l.Called(message, fields)
}
func (l *MockLogger) Error(message string) {
	_ = l.Called(message)
}
func (l *MockLogger) Errorf(message string, args ...interface{}) {
	_ = l.Called(message, args)
}
func (l *MockLogger) ErrorFields(message string, fields logger.Fields) {
	_ = l.Called(message, fields)
}
func (l *MockLogger) Fatal(message string) {
	_ = l.Called(message)
}
func (l *MockLogger) Fatalf(message string, args ...interface{}) {
	_ = l.Called(message, args)
}
func (l *MockLogger) FatalFields(message string, fields logger.Fields) {
	_ = l.Called(message, fields)
}
func (l *MockLogger) Panic(message string) {
	_ = l.Called(message)
}
func (l *MockLogger) Panicf(message string, args ...interface{}) {
	_ = l.Called(message, args)
}
func (l *MockLogger) PanicFields(message string, fields logger.Fields) {
	_ = l.Called(message, fields)
}

// Browser
type MockBrowser struct {
	mock.Mock
}

func (b *MockBrowser) OpenURL(url string) error {
	args := b.Called(url)
	return args.Error(0)
}

// Window
type MockWindow struct {
	mock.Mock
}

func (w *MockWindow) Close() {
	_ = w.Called()
}

// Store
type MockStore struct {
	mock.Mock
}

func (s *MockStore) OnError(callback func(error)) {
	_ = s.Called(callback)
}
func (s *MockStore) Set(data interface{}) error {
	args := s.Called(data)
	return args.Error(0)
}
func (s *MockStore) Update(updater interface{}) {
	_ = s.Called(updater)
}
func (s *MockStore) Subscribe(callback interface{}) {
	_ = s.Called(callback)
}
func (s *MockStore) Get() interface{} {
	args := s.Called()
	return args.Error(0)
}
