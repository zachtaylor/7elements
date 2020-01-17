package request

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/update"
	"ztaylor.me/cast"
)

func connect(g *game.T, s *game.Seat) {
	g.Log().With(cast.JSON{
		"Username": s.Username,
		"State":    g.State.String(),
	}).Debug("engine/connect: seated")
	update.Connect(g, s)
	if connector, ok := g.State.R.(game.ConnectStater); ok {
		connector.OnConnect(g, s)
	}
}
