package trigger

import "github.com/zachtaylor/7elements/game"

func Damage(g *game.T, card *game.Card, n int) []game.Event {
	card.Body.Health -= n
	g.SendCardUpdate(card)
	if card.Body.Health < 1 {
		return Death(g, card)
	}
	return nil
}
