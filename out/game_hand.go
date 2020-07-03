package out

import "ztaylor.me/cast"

func GameHand(t Target, cardsjson cast.JSON) {
	t.Send("/game/hand", cast.JSON{
		"cards": cardsjson,
	})
}
