package trigger

import (
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
)

// Name returns stacking events for each card in seat.Present for a trigger name
func Name(g *game.T, seat *game.Seat, trigger string) []game.Stater {
	events := make([]game.Stater, 0)
	log := g.Log().With(cast.JSON{
		"Seat":    seat.Username,
		"Trigger": trigger,
	}).Tag("game/trigger/name")
	for _, t := range seat.Present {
		if e := g.Runtime.Service.Trigger(g, seat, t, trigger, t); len(e) > 0 {
			log.Copy().Add("Token", t).Debug("trigger")
			events = append(events, e...)
		}
	}
	return events
}
