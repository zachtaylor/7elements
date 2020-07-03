package trigger

import (
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/out"
	"ztaylor.me/cast"
)

func Damage(g *game.T, t *game.Token, n int) []game.Stater {
	t.Body.Health -= n
	g.Settings.Chat.AddMessage(chat.NewMessage(t.Card.Proto.Name, cast.StringI(n)+" damage to "+t.Card.Proto.Name))
	out.GameToken(g, t.JSON())
	if t.Body.Health < 1 {
		return Death(g, t)
	}
	return nil
}
