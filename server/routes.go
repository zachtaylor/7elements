package server

import (
	"net/http"

	"github.com/zachtaylor/7elements/server/api"
	"ztaylor.me/http/handler"
	"ztaylor.me/http/mux"
)

// Routes applies the routes of this server to a Mux
func Routes(router *mux.Mux, fs http.FileSystem) {
	router.MapLit(`/api/ping.json`, http.HandlerFunc(api.PingHandler))
	router.MapLit(`/api/myaccount.json`, http.HandlerFunc(api.MyAccountHandler))
	router.MapLit(`/api/login`, api.LoginHandler)
	router.MapLit(`/api/signup`, api.SignupHandler)
	router.MapLit(`/api/logout`, api.LogoutHandler)

	// Server.MapLit(`/api/newgame.json`, api.NewGameHandler)
	// Server.MapLit(`/api/packs.json`, api.PacksHandler)
	// Server.MapRawLit(`/api/openpack.json`, api.OpenPackJsonHandler)
	// Server.MapLit(`/chat`, api.ChatHandler)
	// Server.MapLit(`/join`, api.JoinHandler)
	// Server.MapLit(`/game`, api.GameHandler)
	// Server.MapRgx(`/api/cards.*\.json`, api.CardsHandler)
	// Server.MapRawLit(`/api/websocket`, http.SocketHandler)

	// Server.MapRawRgx(`/api/decks/.*\.json`, api.DecksIdJsonHandler)

	router.Map(mux.MatcherSPA, handler.Index(fs))
	fsHandler := http.FileServer(fs)
	router.MapRgx(`/img/`, handler.AddPrefix("/assets", fsHandler))
	router.MapRgx(`.`, fsHandler)
}
