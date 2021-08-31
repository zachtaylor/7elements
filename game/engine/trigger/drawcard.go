package trigger

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/wsout"
)

func DrawCard(game *game.T, seat *seat.T, count int) []game.Phaser {
	newcards := make(card.Set)
	for i := 0; i < count && len(seat.Deck.Cards) > 0; i++ {
		if card := seat.Deck.Draw(); card != nil {
			newcards[card.ID] = card
			seat.Hand[card.ID] = card
		}
	}
	if len(newcards) > 0 {
		game.Seats.Write(wsout.GameSeatJSON(seat.Data()))
		seat.Writer.Write(wsout.GameHand(seat.Hand.Keys()).EncodeToJSON())
		for _, card := range newcards {
			seat.Writer.Write(wsout.GameCard(card.Data()).EncodeToJSON())
		}
	}
	return nil // todo
}
