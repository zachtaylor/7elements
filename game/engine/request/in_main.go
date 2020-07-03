package request

import (
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
)

func InMain(g *game.T, seat *game.Seat, uri string, json cast.JSON) []game.Stater {
	g.Log().Add("Data", json).Debug()
	switch uri {
	case "connect":
		g.State.Connect(g, seat)
	case "disconnect":
		g.State.Disconnect(g, seat)
	case g.State.ID():
		g.State.Request(g, seat, json)
	case "chat":
		chat(g, seat, json)
	case "pass":
		pass(g, seat, json)
	case "trigger":
		return trigger(g, seat, json)
	case "attack":
		return attack(g, seat, json)
	case "play":
		return play(g, seat, json, false)
	default:
		g.Log().With(cast.JSON{
			"URI":   uri,
			"Seat":  seat,
			"State": g.State,
		}).Warn("engine/request: 404")
	}
	return nil
}
