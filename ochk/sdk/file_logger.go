package sdk

import (
	"fmt"
	"os"
)

type FileLogger struct {
	filePath string
	file     *os.File
}

func NewFileLogger(filePath string) *FileLogger {
	return &FileLogger{filePath: filePath}
}

func (fl *FileLogger) Init() error {
	file, err := os.OpenFile(fl.filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		return fmt.Errorf("error opening log file %s: %v", fl.filePath, err)
	}

	fl.file = file

	return nil
}

func (fl *FileLogger) Close() error {
	if fl.file == nil {
		return nil
	}

	if err := fl.Close(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error closing log file, ignoring: %v", err)
	}

	return nil
}

func (fl *FileLogger) Printf(format string, args ...interface{}) {
	if fl.file == nil {
		panic("file logger not initialized")
	}

	if len(format) == 0 || format[len(format)-1] != '\n' {
		format += "\n"
	}

	if _, err := fmt.Fprintf(fl.file, format, args...); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error writing to log file %s: %v", fl.filePath, err)
	}
}

func (fl *FileLogger) Debugf(format string, args ...interface{}) {
	fl.Printf(format, args...)
}
