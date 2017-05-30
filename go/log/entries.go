package log

import (
	"github.com/Sirupsen/logrus"
)

type Entries map[string]interface{}

func (entries *Entries) Add(name string, value interface{}) *Entries {
	(*entries)[name] = value
	return entries
}

func (entries *Entries) Debug(message string) {
	log.WithFields(logrus.Fields(*entries)).Debug(message)
	entries.clear()
}

func (entries *Entries) Info(message string) {
	log.WithFields(logrus.Fields(*entries)).Info(message)
	entries.clear()
}

func (entries *Entries) Warn(message string) {
	log.WithFields(logrus.Fields(*entries)).Warn(message)
	entries.clear()
}

func (entries *Entries) Error(message interface{}) {
	log.WithFields(logrus.Fields(*entries)).Error(message)
	entries.clear()
}

func (entries *Entries) clear() {
	*entries = Entries{}
}
