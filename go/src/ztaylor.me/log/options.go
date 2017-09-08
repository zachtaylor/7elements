package log

import (
	"github.com/Sirupsen/logrus"
	"os"
)

var logger = logrus.New()

func SetLevel(level string) {
	switch level {
	case "debug":
		logger.Level = logrus.DebugLevel
		break
	case "info":
		logger.Level = logrus.InfoLevel
		break
	case "warn":
		logger.Level = logrus.WarnLevel
		break
	case "error":
		logger.Level = logrus.ErrorLevel
		break
	case "fatal":
		logger.Level = logrus.FatalLevel
		break
	case "panic":
		logger.Level = logrus.PanicLevel
		break
	default:
		logger.Level = logrus.InfoLevel
		Add("Level", level).Warn("log: level invalid")
		break
	}
}

func SetFile(path string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		Info("log: proceeding on file")
		logger.Out = file
	} else {
		Warn("failed to log to file, using default stderr")
	}
}
