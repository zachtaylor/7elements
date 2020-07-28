package runtime

import (
	"ztaylor.me/cast"
	"ztaylor.me/http/websocket"
)

func (t *T) Ping() {
	data := cast.BytesS(websocket.NewMessage("/ping", cast.JSON{
		"ping":   t.Sockets.Count(),
		"online": t.Sessions.Count(),
	}).JSON().String())
	for _, key := range t.Sockets.Keys() {
		if socket := t.Sockets.Get(key); socket != nil {
			socket.Write(data)
		}
	}
}
