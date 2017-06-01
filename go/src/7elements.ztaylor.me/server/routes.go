package server

import (
	"7elements.ztaylor.me/server/cards"
	"7elements.ztaylor.me/server/cssbuilder"
	"7elements.ztaylor.me/server/jsbuilder"
	"7elements.ztaylor.me/server/login"
	"7elements.ztaylor.me/server/logout"
	"7elements.ztaylor.me/server/myaccount"
	"7elements.ztaylor.me/server/mycards"
	"7elements.ztaylor.me/server/signup"
)

func init() {
	HandleFunc(`/7elements.js`, jsbuilder.Handler)
	HandleFunc(`/7elements.css`, cssbuilder.Handler)
	HandleFunc(`/api/signup`, signup.Handler)
	HandleFunc(`/api/login`, login.Handler)
	HandleFunc(`/api/logout`, logout.Handler)
	HandleFunc(`/api/mycards.json`, mycards.Handler)
	HandleFunc(`/api/myaccount.json`, myaccount.Handler)
	HandleFunc(`/cards.*\.json`, cards.Handler)
	// http.Handle("/api/websocket", WebsocketHandler)
	HandleFunc(`/.*`, PageHandler)
}
