package player

import "ztaylor.me/http/websocket"

type Settings struct {
	Sockets *websocket.Cache
}

func NewSettings(sockets *websocket.Cache) Settings {
	return Settings{
		Sockets: sockets,
	}
}
