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
	rt.Handler.Path(router.Path(`/api/global.json`), apihttp.GlobalDataHandler(rt))
	rt.Handler.Path(router.Path(`/api/myaccount.json`), apihttp.MyAccountHandler(rt))
	// rt.Handler.Path(router.Path(`/api/newgame.json`), apihttp.NewGameHandler(rt))
	rt.Handler.Path(router.Path(`/api/login`), apihttp.LoginHandler(rt))
	rt.Handler.Path(router.Path(`/api/signup`), apihttp.SignupHandler(rt))
	rt.Handler.Path(router.Path(`/api/username`), apihttp.UsernameHandler(rt))
	rt.Handler.Path(router.Path(`/api/logout`), apihttp.LogoutHandler(rt))
	rt.Handler.Path(router.Path(`/api/websocket`), rt.Sockets.NewUpgrader())
	rt.Handler.Path(router.SinglePage, handler.Index(fs))
	fsHandler := http.FileServer(fs)
	// rt.Handler.Path(router.PathStarts(`/img/`), handler.AddPrefix("/assets", fsHandler)) // fix angular asset layout
	rt.Handler.Path(router.Bool(true), fsHandler)
}

func WSRoutes(rt *runtime.T) {
	// rt.WSHandler.Path(websocket.RouterURI("/connect"), apiws.Connect(rt))
	// rt.WSHandler.Path(websocket.RouterURI("/disconnect"), apiws.Disconnect(rt))
	rt.WSHandler.Path(websocket.RouterURI("/ping"), websocket.HandlerFunc(func(*websocket.T, *websocket.Message) {}))
	rt.WSHandler.Path(websocket.RouterURI("/signup"), apiws.Signup(rt))
	rt.WSHandler.Path(websocket.RouterURI("/login"), apiws.Login(rt))
	rt.WSHandler.Path(websocket.RouterURI("/email"), apiws.Email(rt))
	rt.WSHandler.Path(websocket.RouterURI("/password"), apiws.Password(rt))
	rt.WSHandler.Path(websocket.RouterURI("/logout"), apiws.Logout(rt))
	rt.WSHandler.Path(websocket.RouterURI("/chat"), apiws.Chat(rt))
	rt.WSHandler.Path(websocket.RouterURI("/deck"), apiws.UpdateDeck(rt))
	// rt.WSHandler.Path(websocket.RouterURI("/chat/join"), apiws.ChatJoin(rt))
	rt.WSHandler.Path(websocket.RouterURI("/game"), apiws.Game(rt))
	rt.WSHandler.Path(websocket.RouterURI("/game/new"), apiws.GameNew(rt))
	rt.WSHandler.Path(websocket.RouterURI("/game/cancel"), apiws.GameCancel(rt))
	rt.WSHandler.Path(websocket.RouterURI("/packs/buy"), apiws.PacksBuy(rt))

	// route 404
	rt.WSHandler.Path(
		websocket.RouterFunc(func(*websocket.Message) bool { return true }),
		websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
			rt.Logger.With(map[string]interface{}{
				"URI":       m.URI,
				"SocketID":  socket.ID(),
				"SessionID": socket.SessionID(),
			}).Warn("route unknown")
		}),
	)
}
