package server

import (
	"github.com/zachtaylor/7elements/server/api"
	"ztaylor.me/env"
	"ztaylor.me/http"
)

func init() {
	Server.MapLit(`/api/ping.json`, api.PingHandler)
	Server.MapLit(`/api/decks.json`, api.DecksHandler)
	// Server.MapLit(`/api/packs.json`, api.PacksHandler)
	// Server.MapRawLit(`/api/openpack.json`, api.OpenPackJsonHandler)

	Server.MapLit(`/chat`, api.ChatHandler)
	Server.MapLit(`/newgame`, api.NewGameHandler)
	Server.MapLit(`/join`, api.JoinHandler)
	Server.MapLit(`/game`, api.GameHandler)

	Server.MapLit(`/api/coins`, api.CoinsHandler)
	Server.MapLit(`/api/myaccount.json`, api.MyAccountHandler)
	Server.MapLit(`/api/mycards.json`, api.MyCardsHandler)
	Server.MapLit(`/api/mydecks.json`, api.MyDecksHandler)

	Server.MapRgx(`/api/cards.*\.json`, api.CardsHandler)

	Server.MapRawLit(`/api/websocket`, http.SocketHandler)
	Server.MapRawLit(`/api/login`, api.LoginHandler)
	Server.MapRawLit(`/api/signup`, api.SignupHandler)
	Server.MapRawLit(`/api/logout`, api.LogoutHandler)

	imgHandler := http.StripPrefix("/img/", http.FileServer(http.Dir(env.Default("IMG_PATH", "img/"))))
	Server.MapRawRgx(`.*\.png`, imgHandler)
	Server.MapRawRgx(`.*\.jpg`, imgHandler)

	Server.MapRawRgx(`/api/decks/.*\.json`, api.DecksIdJsonHandler)
	Server.MapRawRgx(`.*`, PageHandler)
}
