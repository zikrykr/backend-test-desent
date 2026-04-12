package infrastructure

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	logger *logrus.Logger
}

func NewLogger() *Logger {
	return &Logger{
		logger: logrus.New(),
	}
}

func (l *Logger) Info(message string) {
	l.logger.Info(message)
}

func (l *Logger) Error(message string, err error) {
	l.logger.WithFields(logrus.Fields{
		"error": err,
	}).Error(message)
}
