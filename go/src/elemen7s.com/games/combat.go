package games

func Combat(g *Game, acard *Card, dcard *Card) {
	if acard.Body == nil || dcard.Body == nil {
		return
	}
	acard.Body.Health -= dcard.Body.Attack
	dcard.Body.Health -= acard.Body.Attack
	if acard.Health < 1 {
		delete(g.GetSeat(acard.Username).Alive, acard.Id)
	}
}
