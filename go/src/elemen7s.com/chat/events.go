package chat

import (
	"ztaylor.me/events"
	"ztaylor.me/http"
)

func init() {
	events.On(http.EVTsocket_login, func(args ...interface{}) {
		// session := args[0].(*http.Session)
		socket := args[1].(*http.Socket)
		GetChannel("all").AddSocket(socket)
	})
	events.On(http.EVTsocket_close, func(args ...interface{}) {
		socket := args[0].(*http.Socket)
		if socket.Session != nil {
			GetChannel("all").RemoveSocket(socket)
		}
	})
}
