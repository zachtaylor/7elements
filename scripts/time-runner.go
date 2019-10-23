package scripts

import (
	"github.com/zachtaylor/7elements/game"
)

func init() {
	game.Scripts["time-runner"] = TimeRunner
}

func TimeRunner(g *game.T, seat *game.Seat, target interface{}) []game.Event {
	seat.DrawCard(1)
	seat.SendHandUpdate()
	g.SendSeatUpdate(seat)
	return nil
}
