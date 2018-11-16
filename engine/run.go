package engine

import (
	"time"

	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"ztaylor.me/log"
)

func Run(game *vii.Game) {
	if game == nil {
		log.Error("engine/run: game is nil")
		return
	} else if len(game.Seats) < 2 {
		log.Add("Seats", len(game.Seats)).Error("engine/run: game missing participants")
		return
	}

	p1 := game.GetOpponentSeat("").Username
	game.State = vii.NewGameState(p1, game.Settings, Start(game))
	game.State.Event.OnStart(game)
	one2nd := time.NewTimer(time.Second)
	var tStart time.Time

	for {
		tStart = time.Now()
		select {
		case r, ok := <-game.In:
			// stop and/or flush one2nd channel
			if !one2nd.Stop() {
				<-one2nd.C
			}

			log := game.Log().Add("GameEvent", game.State.EventName())
			if !ok {
				log.Debug("engine/run: timeline finished")
				return
			}
			log.Add("Request", r).Info("engine/run: begin")
			if event := Request(game, game.GetSeat(r.Username), r.Data); event != nil {
				log.Add("Future", event.Name()).Info("engine/run: timeline forked")
				game.State = vii.NewGameState(game.State.Seat, game.Settings, event)
				event.OnStart(game)
			}
		case <-one2nd.C:
		}

		game.State.Timer -= time.Now().Sub(tStart)
		// } else if !t.HasPause() {

		if game.State.Timer < time.Second {
			game.State.Timer = 0
			event := game.State.Event.NextEvent(game)
			game.State = vii.NewGameState(game.State.Seat, game.Settings, event)
			game.State.Event.OnStart(game)
			animate.GameState(game)
		}
		// if game.Results != nil {
		// game.State.Timer = 0
		// game.State.Event.OnStop(game)
		// game.State = vii.NewGameState("", game.Settings, End(game))
		// game.State.Event.OnStart(game)
		// return
		// }

		one2nd.Reset(time.Second)
	}
}
