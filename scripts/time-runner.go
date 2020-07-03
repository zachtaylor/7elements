package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/out"
)

func init() {
	game.Scripts["time-runner"] = TimeRunner
}

func TimeRunner(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	s.DrawCard(1)
	out.GameSeat(g, s.JSON())
	out.GameHand(s.Player, s.Hand.JSON())
	return
}
