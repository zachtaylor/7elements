package game

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/out"
)

// Data populates game data in multiple writes to stay under ws frame limit
func Data(player out.Player, g *game.T, seat *game.Seat) {
	// seat.Send(Build("/connect/game", nil))
	player.Send("/game", g.JSON(seat))
	for _, s := range g.Seats {
		for _, t := range s.Present {
			player.Send("/game/token", t.JSON())
		}
		for _, c := range s.Past {
			player.Send("/game/card", c.JSON())
		}
	}
}
