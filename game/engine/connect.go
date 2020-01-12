package engine

import (
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
)

func connect(g *game.T, seat *game.Seat) {
	log := g.Log().With(cast.JSON{
		"State": g.State,
		"Seat":  seat,
	}).Source()
	if connector, _ := g.State.R.(game.ConnectStater); connector != nil {
		log.Debug()
		connector.OnConnect(g, seat)
	} else {
		log.Debug("none")
	}
}
