package trigger

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/update"
	"ztaylor.me/cast"
)

func Damage(g *game.T, t *game.Token, n int) []game.Stater {
	t.Body.Health -= n
	update.GameChat(g, t.Card.Card.Name, cast.StringI(n)+" damage to "+t.Card.Card.Name)
	update.Token(g, t)
	if t.Body.Health < 1 {
		return Death(g, t)
	}
	return nil
}
