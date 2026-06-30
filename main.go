package logger

import (
	"github.com/charmbracelet/log"
	"time"
	"os"
)

func NewLogger (prefix string) (*log.Logger) {

	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller: true,
		ReportTimestamp: true,
		TimeFormat: time.RFC3339,
		Prefix: prefix,
	})

	return logger

}

func SetLogLevel (level string) (log.Level) {

	switch level {
		case "Debug":
			return log.DebugLevel
		case "Warn":
			return log.WarnLevel
		case "Fatal":
			return log.FatalLevel
		case "Error":
			return log.ErrorLevel
	}

	return log.InfoLevel

}

