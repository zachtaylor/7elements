package server

import (
	"7elements.ztaylor.me/server/cssbuilder"
	"7elements.ztaylor.me/server/jsbuilder"
	"7elements.ztaylor.me/server/routes/cards"
	"7elements.ztaylor.me/server/routes/login"
	"7elements.ztaylor.me/server/routes/logout"
	"7elements.ztaylor.me/server/routes/myaccount"
	"7elements.ztaylor.me/server/routes/mycards"
	"7elements.ztaylor.me/server/routes/openpack"
	"7elements.ztaylor.me/server/routes/signup"
)

func init() {
	HandleFunc(`/7elements.js`, jsbuilder.Handler)
	HandleFunc(`/7elements.css`, cssbuilder.Handler)
	HandleFunc(`/api/signup`, signup.Handler)
	HandleFunc(`/api/login`, login.Handler)
	HandleFunc(`/api/logout`, logout.Handler)
	HandleFunc(`/api/mycards.json`, mycards.Handler)
	HandleFunc(`/api/myaccount.json`, myaccount.Handler)
	HandleFunc(`/api/openpack.json`, openpack.Handler)
	HandleFunc(`/cards.*\.json`, cards.Handler)
	// http.Handle("/api/websocket", WebsocketHandler)
	HandleFunc(`/.*`, PageHandler)
}
