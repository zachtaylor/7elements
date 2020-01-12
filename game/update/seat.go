package update

import "github.com/zachtaylor/7elements/game"

// Seat sends game seat data using SendUpdate
func Seat(g *game.T, seat *game.Seat) {
	g.Log().Add("Username", seat.Username).Source().Debug()
	Game(g, "/game/seat", seat.JSON())
}
