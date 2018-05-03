package games

import (
	"elemen7s.com"
)

func Combat(game *Game, acard *vii.GameCard, dcard *vii.GameCard) {
	if acard.CardBody == nil || dcard.CardBody == nil {
		return
	}
	Damage(game, acard, dcard.CardBody.Attack)
	Damage(game, dcard, acard.CardBody.Attack)
}

func Damage(game *Game, card *vii.GameCard, n int) {
	card.CardBody.Health -= n
	if card.Health < 1 {
		seat := game.GetSeat(card.Username)
		delete(seat.Alive, card.Id)

		if !card.IsToken {
			seat.Graveyard[card.Id] = card
		}
	}
}
