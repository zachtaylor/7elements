package server

import (
	"net/http"

	"github.com/zachtaylor/7elements/server/api"
	"github.com/zachtaylor/7elements/server/apiws"
	"ztaylor.me/cast"
	"ztaylor.me/http/handler"
	"ztaylor.me/http/mux"
	"ztaylor.me/http/router"
	"ztaylor.me/http/websocket"
)

func Routes(rt *api.Runtime) mux.Mux {
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

func WSRoutes(apirt *api.Runtime) http.Handler {
	mux := websocket.NewCache(apirt.Sessions)
	rt := &apiws.Runtime{apirt, mux}
	mux.Route(&websocket.Route{websocket.RouterLit("/connect"), apiws.Connect(rt)})
	mux.Route(&websocket.Route{websocket.RouterLit("/disconnect"), apiws.Disconnect(rt)})
	mux.Route(&websocket.Route{websocket.RouterLit("/ping"), websocket.HandlerFunc(ping)})
	mux.Route(&websocket.Route{websocket.RouterLit("/signup"), apiws.Signup(rt)})
	mux.Route(&websocket.Route{websocket.RouterLit("/login"), apiws.Login(rt)})
	mux.Route(&websocket.Route{websocket.RouterLit("/email"), apiws.Email(rt)})
	mux.Route(&websocket.Route{websocket.RouterLit("/password"), apiws.Password(rt)})
	mux.Route(&websocket.Route{websocket.RouterLit("/logout"), apiws.Logout(rt)})
	mux.Route(&websocket.Route{websocket.RouterLit("/chat"), apiws.Chat(rt)})
	mux.Route(&websocket.Route{websocket.RouterLit("/chat/join"), apiws.ChatJoin(rt)})
	mux.Route(&websocket.Route{websocket.RouterLit("/game"), apiws.Game(rt)})
	mux.Route(&websocket.Route{websocket.RouterLit("/game/new"), apiws.GameNew(rt)})
	mux.Route(&websocket.Route{websocket.RouterLit("/packs/buy"), apiws.PacksBuy(rt)})

	// route 404
	mux.Route(&websocket.Route{
		Router: websocket.RouterFunc(func(*websocket.Message) bool {
			return true
		}),
		Handler: websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
			wsroutefailed(apirt, socket, m)
		}),
	})
	return websocket.UpgradeHandler(mux)
}

// ping does nothing
func ping(*websocket.T, *websocket.Message) {
}

func wsroutefailed(rt *api.Runtime, socket *websocket.T, m *websocket.Message) {
	rt.Root.Logger.New().With(cast.JSON{
		"Session": socket.Session,
		"URI":     m.URI,
	}).Source().Warn("routing failed")
}
