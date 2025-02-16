package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Entry

func Init(path string) error {
	logFile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	log := logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	log.Out = logFile
	logger = logrus.NewEntry(log)
	return nil
}

func GetLogger() *logrus.Entry {
	return logger
}
