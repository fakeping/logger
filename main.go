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
