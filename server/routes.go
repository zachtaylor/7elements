package server

import (
	"net/http"

	"github.com/zachtaylor/7elements/server/api"
	"ztaylor.me/http/handlers"
	"ztaylor.me/http/mux"
	"ztaylor.me/http/ws"
)

// Routes applies the routes of this server to a Mux
func Routes(router *mux.Mux, fs http.FileSystem, dbsalt string) {
	router.MapLit(`/api/ping.json`, api.PingHandler())
	router.MapLit(`/api/myaccount.json`, api.MyAccountHandler())
	router.MapLit(`/api/newgame.json`, api.NewGameHandler())
	router.MapLit(`/api/login`, api.LoginHandler(dbsalt))
	router.MapLit(`/api/signup`, api.SignupHandler(dbsalt))
	router.MapLit(`/api/logout`, api.LogoutHandler)
	router.MapLit(`/api/websocket`, NewWebsocketHandler())

	// Server.MapLit(`/api/packs.json`, api.PacksHandler)
	// Server.MapRawLit(`/api/openpack.json`, api.OpenPackJsonHandler)
	// Server.MapLit(`/chat`, api.ChatHandler)
	// Server.MapLit(`/join`, api.JoinHandler)
	// Server.MapLit(`/game`, api.GameHandler)
	// Server.MapRgx(`/api/cards.*\.json`, api.CardsHandler)

	// Server.MapRawRgx(`/api/decks/.*\.json`, api.DecksIdJsonHandler)

	router.Map(mux.MatcherSPA, handlers.Index(fs))
	fsHandler := http.FileServer(fs)
	router.MapRgx(`/img/`, handlers.AddPrefix("/assets", fsHandler))
	router.MapRgx(`.`, fsHandler)
}

func NewWebsocketHandler() http.Handler {
	mux := ws.NewMux()
	mux.MapLit("/chat", api.WSChat())
	mux.MapLit("/chat/join", api.WSChatJoin())
	mux.MapLit("/game", api.WSGame())
	mux.MapLit("/game/join", api.WSGameJoin())

	return ws.UpgradeHandler(mux)
}
