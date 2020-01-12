package update

import "ztaylor.me/cast"

type Writer interface {
	WriteJSON(cast.JSON)
}
