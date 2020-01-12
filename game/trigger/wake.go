package trigger

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/update"
)

func Wake(g *game.T, token *game.Token) []game.Stater {
	wasAwake := token.IsAwake
	token.IsAwake = true
	if !wasAwake {
		update.Token(g, token)
	}
	return []game.Stater{}
}
