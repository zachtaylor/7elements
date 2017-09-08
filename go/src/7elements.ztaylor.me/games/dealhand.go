package games

import (
	"7elements.ztaylor.me/games/cards"
	"7elements.ztaylor.me/games/seats"
	"ztaylor.me/log"
)

func DealHand(game *Game, seat *gameseats.GameSeat) {
	if game.GamePhase != GPHSbegin {
		log.Add("GamePhase", game.GamePhase).Warn("dealhand: rejected")
		return
	}

	for _, card := range seat.Hand {
		seat.Deck.Append(card)
	}
	seat.Hand = make([]*gamecards.GameCard, 0)

	seat.Deck.Shuffle()

	for i := 0; i < 3; i++ {
		seat.Hand = append(seat.Hand, seat.Deck.Draw())
	}

	log.Add("Username", seat.Username).Add("GameId", game.Id).Debug("dealhand")
}
