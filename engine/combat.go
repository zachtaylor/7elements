package engine

import "github.com/zachtaylor/7elements"

func Combat(game *vii.Game, acard *vii.GameCard, dcard *vii.GameCard) {
	if acard.Body == nil || dcard.Body == nil {
		return
	}
	Damage(game, acard, dcard.Body.Attack)
	Damage(game, dcard, acard.Body.Attack)
}

func Damage(game *vii.Game, card *vii.GameCard, n int) {
	card.Body.Health -= n
	if card.Body.Health < 1 {
		seat := game.GetSeat(card.Username)
		delete(seat.Alive, card.Id)

		if !card.IsToken {
			seat.Graveyard[card.Id] = card
		}
	}
}
