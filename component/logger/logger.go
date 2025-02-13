package logger

import "github.com/sirupsen/logrus"

func Init() {
	log := logrus.New()
	logrus.NewEntry(log)
}
