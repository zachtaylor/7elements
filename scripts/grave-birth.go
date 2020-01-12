package scripts

import (
	"github.com/zachtaylor/7elements/game/target"

	"github.com/zachtaylor/7elements/game"
)

const GraveBirthID = "grave-birth"

func init() {
	game.Scripts[GraveBirthID] = GraveBirth
}

func GraveBirth(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	if len(args) < 1 {
		err = game.ErrNoTarget
	} else if card, e := target.MyPastBeing(g, s, args[0]); e != nil {
		err = e
	} else if card == nil {
		err = game.ErrNoTarget
	} else {
		// TODO
	}
	return
}
