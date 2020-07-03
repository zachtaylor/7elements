package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/target"
	"github.com/zachtaylor/7elements/out"
)

const energizeID = "energize"

func init() {
	game.Scripts[energizeID] = Energize
}

func Energize(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	var token *game.Token
	if len(args) < 1 {
		err = ErrNoTarget
	} else if token, err = target.PresentBeingItem(g, s, args[0]); err != nil {

	} else {
		token.IsAwake = true
		out.GameToken(g, token.JSON())
	}
	return
}
