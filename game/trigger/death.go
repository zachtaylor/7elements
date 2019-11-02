package trigger

import (
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/game"
)

func Death(g *game.T, card *game.Card) []game.Event {
	if card.Body != nil {
		card.Body.Health = 0
	}
	seat := g.GetSeat(card.Username)
	go g.GetChat().AddMessage(chat.NewMessage(card.Card.Name, "Died"))
	delete(seat.Present, card.Id)
	return g.Runtime.Service.CardTriggeredEvents(g, seat, card, "death", card)
}
