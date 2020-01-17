package request

import (
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
)

func disconnect(g *game.T, s *game.Seat) {
	g.Log().With(cast.JSON{
		"Username": s.Username,
		"State":    g.State,
	}).Tag("engine/disconnect: left")

	if disconnector, ok := g.State.R.(game.DisconnectStater); ok {
		disconnector.OnDisconnect(g, s)
	}
}
