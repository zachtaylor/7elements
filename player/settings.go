package player

import "ztaylor.me/http/websocket"

type Settings struct {
	Sockets websocket.Service
}

func NewSettings(sockets websocket.Service) Settings {
	return Settings{
		Sockets: sockets,
	}
}
