package game

import (
	"time"

	"taylz.io/log"
)

func Sandbox(syslog *log.T, game *T) {
	tStart := time.Now()
	defer func() {
		if err := recover(); err != nil {
			syslog.With(map[string]interface{}{
				"GameID":  game.id,
				"Error":   err,
				"Runtime": time.Since(tStart),
			}).Error("panic")
		}
	}()
	syslog.Add("GameID", game.id).Info("start")
	game.Engine().Run(syslog, game)
	syslog.With(map[string]interface{}{
		"GameID":  game.id,
		"Runtime": time.Since(tStart),
	}).Info("exit")
}
