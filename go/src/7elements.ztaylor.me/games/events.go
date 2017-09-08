package games

import (
	"7elements.ztaylor.me/event"
	"7elements.ztaylor.me/server/sessionman"
	"ztaylor.me/json"
)

func init() {
	event.On("WebsocketOpen", func(args ...interface{}) {
		s := args[0].(*sessionman.Socket)
		WebsocketOpen(s)
	})

	event.On("WebsocketMessage", func(args ...interface{}) {
		s := args[0].(*sessionman.Socket)
		name := args[1].(string)
		data := args[2].(json.Json)
		WebsocketMessage(s, name, data)
	})

	event.On("WebsocketClose", func(args ...interface{}) {
		s := args[0].(*sessionman.Socket)
		WebsocketClose(s)
	})
}
