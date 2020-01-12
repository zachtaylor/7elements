package update

import "github.com/zachtaylor/7elements/game"

// Connect populates game data in multiple writes to stay under ws frame limit
func Connect(g *game.T, seat *game.Seat) {
	// seat.Send(Build("/connect/game", nil))
	seat.WriteJSON(Build("/game", g.JSON(seat)))
	for _, s := range g.Seats {
		for _, t := range s.Present {
			seat.WriteJSON(Build("/game/token", t.JSON()))
		}
		for _, c := range s.Past {
			seat.WriteJSON(Build("/game/card", c.JSON()))
		}
	}
}
