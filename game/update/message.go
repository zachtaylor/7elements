package update

import "ztaylor.me/cast"

func Message(writer Writer, uri string, data cast.JSON) {
	writer.WriteJSON(Build(uri, data))
}
