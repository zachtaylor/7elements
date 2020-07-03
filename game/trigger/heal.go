package trigger

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/out"
)

func Heal(g *game.T, t *game.Token, n int) []game.Stater {
	t.Body.Health += n
	out.GameToken(g, t.JSON())
	return nil
}
