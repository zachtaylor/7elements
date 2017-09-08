package sessionman

import (
	"ztaylor.me/json"
)

type WebsocketMessage struct {
	Name string
	Data json.Json
}

func NewWebsocketMessage() *WebsocketMessage {
	return &WebsocketMessage{
		Data: json.Json{},
	}
}
