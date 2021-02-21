package mock

import (
	"github.com/stretchr/testify/mock"
	"github.com/wailsapp/wails/lib/logger"
)

// System Handler
type SystemHandler struct {
	mock.Mock
}

func (w *SystemHandler) WriteFile(filePath string, data []byte) error {
	args := w.Called(filePath, data)
	return args.Error(0)
}

func (w *SystemHandler) ReadFile(filePath string) ([]byte, error) {
	args := w.Called(filePath)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]byte), args.Error(1)
}

func (w *SystemHandler) MoveFile(source, destination string) error {
	args := w.Called(source, destination)
	return args.Error(0)
}

func (w *SystemHandler) Executable() (string, error) {
	args := w.Called()
	return args.Get(0).(string), args.Error(1)
}

func (w *SystemHandler) Getenv(key string) string {
	args := w.Called(key)
	return args.Get(0).(string)
}

// Logger
type Logger struct {
	mock.Mock
}

func (l *Logger) Infof(message string, args ...interface{}) {
	x := append([]interface{}{message}, args...)
	_ = l.Called(x)
}
func (l *Logger) Debug(message string) {
	_ = l.Called(message)
}
func (l *Logger) Info(message string) {
	_ = l.Called(message)
}
func (l *Logger) InfoFields(message string, fields logger.Fields) {
	_ = l.Called(message, fields)
}
func (l *Logger) Debugf(message string, args ...interface{}) {
	_ = l.Called(message, args)
}
func (l *Logger) DebugFields(message string, fields logger.Fields) {
	_ = l.Called(message, fields)
}
func (l *Logger) Warn(message string) {
	_ = l.Called(message)
}
func (l *Logger) Warnf(message string, args ...interface{}) {
	_ = l.Called(message, args)
}
func (l *Logger) WarnFields(message string, fields logger.Fields) {
	_ = l.Called(message, fields)
}
func (l *Logger) Error(message string) {
	_ = l.Called(message)
}
func (l *Logger) Errorf(message string, args ...interface{}) {
	_ = l.Called(message, args)
}
func (l *Logger) ErrorFields(message string, fields logger.Fields) {
	_ = l.Called(message, fields)
}
func (l *Logger) Fatal(message string) {
	_ = l.Called(message)
}
func (l *Logger) Fatalf(message string, args ...interface{}) {
	_ = l.Called(message, args)
}
func (l *Logger) FatalFields(message string, fields logger.Fields) {
	_ = l.Called(message, fields)
}
func (l *Logger) Panic(message string) {
	_ = l.Called(message)
}
func (l *Logger) Panicf(message string, args ...interface{}) {
	_ = l.Called(message, args)
}
func (l *Logger) PanicFields(message string, fields logger.Fields) {
	_ = l.Called(message, fields)
}

// Browser
type Browser struct {
	mock.Mock
}

func (b *Browser) OpenURL(url string) error {
	args := b.Called(url)
	return args.Error(0)
}

// Window
type Window struct {
	mock.Mock
}

func (w *Window) Close() {
	_ = w.Called()
}
