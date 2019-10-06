package server

import (
	"net/http"

	"github.com/zachtaylor/7elements/server/api"
	"github.com/zachtaylor/7elements/server/api/apiws"
	"ztaylor.me/http/handler"
	"ztaylor.me/http/mux"
	"ztaylor.me/http/router"
	"ztaylor.me/http/websocket"
	"ztaylor.me/log"
)

func Routes(rt *Runtime) mux.Mux {
	m := mux.Mux{}
	m.Add(router.Path(`/api/global.json`), api.GlobalDataHandler(rt))
	m.Add(router.Path(`/api/myaccount.json`), api.MyAccountHandler(rt))
	m.Add(router.Path(`/api/newgame.json`), api.NewGameHandler(rt))
	m.Add(router.Path(`/api/login`), api.LoginHandler(rt))
	m.Add(router.Path(`/api/signup`), api.SignupHandler(rt))
	m.Add(router.Path(`/api/logout`), api.LogoutHandler(rt))
	m.Add(router.Path(`/api/websocket`), WSRoutes(rt))
	m.Add(router.SinglePageApp, handler.Index(rt.FileSystem))
	fsHandler := http.FileServer(rt.FileSystem)
	m.Add(router.PathStarts(`/img/`), handler.AddPrefix("/assets", fsHandler)) // fix angular asset layout
	m.Add(router.Bool(true), fsHandler)
	return m
}

func WSRoutes(rt *Runtime) http.Handler {
	mux := websocket.NewCache(rt.Sessions)
	mux.Route(&websocket.Route{websocket.RouterLit("/connect"), apiws.Connect(rt)})
	mux.Route(&websocket.Route{websocket.RouterLit("/disconnect"), apiws.Disconnect(rt)})
	mux.Route(&websocket.Route{websocket.RouterLit("/chat"), apiws.Chat(rt)})
	mux.Route(&websocket.Route{websocket.RouterLit("/chat/join"), apiws.ChatJoin(rt)})
	mux.Route(&websocket.Route{websocket.RouterLit("/game"), apiws.Game(rt)})
	mux.Route(&websocket.Route{websocket.RouterLit("/game/new"), apiws.GameNew(rt)})
	mux.Route(&websocket.Route{websocket.RouterLit("/packs/buy"), apiws.PacksBuy(rt)})
	mux.Route(&websocket.Route{websocket.RouterLit("/signup"), apiws.Signup(rt)})
	mux.Route(&websocket.Route{websocket.RouterLit("/login"), apiws.Login(rt)})
	mux.Route(&websocket.Route{websocket.RouterLit("/logout"), apiws.Logout(rt)})

	// route 404
	mux.Route(&websocket.Route{
		websocket.RouterFunc(func(*websocket.Message) bool {
			return true
		}),
		websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
			rt.Root.Logger.New().With(log.Fields{
				"Username": m.User,
				"URI":      m.URI,
			}).Warn("api/ws: message routing failed")
		}),
	})

	return websocket.UpgradeHandler(mux)
}
