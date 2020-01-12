package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/update"
)

func init() {
	game.Scripts["time-runner"] = TimeRunner
}

func TimeRunner(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	s.DrawCard(1)
	update.Seat(g, s)
	update.Hand(s)
	return
}
