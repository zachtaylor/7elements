package log

import (
	"github.com/Sirupsen/logrus"
)

type Log map[string]interface{}

func (log Log) Clone() Log {
	clone := new()
	for key, val := range log {
		clone.Add(key, val)
	}
	return clone
}

func (log Log) Add(name string, value interface{}) Log {
	log[name] = value
	return log
}

func (log Log) Debug(message string) {
	logger.WithFields(logrus.Fields(log)).Debug(message)
}

func (log Log) Info(message string) {
	logger.WithFields(logrus.Fields(log)).Info(message)
}

func (log Log) Warn(message string) {
	logger.WithFields(logrus.Fields(log)).Warn(message)
}

func (log Log) Error(message interface{}) {
	logger.WithFields(logrus.Fields(log)).Error(message)
}
