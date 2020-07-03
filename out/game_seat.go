package out

import "ztaylor.me/cast"

func GameSeat(t Target, data cast.JSON) {
	t.Send("/game/seat", data)
}
