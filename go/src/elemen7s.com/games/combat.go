package games

func Combat(game *Game, acard *Card, dcard *Card) {
	if acard.CardBody == nil || dcard.CardBody == nil {
		return
	}
	Damage(game, acard, dcard.CardBody.Attack)
	Damage(game, dcard, acard.CardBody.Attack)
}

func Damage(game *Game, card *Card, n int) {
	card.CardBody.Health -= n
	if card.Health < 1 {
		seat := game.GetSeat(card.Username)
		delete(seat.Alive, card.Id)

		if !card.SkipPast {
			seat.Graveyard[card.Id] = card
		}
	}
}
