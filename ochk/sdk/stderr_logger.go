package sdk

import (
	"log"
	"os"
)

type StdErrLogger struct {
	logger *log.Logger
}

func NewStdErrLogger() *StdErrLogger {
	return &StdErrLogger{logger: log.New(os.Stderr, "[DEBUG]", log.LstdFlags)}
}

func (fl *StdErrLogger) Printf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (fl *StdErrLogger) Debugf(format string, args ...interface{}) {
	fl.Printf(format, args...)
}
