package websocket

import (
	"golang.org/x/net/websocket"
	"ztaylor.me/http/sessions"
	"ztaylor.me/log"
)

var Handler = websocket.Handler(func(conn *websocket.Conn) {
	log := log.Add("RemoteAddr", conn.Request().RemoteAddr)
	session, err := sessions.ReadRequestCookie(conn.Request())
	if session == nil {
		log.Add("Error", err).Warn("websocket: session error")
		return
	}

	sessions.OpenSocket(session, conn)
})
