package scripts

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/target"
	"github.com/zachtaylor/7elements/out"
)

const memorializeID = "memorialize"

func init() {
	game.Scripts[memorializeID] = Memorialize
}

func Memorialize(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	if len(args) < 1 {
		err = ErrNoTarget
	} else if c, e := target.MyPastBeing(g, s, args[0]); e != nil {
		err = e
	} else if c == nil {
		err = ErrNoTarget
	} else {
		c := card.New(c.Proto)
		c.Username = s.Username
		g.RegisterCard(c)
		s.Hand[c.ID] = c
		out.GameSeat(g, s.JSON())
		out.GameHand(s.Player, s.Hand.JSON())
	}
	return
}
