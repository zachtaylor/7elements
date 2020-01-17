package request

import (
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
)

func InResponse(g *game.T, seat *game.Seat, uri string, json cast.JSON) []game.Stater {
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
	case "play":
		return play(g, seat, json, true)
	default:
		g.Log().Add("Data", json).Source().Warn("404")
	}
	return nil
}
