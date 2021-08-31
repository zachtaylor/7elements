package engine

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/phase"
	"taylz.io/http/websocket"
)

// Stack adds new States as Stacked States
func Stack(g *game.T, stack []game.Phaser) {
	log := g.Log().With(websocket.MsgData{
		"State": g.State,
		"Stack": stack,
	})
	log.Trace()
	next := make([]game.Phaser, 0)
	for _, r := range stack {
		state := g.NewState(r)
		state.Stack = g.State
		g.State = state

		if addnext := phase.TryOnActivate(g); len(addnext) > 0 {
			g.Log().With(websocket.MsgData{
				"State": g.State,
				"Stack": stack,
			}).Debug("activate trigger")
			next = append(next, addnext...)
		}
	}
	if len(next) > 0 {
		log.Debug("Hold my cards, lads, I'm going in!")
		Stack(g, next)
	}
}
