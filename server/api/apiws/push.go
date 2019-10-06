package apiws

import (
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
	"ztaylor.me/http/websocket"
)

func pushJSON(socket *websocket.T, uri string, data cast.JSON) {
	socket.Write(cast.BytesS(game.BuildPushJSON(uri, data).String()))
}

func pushError(socket *websocket.T, error string) {
	pushJSON(socket, "/error", cast.JSON{
		"error": error,
	})
}

func pushRedirectJSON(socket *websocket.T, location string) {
	pushJSON(socket, "/redirect", cast.JSON{
		"location": location,
	})
}
