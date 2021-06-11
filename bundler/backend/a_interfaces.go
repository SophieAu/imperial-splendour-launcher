package backend

import "github.com/wailsapp/wails/lib/logger"

type Logger interface {
	Info(message string)
	Infof(message string, args ...interface{})
	Warn(message string)
	Warnf(message string, args ...interface{})
	Debug(message string)
	Debugf(message string, args ...interface{})

	InfoFields(message string, fields logger.Fields)
	DebugFields(message string, fields logger.Fields)
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

type Browser interface {
	OpenURL(url string) error
}

type Window interface {
	Close()
}
