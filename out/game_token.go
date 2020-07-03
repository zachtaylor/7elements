package out

import "ztaylor.me/cast"

func GameToken(t Target, tokenjson cast.JSON) {
	t.Send("/game/token", tokenjson)
}
