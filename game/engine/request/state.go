package request

import (
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
)

func state(g *game.T, seat *game.Seat, json cast.JSON) {
	if requester, ok := g.State.R.(game.RequestStater); ok {
		requester.Request(g, seat, json)
	} else {
		g.Log().With(cast.JSON{
			"State":    g.State.Name(),
			"Username": seat.Username,
		}).Warn("engine/state: request failed")
	}
}
