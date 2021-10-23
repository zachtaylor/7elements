package game

import (
	"github.com/zachtaylor/7elements/card"
	"taylz.io/keygen"
	"taylz.io/log"
)

type Settings struct {
	// LogDir is a system dir path for game log files with trailing slash
	LogDir string
	// Cards is the card pool
	Cards card.Prototypes
	// Logger is a pointer to the system logger
	Logger *log.T
	// Engine is the dependency injection point
	Engine Engine
	// Keygen proposes new game keys
	Keygen keygen.Func
}

func NewSettings(logdir string, cards card.Prototypes, syslog *log.T, engine Engine, keygen keygen.Func) Settings {
	return Settings{
		LogDir: logdir,
		Cards:  cards,
		Logger: syslog,
		Engine: engine,
		Keygen: keygen,
	}
}
