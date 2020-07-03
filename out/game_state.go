package out

import "ztaylor.me/cast"

func GameState(t Target, statejson cast.JSON) {
	t.Send("/game/state", statejson)
}
