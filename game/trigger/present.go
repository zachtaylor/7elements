package trigger

import (
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/log"
)

// SeatPresent returns stacking events for each card in seat.Present for a trigger name
func SeatPresent(g *game.T, seat *game.Seat, trigger string) []game.Event {
	events := make([]game.Event, 0)
	log := g.Log().With(log.Fields{
		"Seat":    seat.Username,
		"Trigger": trigger,
	}).Tag("game/trigger-seat")
	for _, c := range seat.Present {
		if e := g.Runtime.Service.CardTriggeredEvents(g, seat, c, trigger, c); len(e) > 0 {
			log.Copy().Add("Card", c).Debug("trigger")
			events = append(events, e...)
		}
	}
	return events
}
