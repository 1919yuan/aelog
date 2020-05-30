package aelog

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strings"

	"cloud.google.com/go/logging"
	"google.golang.org/appengine"
)

type StackDriverLogger struct {
	client *logging.Client
	logger *logging.Logger
}

var _ Logger = &StackDriverLogger{}

func NewStackDriverLogger(
	ctx context.Context, client *logging.Client) Logger {
	c := client
	project_id := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if c == nil {
		var err error = nil
		c, err = logging.NewClient(ctx, project_id)
		if err != nil {
			log.Fatalf("Failed to connect to stackdriver: %v", err)
		}
	}
	logger := c.Logger(project_id)
	return &StackDriverLogger{
		client: c,
		logger: logger,
	}
}

func (l *StackDriverLogger) Close() {
	l.client.Close()
}

func UseSyslog() bool {
	return appengine.IsDevAppServer() || strings.HasSuffix(os.Args[0], ".test")
}

func (l *StackDriverLogger) Fatal(format string, args ...interface{}) {
	if UseSyslog() {
		_, file, no, ok := runtime.Caller(2)
		if ok {
			log.Printf(
				fmt.Sprintf("[FATAL %s:%d] ", path.Base(file), no)+
					format, args...)
		} else {
			log.Printf(format, args...)
		}
	} else {
		l.logger.Log(
			logging.Entry{
				Payload:  fmt.Sprintf(format, args...),
				Severity: logging.Emergency,
			})
	}
	log.Fatalln("Exit: encoutered Log Fatal.")
}

func (l *StackDriverLogger) Error(format string, args ...interface{}) {
	if UseSyslog() {
		_, file, no, ok := runtime.Caller(2)
		if ok {
			log.Printf(
				fmt.Sprintf("[ERROR %s:%d] ", path.Base(file), no)+
					format, args...)
		} else {
			log.Printf(format, args...)
		}
	} else {
		l.logger.Log(
			logging.Entry{
				Payload:  fmt.Sprintf(format, args...),
				Severity: logging.Error,
			})
	}
}

func (l *StackDriverLogger) Warning(format string, args ...interface{}) {
	if UseSyslog() {
		_, file, no, ok := runtime.Caller(2)
		if ok {
			log.Printf(
				fmt.Sprintf("[WARNING %s:%d] ", path.Base(file), no)+
					format, args...)
		} else {
			log.Printf(format, args...)
		}
	} else {
		l.logger.Log(
			logging.Entry{
				Payload:  fmt.Sprintf(format, args...),
				Severity: logging.Warning,
			})
	}
}

func (l *StackDriverLogger) Info(format string, args ...interface{}) {
	if UseSyslog() {
		_, file, no, ok := runtime.Caller(2)
		if ok {
			log.Printf(
				fmt.Sprintf("[INFO %s:%d] ", path.Base(file), no)+
					format, args...)
		} else {
			log.Printf(format, args...)
		}
	} else {
		l.logger.Log(
			logging.Entry{
				Payload:  fmt.Sprintf(format, args...),
				Severity: logging.Info,
			})
	}
}

func (l *StackDriverLogger) Debug(format string, args ...interface{}) {
	if UseSyslog() {
		_, file, no, ok := runtime.Caller(2)
		if ok {
			log.Printf(
				fmt.Sprintf("[DEBUG %s:%d] ", path.Base(file), no)+
					format, args...)
		} else {
			log.Printf(format, args...)
		}
	} else {
		l.logger.Log(
			logging.Entry{
				Payload:  fmt.Sprintf(format, args...),
				Severity: logging.Debug,
			})
	}
}
