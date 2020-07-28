package trigger

import (
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
)

// AllPresent returns stacking events for each token in seat.Present for a trigger name
func AllPresent(g *game.T, seat *game.Seat, trigger string) []game.Stater {
	events := make([]game.Stater, 0)
	log := g.Log().With(cast.JSON{
		"Seat":    seat.Username,
		"Trigger": trigger,
	})
	for _, token := range seat.Present {
		if e := g.Settings.Engine.TriggerTokenEvent(g, seat, token, trigger); len(e) > 0 {
			log.Copy().Add("Token", trigger).Debug("trigger")
			events = append(events, e...)
		}
	}
	return events
}
