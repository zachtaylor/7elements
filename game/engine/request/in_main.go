package request

import (
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
)

func InMain(g *game.T, seat *game.Seat, uri string, json cast.JSON) []game.Stater {
	g.Log().Add("Data", json).Source().Debug()
	switch uri {
	case "connect":
		connect(g, seat)
	case "disconnect":
		disconnect(g, seat)
	case "chat":
		chat(g, seat, json)
	case "pass":
		pass(g, seat, json)
	case g.State.ID():
		state(g, seat, json)
	case "trigger":
		return trigger(g, seat, json)
	case "attack":
		return attack(g, seat, json)
	case "play":
		return play(g, seat, json, false)
	default:
		g.Log().With(cast.JSON{
			"Seat":  seat.String(),
			"URI":   uri,
			"State": g.State.Print(),
		}).Warn("engine/request: 404")
	}
	return nil
}
