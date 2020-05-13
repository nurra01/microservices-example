package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// NewLogger creates a new logger instance
func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "02-01-2006 15:04:05",
	})
	return logger
}
