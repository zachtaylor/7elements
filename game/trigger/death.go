package trigger

import (
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/game"
)

func Death(g *game.T, t *game.Token) []game.Stater {
	if t.Body != nil {
		t.Body.Health = 0
	}
	seat := g.GetSeat(t.Username)
	g.Settings.Chat.AddMessage(chat.NewMessage(t.Card.Proto.Name, "died"))
	delete(seat.Present, t.ID)
	return g.Settings.Engine.TriggerTokenEvent(g, seat, t, "death")
}
