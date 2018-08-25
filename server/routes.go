package server

import (
	"github.com/zachtaylor/7elements/server/api"
	"ztaylor.me/env"
	"ztaylor.me/http"
	"ztaylor.me/vfs"
)

func init() {
	http.MapLit(`/ping`, api.PingHandler)
	http.MapLit(`/chat`, api.ChatHandler)
	http.MapLit(`/newgame`, api.NewGameHandler)
	http.MapLit(`/join`, api.JoinHandler)
	http.MapLit(`/game`, api.GameHandler)
	http.MapLit(`/api/coins`, api.CoinsHandler)
	http.MapLit(`/api/myaccount.json`, api.MyAccountHandler)
	http.MapLit(`/api/mycards.json`, api.MyCardsHandler)
	http.MapLit(`/api/decks.json`, api.DecksHandler)
	http.MapLit(`/api/buypack.json`, api.PackHandler)

	http.MapRgx(`/api/cards.*\.json`, api.CardsHandler)

	http.MapRawLit(`/7elements.js`, vfs.NewFile(vfs.JS, vfs.NewDirSource(env.Default("JS_PATH", "js/"))))
	http.MapRawLit(`/7elements.css`, vfs.NewFile(vfs.CSS, vfs.NewDirSource(env.Default("CSS_PATH", "css/"))))
	http.MapRawLit(`/api/websocket`, http.SocketHandler)
	http.MapRawLit(`/api/login`, api.LoginHandler)
	http.MapRawLit(`/api/signup`, api.SignupHandler)
	http.MapRawLit(`/api/logout`, api.LogoutHandler)
	http.MapRawLit(`/api/openpack.json`, api.OpenPackJsonHandler)

	imgHandler := http.StripPrefix("/img/", http.FileServer(http.Dir(env.Default("IMG_PATH", "img/"))))
	http.MapRawRgx(`.*\.png`, imgHandler)
	http.MapRawRgx(`.*\.jpg`, imgHandler)

	http.MapRawRgx(`/api/decks/.*\.json`, api.DecksIdJsonHandler)
	http.MapRawRgx(`.*`, PageHandler)
}
