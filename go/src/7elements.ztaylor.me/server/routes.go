package server

import (
	"7elements.ztaylor.me/server/api"
	"7elements.ztaylor.me/server/cssbuilder"
	"7elements.ztaylor.me/server/jsbuilder"
	"7elements.ztaylor.me/server/routes/cards"
	"7elements.ztaylor.me/server/routes/login"
	"7elements.ztaylor.me/server/routes/logout"
	"7elements.ztaylor.me/server/routes/signup"
	"7elements.ztaylor.me/server/routes/websocket"
)

func init() {
	HandleFunc(`/7elements\.js`, jsbuilder.Handler)
	HandleFunc(`/7elements\.css`, cssbuilder.Handler)

	HandleFunc(`/api/myaccount\.json`, api.MyAccountJsonHandler)
	HandleFunc(`/api/mycards\.json`, api.MyCardsJsonHandler)
	HandleFunc(`/api/navbar\.json`, api.NavbarJsonHandler)
	HandleFunc(`/api/decks\.json`, api.DecksJsonHandler)
	HandleFunc(`/api/decks/.*\.json`, api.DecksIdJsonHandler)
	HandleFunc(`/api/openpack\.json`, api.OpenPackJsonHandler)
	HandleFunc(`/api/newgame\.json(\?deckid=\d*)?`, api.NewGameJsonHandler)

	HandleFunc(`/api/signup`, signup.Handler)
	HandleFunc(`/api/login`, login.Handler)
	HandleFunc(`/api/logout`, logout.Handler)
	Handler(`/api/websocket`, websocket.Handler)
	HandleFunc(`/api/cards.*\.json`, cards.Handler)
	HandleFunc(`/.*`, PageHandler)
}
