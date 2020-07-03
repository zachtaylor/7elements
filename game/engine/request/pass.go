package request

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/out"
	"ztaylor.me/cast"
)

func pass(g *game.T, seat *game.Seat, json cast.JSON) {
	log := g.Log().With(cast.JSON{
		"State":    g.State.String(),
		"Username": seat.Username,
	}).Tag("engine/pass")
	if pass := json.GetS("pass"); pass == "" {
		log.Warn("target missing")
	} else if pass != g.State.ID() {
		log.Add("PassID", pass).Warn("target mismatch")
	} else if len(g.State.Reacts[seat.Username]) > 0 {
		out.GameError(seat.Player, "pass", "response already recorded")
	} else {
		g.State.Reacts[seat.Username] = "pass"
		out.GameReact(g, g.State.ID(), seat.Username, g.State.Reacts[seat.Username], g.State.Timer)
	}
}
