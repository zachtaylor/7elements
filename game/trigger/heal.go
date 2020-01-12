package trigger

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/update"
)

func Heal(g *game.T, t *game.Token, n int) []game.Stater {
	t.Body.Health += n
	update.Token(g, t)
	return nil
}
