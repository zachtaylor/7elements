package update

import (
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
)

func Hand(seat *game.Seat) {
	seat.WriteJSON(Build("/game/hand", cast.JSON{
		"cards": seat.Hand.JSON(),
	}))
}
