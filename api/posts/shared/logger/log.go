package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

const (
	appName    = "posts"
	appKey     = "app"
	timeFormat = "2006-01-02 15:04:05"
)

func NewPostLogger() *logrus.Entry {
	logger := &logrus.Logger{
		Out: os.Stdout,
		Formatter: &logrus.JSONFormatter{
			TimestampFormat: timeFormat,
		},
		Level: logrus.InfoLevel,
	}

	return logger.WithField(appKey, appName)
}
