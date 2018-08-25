package engine

import (
	"github.com/zachtaylor/7elements"
	"time"
)

func Run(game *vii.Game) {
	p1 := game.GetOpponentSeat("").Username
	t := NewTimeline(p1, game.Settings.Timeout, Start(game))
	for {
		tStart := time.Now()
		select {
		case r, ok := <-game.In:
			log := game.Log().Add("Timeline", t)
			if !ok {
				log.Debug("engine-run: timeline finished")
				return
			}
			log.Add("Request", r).Info("engine-run: begin")
			if tf := Request(game, t, r.Username, r.Data); tf != nil {
				log.Add("Future", tf).Info("engine-run: timeline forked")
				t = tf
			}
		case <-time.After(time.Second):
		}
		if !t.HasPause() {
			t.Lifetime -= time.Now().Sub(tStart)
		}
		if game.Results != nil {
			t = t.Fork(game, End(game, t))
		} else if t.Lifetime < 1 {
			t.Lifetime = 0
			t = t.OnStop(game)
		}
	}
}
