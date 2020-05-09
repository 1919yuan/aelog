package aelog

import (
	"context"
	"os"
)

type Logger interface {
	Fatal(format string, args ...interface{})
	Error(format string, args ...interface{})
	Warning(format string, args ...interface{})
	Info(format string, args ...interface{})
	Debug(format string, args ...interface{})
	Close()
}

var (
	Log Logger
)

func init() {
	if os.Getenv("GOOGLE_CLOUD_PROJECT") != "" {
		Log = NewStackDriverLogger(context.Background(), nil)
	} else {
		Log = NewSysLogger()
	}
}

func Debug(format string, args ...interface{}) {
	Log.Debug(format, args...)
}

func Error(format string, args ...interface{}) {
	Log.Error(format, args...)
}

func Fatal(format string, args ...interface{}) {
	Log.Fatal(format, args...)
}

func Info(format string, args ...interface{}) {
	Log.Info(format, args...)
}

func Warning(format string, args ...interface{}) {
	Log.Warning(format, args...)
}
