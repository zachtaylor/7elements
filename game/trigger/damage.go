package trigger

import (
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
)

func Damage(g *game.T, card *game.Card, n int) []game.Event {
	card.Body.Health -= n
	go g.GetChat().AddMessage(chat.NewMessage(card.Card.Name, cast.StringI(n)+" damage to "+card.Card.Name))
	if card.Body.Health < 1 {
		return Death(g, card)
	}
	g.SendCardUpdate(card)
	return nil
}
