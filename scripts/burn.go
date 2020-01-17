package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/target"
	"github.com/zachtaylor/7elements/game/trigger"
)

const burnID = "burn"

func init() {
	game.Scripts[burnID] = Burn
}

func Burn(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	if len(args) < 1 {
		err = ErrNoTarget
	} else if card, e := target.PresentBeing(g, s, args[0]); e != nil {
		err = e
	} else {
		events = trigger.Damage(g, card, 2)
	}
	return
}
