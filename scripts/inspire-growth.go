package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/target"
	"github.com/zachtaylor/7elements/game/update"
)

const inspiregrowthID = "inspire-growth"

func init() {
	game.Scripts[inspiregrowthID] = InspireGrowth
}

func InspireGrowth(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	if len(args) < 1 {
		err = ErrNoTarget
	} else if token, e := target.PresentBeing(g, s, args[0]); e != nil {
		err = e
	} else if token == nil {
		err = ErrNoTarget
	} else {
		token.Body.Attack++
		update.Token(g, token)
	}
	return
}
