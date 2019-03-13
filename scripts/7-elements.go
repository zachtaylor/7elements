package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/engine"
)

func init() {
	engine.Scripts["7-elements"] = Elemen7s
}

func Elemen7s(game *game.T, seat *game.Seat, target interface{}) game.Event {
	return engine.End(game, seat.Username, game.GetOpponentSeat(seat.Username).Username)
}
