package trigger

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/out"
)

func Wake(g *game.T, token *game.Token) []game.Stater {
	wasAwake := token.IsAwake
	token.IsAwake = true
	if !wasAwake {
		out.GameToken(g, token.JSON())
	}
	return []game.Stater{}
}
