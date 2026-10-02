package logger

import (
	"fmt"
	"log"
	"time"
)

var EnableDebug bool

func init() {
	log.SetFlags(0)
}

func logf(level string, format string, args ...any) {
	log.Printf(
		"%s [%-5s] %s",
		time.Now().Format("2006-01-02 15:04:05.000000"),
		level,
		fmt.Sprintf(format, args...),
	)
}

func Debugf(format string, args ...any) {
	if EnableDebug {
		logf("[DEBUG]", format, args...)
	}
}

func Infof(format string, args ...any) {
	logf("[INFO ]", format, args...)
}

func Warnf(format string, args ...any) {
	logf("[WARN ]", format, args...)
}

func Errorf(format string, args ...any) {
	logf("[ERROR]", format, args...)
}

func Fatalf(format string, args ...any) {
	logf("[FATAL]", format, args...)
	panic(fmt.Sprintf(format, args...))
}
