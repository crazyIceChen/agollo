package agollo

import (
	"github.com/sirupsen/logrus"
)

var logger LoggerInterface

func init() {
	initLogger(initLogrusLog())
}

func initLogger(ILogger LoggerInterface) {
	logger = ILogger
}

type LoggerInterface interface {
	Debugf(format string, params ...interface{})

	Infof(format string, params ...interface{})

	Warnf(format string, params ...interface{})

	Errorf(format string, params ...interface{})

	Debug(v ...interface{})

	Info(v ...interface{})

	Warn(v ...interface{})

	Error(v ...interface{})
}

func initLogrusLog() LoggerInterface {
	logrus.New()
	return logger
}
