package trigger

import "github.com/zachtaylor/7elements/game"

func Heal(g *game.T, card *game.Card, n int) []game.Event {
	card.Body.Health += n
	g.SendCardUpdate(card)
	return nil
}
