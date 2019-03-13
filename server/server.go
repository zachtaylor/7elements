package server // import "github.com/zachtaylor/7elements/server"

import (
	"net/http"

	"github.com/zachtaylor/7elements/server/api"
	"ztaylor.me/http/handlers"
	"ztaylor.me/http/mux"
	"ztaylor.me/http/routers"
	"ztaylor.me/http/sessions"
	"ztaylor.me/http/ws"
)

// Server creates routes of this server to a Mux
func Server(fs http.FileSystem, sessions *sessions.Service, dbsalt string) (m mux.Mux) {
	m.Add(mux.RouterPath(`/api/ping.json`), api.PingHandler(sessions))
	m.Add(mux.RouterPath(`/api/myaccount.json`), api.MyAccountHandler(sessions))
	m.Add(mux.RouterPath(`/api/newgame.json`), api.NewGameHandler(sessions))
	m.Add(mux.RouterPath(`/api/login`), api.LoginHandler(sessions, dbsalt))
	m.Add(mux.RouterPath(`/api/signup`), api.SignupHandler(sessions, dbsalt))
	m.Add(mux.RouterPath(`/api/logout`), api.LogoutHandler(sessions))
	m.Add(mux.RouterPath(`/api/websocket`), NewWebsocketHandler(sessions))

	// Server= append(mux, gops.New(gops.RouterPath(`/api/packs.json`),  api.PacksHandler))
	// Server.MapRawLit(`/api/openpack.json`, api.OpenPackJsonHandler)
	// Server= append(mux, gops.New(gops.RouterPath(`/chat`),  api.ChatHandler))
	// Server= append(mux, gops.New(gops.RouterPath(`/join`),  api.JoinHandler))
	// Server= append(mux, gops.New(gops.RouterPath(`/game`),  api.GameHandler))
	// Server= append(mux, gops.New(gops.RouterPath(`/api/cards.*\.json`),  api.CardsHandler))

	// Server.MapRawRgx(`/api/decks/.*\.json`, api.DecksIdJsonHandler)

	m.Add(routers.SinglePageApp, handlers.Index(fs))
	fsHandler := http.FileServer(fs)
	m.Add(mux.RouterPathStarts(`/img/`), handlers.AddPrefix("/assets", fsHandler)) // fixes angular asset layout
	m.Add(mux.RouterAny(true), fsHandler)
	return
}

func NewWebsocketHandler(sessions *sessions.Service) http.Handler {
	mux := make(ws.Mux, 0)
	mux = append(mux, ws.Route{ws.RouterLit("/chat"), api.WSChat()})
	mux = append(mux, ws.Route{ws.RouterLit("/chat/join"), api.WSChatJoin()})
	mux = append(mux, ws.Route{ws.RouterLit("/game"), api.WSGame()})
	mux = append(mux, ws.Route{ws.RouterLit("/game/join"), api.WSGameJoin()})

	return ws.UpgradeHandler(sessions, mux)
}
