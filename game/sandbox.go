package game

import (
	"time"

	"taylz.io/log"
)

func Sandbox(game *G, syslog log.Writer, runner Runner) {
	tStart := time.Now()
	defer func() {
		if err := recover(); err != nil {
			syslog.Add("GameID", game.ID()).Error("recovered", time.Since(tStart), err)
		}
	}()
	syslog.Add("GameID", game.ID()).Info("start")
	runner.Run(game, syslog)
	syslog.Add("GameID", game.ID()).Info("stopped", time.Since(tStart))
}
