package backend

import "github.com/wailsapp/wails/lib/logger"

// Logger .
type Logger interface {
	Infof(message string, args ...interface{})
	Debug(message string)
	Info(message string)
	InfoFields(message string, fields logger.Fields)
	Debugf(message string, args ...interface{})
	DebugFields(message string, fields logger.Fields)
	Warn(message string)
	Warnf(message string, args ...interface{})
	WarnFields(message string, fields logger.Fields)
	Error(message string)
	Errorf(message string, args ...interface{})
	ErrorFields(message string, fields logger.Fields)
	Fatal(message string)
	Fatalf(message string, args ...interface{})
	FatalFields(message string, fields logger.Fields)
	Panic(message string)
	Panicf(message string, args ...interface{})
	PanicFields(message string, fields logger.Fields)
}

// Browser .
type Browser interface {
	OpenURL(url string) error
}

// Window .
type Window interface {
	Close()
}
