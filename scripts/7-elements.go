package scripts

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/engine"
)

func init() {
	engine.Scripts["7-elements"] = Elemen7s
}

func Elemen7s(game *vii.Game, seat *vii.GameSeat, target interface{}) vii.GameEvent {
	return engine.End(game, seat.Username, game.GetOpponentSeat(seat.Username).Username)
}
