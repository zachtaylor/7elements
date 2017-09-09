package games

import (
	"ztaylor.me/events"
	"ztaylor.me/http/sessions"
	"ztaylor.me/json"
)

func init() {
	events.On("WebsocketOpen", func(args ...interface{}) {
		s := args[0].(*sessions.Socket)
		WebsocketOpen(s)
	})

	events.On("WebsocketMessage", func(args ...interface{}) {
		s := args[0].(*sessions.Socket)
		name := args[1].(string)
		data := args[2].(json.Json)
		WebsocketMessage(s, name, data)
	})

	events.On("WebsocketClose", func(args ...interface{}) {
		s := args[0].(*sessions.Socket)
		WebsocketClose(s)
	})
}
