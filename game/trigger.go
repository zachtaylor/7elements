package game

import "ztaylor.me/log"

// SeatTriggeredEvents returns stacking events for each card in seat.Present for a trigger name
func SeatTriggeredEvents(g *T, seat *Seat, trigger string) []Event {
	events := make([]Event, 0)
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

func TriggerDamage(g *T, card *Card, n int) []Event {
	card.Body.Health -= n
	g.SendAll(BuildCardUpdate(card))
	if card.Body.Health < 1 {
		card.Body.Health = 0
		seat := g.GetSeat(card.Username)
		delete(seat.Present, card.Id)
		return g.Runtime.Service.CardTriggeredEvents(g, seat, card, "death", card)
	}
	return nil
}
