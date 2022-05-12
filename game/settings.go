package game

import (
	"github.com/zachtaylor/7elements/content"
	"taylz.io/keygen"
	"taylz.io/log"
)

type Settings struct {
	// Content is game content pointer
	Content *content.T
	// LogDir is a system dir path for game log files with trailing slash
	LogDir string
	// Logger is a pointer to the system logger
	Logger *log.T
	// Keygen proposes new game ids
	Keygen keygen.Func
}

func NewSettings(content *content.T, logdir string, syslog *log.T, keygen keygen.Func) Settings {
	return Settings{
		Content: content,
		LogDir:  logdir,
		Logger:  syslog,
		Keygen:  keygen,
	}
}
