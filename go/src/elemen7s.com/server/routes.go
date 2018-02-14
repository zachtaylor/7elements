package server

import (
	"elemen7s.com/options"
	"elemen7s.com/server/api"
	nhttp "net/http"
	"ztaylor.me/buildir"
	"ztaylor.me/env"
	"ztaylor.me/http"
)

func init() {
	http.MapLit(`/ping`, api.PingHandler)
	http.MapLit(`/chat`, api.ChatHandler)
	http.MapLit(`/newgame`, api.NewGameHandler)
	http.MapLit(`/join`, api.JoinHandler)
	http.MapLit(`/api/coins`, api.CoinsHandler)
	http.MapLit(`/api/myaccount.json`, api.MyAccountHandler)
	http.MapLit(`/api/mycards.json`, api.MyCardsHandler)
	http.MapLit(`/api/decks.json`, api.DecksHandler)
	http.MapLit(`/api/buypack.json`, api.PackHandler)
	http.MapRgx(`/api/cards.*\.json`, api.CardsHandler)

	http.MapRawLit(`/7elements.js`, buildir.GetFile(".js", options.String("js-path")))
	http.MapRawLit(`/7elements.css`, buildir.GetFile(".css", options.String("css-path")))

	http.MapRawLit(`/api/websocket`, http.SocketHandler)
	http.MapRawLit(`/api/login`, api.LoginHandler)
	http.MapRawLit(`/api/signup`, api.SignupHandler)
	http.MapRawLit(`/api/logout`, api.LogoutHandler)
	http.MapRawLit(`/api/openpack.json`, api.OpenPackJsonHandler)

	http.MapRawRgx(`/api/decks/.*\.json`, api.DecksIdJsonHandler)
	imgPath := env.Default("IMG_PATH", "img")
	imgHandler := nhttp.StripPrefix("/"+imgPath+"/", http.FileServer(http.Dir(imgPath)))
	http.MapRawRgx(`.*\.png`, imgHandler)
	http.MapRawRgx(`.*\.jpg`, imgHandler)
	http.MapRawRgx(`.*`, PageHandler)
}
