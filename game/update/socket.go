package update

import (
	"ztaylor.me/cast"
	"ztaylor.me/http/websocket"
)

func Socket(socket *websocket.T, uri string, data cast.JSON) {
	socket.SendMessage(websocket.NewMessage(uri, data))
}
