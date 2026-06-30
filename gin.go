package logger

import (
	"strings"
	"time"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func GinLogger(l *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		l.Info(
			"request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"latency", time.Since(start),
			"ip", c.ClientIP(),
		)
	}
}

type LogWriter struct {
	logger *log.Logger
}

func NewLogWriter(l *log.Logger) *LogWriter {
	return &LogWriter{
		logger: l,
	}
}

func (w *LogWriter) Write(p []byte) (n int, err error) {
	w.logger.Debug(strings.TrimSpace(string(p)))
	return len(p), nil
}
