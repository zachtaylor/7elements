package game

type Settings struct {
	// LogDir is a system dir path for game log files with trailing slash
	LogDir string
}

func NewSettings(logdir string) Settings {
	return Settings{
		LogDir: logdir,
	}
}
