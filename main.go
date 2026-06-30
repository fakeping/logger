package logger

import (
	"github.com/charmbracelet/log"
	"time"
	"os"
	"strings"
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

	switch strings.ToLower(level) {
		case "debug":
			return log.DebugLevel
		case "warn":
			return log.WarnLevel
		case "fatal":
			return log.FatalLevel
		case "error":
			return log.ErrorLevel
	}

	return log.InfoLevel

}

