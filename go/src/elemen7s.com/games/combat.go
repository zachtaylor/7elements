package games

func Combat(g *Game, acard *Card, dcard *Card) {
	if acard.Body == nil || dcard.Body == nil {
		return
	}
	acard.Body.Health -= dcard.Body.Attack
	dcard.Body.Health -= acard.Body.Attack
	if acard.Health < 1 {
		seat := g.GetSeat(acard.Username)
		delete(seat.Alive, acard.Id)
		seat.Graveyard[acard.Id] = acard
	}
	if dcard.Health < 1 {
		seat := g.GetSeat(dcard.Username)
		delete(seat.Alive, dcard.Id)
		seat.Graveyard[dcard.Id] = dcard
	}
}
