package server

import (
	"net/http"

	"github.com/zachtaylor/7elements/server/apihttp"
	"github.com/zachtaylor/7elements/server/apiws"
	"github.com/zachtaylor/7elements/server/runtime"
	"taylz.io/http/handler"
	"taylz.io/http/router"
	"taylz.io/http/websocket"
)

func Routes(rt *runtime.T, fs http.FileSystem) {
	WSRoutes(rt)
	rt.Server.Path(router.Path(`/api/global.json`), apihttp.GlobalDataHandler(rt))
	rt.Server.Path(router.Path(`/api/myaccount.json`), apihttp.MyAccountHandler(rt))
	// rt.Server.Path(router.Path(`/api/newgame.json`), apihttp.NewGameHandler(rt))
	rt.Server.Path(router.Path(`/api/login`), apihttp.LoginHandler(rt))
	rt.Server.Path(router.Path(`/api/signup`), apihttp.SignupHandler(rt))
	rt.Server.Path(router.Path(`/api/username`), apihttp.UsernameHandler(rt))
	rt.Server.Path(router.Path(`/api/logout`), apihttp.LogoutHandler(rt))
	rt.Server.Path(router.Path(`/api/websocket`), rt.Sockets.Upgrader())
	rt.Server.Path(router.SinglePage, handler.Index(fs))
	fsHandler := http.FileServer(fs)
	// rt.Server.Path(router.PathStarts(`/img/`), handler.AddPrefix("/assets", fsHandler)) // fix angular asset layout
	rt.Server.Path(router.Bool(true), fsHandler)
}

func WSRoutes(rt *runtime.T) {
	// rt.WSServer.Path(websocket.RouterURI("/connect"), apiws.Connect(rt))
	// rt.WSServer.Path(websocket.RouterURI("/disconnect"), apiws.Disconnect(rt))
	rt.WSServer.Path(websocket.RouterURI("/ping"), websocket.HandlerFunc(func(*websocket.T, *websocket.Message) {}))
	rt.WSServer.Path(websocket.RouterURI("/signup"), apiws.Signup(rt))
	rt.WSServer.Path(websocket.RouterURI("/login"), apiws.Login(rt))
	rt.WSServer.Path(websocket.RouterURI("/email"), apiws.Email(rt))
	rt.WSServer.Path(websocket.RouterURI("/password"), apiws.Password(rt))
	rt.WSServer.Path(websocket.RouterURI("/logout"), apiws.Logout(rt))
	rt.WSServer.Path(websocket.RouterURI("/chat"), apiws.Chat(rt))
	rt.WSServer.Path(websocket.RouterURI("/deck"), apiws.UpdateDeck(rt))
	rt.WSServer.Path(websocket.RouterURI("/chat/join"), apiws.ChatJoin(rt))
	rt.WSServer.Path(websocket.RouterURI("/game"), apiws.Game(rt))
	rt.WSServer.Path(websocket.RouterURI("/game/new"), apiws.GameNew(rt))
	rt.WSServer.Path(websocket.RouterURI("/packs/buy"), apiws.PacksBuy(rt))

	// route 404
	rt.WSServer.Path(
		websocket.RouterFunc(func(*websocket.Message) bool {
			return true
		}),
		websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
			wsroutefailed(rt, socket, m)
		}),
	)
}

func wsroutefailed(rt *runtime.T, socket *websocket.T, m *websocket.Message) {
	rt.Log().With(websocket.MsgData{
		"User": socket.Name(),
		"URI":  m.URI,
	}).Warn("routing failed")
}
