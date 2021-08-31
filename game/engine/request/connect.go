package request

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/phase"
	"github.com/zachtaylor/7elements/game/seat"
)

func Connect(game *game.T, seat *seat.T) {
	if seat != nil {
		game.SendData(seat.Username)
	}
	phase.TryOnConnect(game, seat)
}
