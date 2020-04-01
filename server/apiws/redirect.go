package apiws

import (
	"ztaylor.me/cast"
	"ztaylor.me/http/websocket"
)

// redirect sends a "/redirect" Message
//
// path is expected to be like "/login" or something
func redirect(ws *websocket.T, location string) {
	ws.Message("/redirect", cast.JSON{
		"location": location,
	})
}
