package target

import "github.com/zachtaylor/7elements/game"

func PlayerPresentBeing(g *game.T, seat *game.Seat, arg interface{}) (interface{}, error) {
	if arg == "self" {
		return seat, nil
	} else if arg == "enemy" {
		return g.GetOpponentSeat(seat.Username), nil
	}
	return PresentBeing(g, seat, arg)
}
