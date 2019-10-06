package scripts

import (
	"github.com/zachtaylor/7elements/game"
)

func init() {
	game.Scripts["time-runner"] = TimeRunner
}

func TimeRunner(g *game.T, s *game.Seat, target interface{}) []game.Event {
	s.DrawCard(1)
	s.Send(game.BuildHandUpdate(s))
	g.SendAll(game.BuildSeatUpdate(s))
	return nil
}
