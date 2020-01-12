package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/update"
)

const vinespiritID = "vine-spirit"

func init() {
	game.Scripts[vinespiritID] = VineSpirit
}

func VineSpirit(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	if token, ok := me.(*game.Token); !ok {
		err = game.ErrMeToken
	} else if token == nil {
		err = game.ErrMeNil
	} else {
		token.Body.Attack++
		update.Token(g, token)
	}
	return
}
