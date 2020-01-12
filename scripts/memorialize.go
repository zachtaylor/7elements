package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/target"
	"github.com/zachtaylor/7elements/game/update"
)

const memorializeID = "memorialize"

func init() {
	game.Scripts[memorializeID] = Memorialize
}

func Memorialize(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	if len(args) < 1 {
		err = game.ErrNoTarget
	} else if card, e := target.MyPastBeing(g, s, args[0]); e != nil {
		err = e
	} else if card == nil {
		err = game.ErrNoTarget
	} else {
		c := game.NewCard(card.Card)
		c.Username = s.Username
		g.RegisterCard(c)
		s.Hand[c.ID] = c
		update.Seat(g, s)
		update.Hand(s)
	}
	return
}
