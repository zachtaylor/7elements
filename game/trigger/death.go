package trigger

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/update"
)

func Death(g *game.T, t *game.Token) []game.Stater {
	if t.Body != nil {
		t.Body.Health = 0
	}
	seat := g.GetSeat(t.Username)
	update.GameChat(g, t.Card.Proto.Name, "died")
	delete(seat.Present, t.ID)
	return g.Runtime.Service.Trigger(g, seat, t, "death", t)
}
