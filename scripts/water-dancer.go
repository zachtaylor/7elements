package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/target"
	"github.com/zachtaylor/7elements/game/update"
)

const WaterDancerID = "water-dancer"

func init() {
	game.Scripts[WaterDancerID] = WaterDancer
}

func WaterDancer(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	if len(args) < 1 {
		err = ErrNoTarget
	} else if token, _err := target.PresentBeing(g, s, args[0]); _err != nil {
		err = _err
	} else if token == nil {
		err = ErrBadTarget
	} else {
		token.IsAwake = false
		update.Token(g, token)
	}
	return
}
