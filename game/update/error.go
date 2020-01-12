package update

import (
	"ztaylor.me/cast"
	"ztaylor.me/http/websocket"
)

func ErrorW(writer Writer, source, message string) {
	writer.WriteJSON(Build("/game/error", cast.JSON{
		"source":  source,
		"message": message,
	}))
}

func ErrorSock(socket *websocket.T, source, message string) {
	socket.SendMessage(websocket.NewMessage("/game/error", cast.JSON{
		"source":  source,
		"message": message,
	}))
}
