package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/out"
)

const vinespiritID = "vine-spirit"

func init() {
	game.Scripts[vinespiritID] = VineSpirit
}

func VineSpirit(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	if token, ok := me.(*game.Token); !ok || token == nil {
		err = ErrMeToken
	} else {
		token.Body.Attack++
		out.GameToken(g, token.JSON())
	}
	return
}
