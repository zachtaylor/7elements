package checktarget

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/seat"
)

func PlayerPresentBeing(game *game.T, seat *seat.T, arg interface{}) (interface{}, error) {
	if arg == "self" {
		return seat, nil
	} else if arg == "enemy" {
		return game.Seats.GetOpponent(seat.Username), nil
	}
	return PresentBeing(game, seat, arg)
}
