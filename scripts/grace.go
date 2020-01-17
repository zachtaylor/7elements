package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/target"
	"github.com/zachtaylor/7elements/game/trigger"
)

const GraceID = "grace"

func init() {
	game.Scripts[GraceID] = Grace
}

func Grace(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	if len(args) < 1 {
		err = ErrNoTarget
	} else if token, e := target.MyPresentBeing(g, s, args[0]); e != nil {
		err = e
	} else if token == nil {
		err = ErrNoTarget
	} else {
		events = trigger.Heal(g, token, 3)
	}
	return
}
