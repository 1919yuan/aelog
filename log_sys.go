package aelog

import (
	"fmt"
	"log"
	"path"
	"runtime"
)

type SysLogger struct {
}

var _ Logger = &SysLogger{}

func NewSysLogger() Logger {
	return &SysLogger{}
}

func (l *SysLogger) Close() {}

func (l *SysLogger) Fatal(format string, args ...interface{}) {
	_, file, no, ok := runtime.Caller(1)
	if ok {
		log.Printf(
			fmt.Sprintf("[FATAL %s:%d] ", path.Base(file), no)+
				format, args...)
	} else {
		log.Printf(format, args...)
	}
	log.Fatalln("Exit: encoutered Log Fatal.")
}

func (l *SysLogger) Error(format string, args ...interface{}) {
	_, file, no, ok := runtime.Caller(1)
	if ok {
		log.Printf(
			fmt.Sprintf("[ERROR %s:%d] ", path.Base(file), no)+
				format, args...)
	} else {
		log.Printf(format, args...)
	}
}

func (l *SysLogger) Warning(format string, args ...interface{}) {
	_, file, no, ok := runtime.Caller(1)
	if ok {
		log.Printf(
			fmt.Sprintf("[WARNING %s:%d] ", path.Base(file), no)+
				format, args...)
	} else {
		log.Printf(format, args...)
	}
}

func (l *SysLogger) Info(format string, args ...interface{}) {
	_, file, no, ok := runtime.Caller(1)
	if ok {
		log.Printf(
			fmt.Sprintf("[INFO %s:%d] ", path.Base(file), no)+
				format, args...)
	} else {
		log.Printf(format, args...)
	}
}

func (l *SysLogger) Debug(format string, args ...interface{}) {
	_, file, no, ok := runtime.Caller(1)
	if ok {
		log.Printf(
			fmt.Sprintf("[DEBUG %s:%d] ", path.Base(file), no)+
				format, args...)
	} else {
		log.Printf(format, args...)
	}
}
