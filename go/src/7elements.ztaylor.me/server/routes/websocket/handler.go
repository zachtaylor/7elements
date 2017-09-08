package websocket

import (
	"7elements.ztaylor.me/server/sessionman"
	"golang.org/x/net/websocket"
	"ztaylor.me/log"
)

var Handler = websocket.Handler(func(conn *websocket.Conn) {
	log := log.Add("RemoteAddr", conn.Request().RemoteAddr)
	session, err := sessionman.ReadRequestCookie(conn.Request())
	if session == nil {
		log.Add("Error", err).Warn("websocket: session error")
		return
	}

	sessionman.OpenSocket(session, conn)
})
