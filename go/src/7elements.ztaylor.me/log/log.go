package log

import (
	"github.com/Sirupsen/logrus"
	"os"
)

var log = logrus.New()

func SetLevel(level string) {
	switch level {
	case "debug":
		log.Level = logrus.DebugLevel
		break
	case "info":
		log.Level = logrus.InfoLevel
		break
	case "warn":
		log.Level = logrus.WarnLevel
		break
	case "error":
		log.Level = logrus.ErrorLevel
		break
	case "fatal":
		log.Level = logrus.FatalLevel
		break
	case "panic":
		log.Level = logrus.PanicLevel
		break
	default:
		log.Level = logrus.InfoLevel
		globalEntries().Add("Level", level).Warn("log: level invalid")
		break
	}
}

func SetFile(path string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		globalEntries().Info("log: proceeding on file")
		log.Out = file
	} else {
		globalEntries().Warn("failed to log to file, using default stderr")
	}
}
