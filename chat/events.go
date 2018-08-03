package chat

import (
	"ztaylor.me/events"
	"ztaylor.me/http"
	"ztaylor.me/log"
)

func init() {
	events.On(http.EVTsocket_login, func(args ...interface{}) {
		socket := args[0].(*http.Socket)
		session := args[1].(*http.Session)
		GetChannel("all").AddSocket(socket)
		log.Add("Session", session).Add("Socket", socket).Debug("chat: added")
	})
	events.On(http.EVTsocket_close, func(args ...interface{}) {
		socket := args[0].(*http.Socket)
		if socket.Session != nil {
			GetChannel("all").RemoveSocket(socket)
			log.Add("Session", socket.Session).Add("Socket", socket).Debug("chat: lost")
		}
	})
}
