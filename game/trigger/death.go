package trigger

import "github.com/zachtaylor/7elements/game"

func Death(g *game.T, card *game.Card) []game.Event {
	card.Body.Health = 0
	seat := g.GetSeat(card.Username)
	delete(seat.Present, card.Id)
	return g.Runtime.Service.CardTriggeredEvents(g, seat, card, "death", card)
}
