package ai

import "time"

type Settings struct {
	Delay time.Duration
	Aggro bool
}

func DefaultSettings() Settings {
	return Settings{
		Delay: 2 * time.Second,
	}
}
