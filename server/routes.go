package server

import (
	"github.com/zachtaylor/7elements/server/apihttp"
	"github.com/zachtaylor/7elements/server/apiws"
	"github.com/zachtaylor/7elements/server/internal"
	"taylz.io/http"
	"taylz.io/http/router"
	"taylz.io/http/websocket"
)

// Routes returns the default routing configuration
func Routes(wsUpgrader http.Handler) []internal.RouteBuilder {
	return []internal.RouteBuilder{
		internal.NewRouteBuilder(router.Path(`/global.json`), apihttp.GlobalDataHandler),
		internal.NewRouteBuilder(router.Path(`/myaccount.json`), apihttp.MyAccountHandler),
		// internal.NewRouteBuilder(router.Path(`/newgame.json`), apihttp.NewGameHandler),
		internal.NewRouteBuilder(router.Path(`/login`), apihttp.LoginHandler),
		internal.NewRouteBuilder(router.Path(`/signup`), apihttp.SignupHandler),
		internal.NewRouteBuilder(router.Path(`/username`), apihttp.UsernameHandler),
		internal.NewRouteBuilder(router.Path(`/logout`), apihttp.LogoutHandler),
		internal.NewRouteBuilder(router.Path(`/websocket`), func(server internal.Server) http.Handler {
			return wsUpgrader
		}),
		{ // 404
			Router: router.Yes(),
			Provider: func(internal.Server) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(404)
				})
			},
		},
		// internal.NewRouteBuilder(router.SinglePage, func(server internal.Server) http.Handler { // SPA override
		// 	return handler.Index(server.FileSystem())
		// }),
		// internal.NewRouteBuilder(router.Yes(), func(server internal.Server) http.Handler { // assets
		// 	return http.FileServer(server.FileSystem())
		// }),
	}
}

// WSRoutes returns the websocket api
func WSRoutes() []internal.WSRouteBuilder {

	return []internal.WSRouteBuilder{
		// NewWSRoute(`/ping`, func(internal.Server) websocket.MessageHandler {
		// 	return websocket.MessageHandlerFunc(func(*websocket.T, *websocket.Message) {})
		// }),
		internal.NewWSRouteBuilder("/signup", apiws.Signup),
		internal.NewWSRouteBuilder("/login", apiws.Login),
		internal.NewWSRouteBuilder("/email", apiws.Email),
		internal.NewWSRouteBuilder("/password", apiws.Password),
		internal.NewWSRouteBuilder("/logout", apiws.Logout),
		// NewWSRoute("/chat", apiws.Chat),
		// NewWSRoute("/chat/join", apiws.ChatJoin),
		internal.NewWSRouteBuilder("/deck", apiws.UpdateDeck),
		internal.NewWSRouteBuilder("/game", apiws.Game),
		internal.NewWSRouteBuilder("/game/new", apiws.GameNew),
		internal.NewWSRouteBuilder("/game/cancel", apiws.GameCancel),
		internal.NewWSRouteBuilder("/packs/buy", apiws.PacksBuy),
		{ // route 404
			Router: websocket.MessageRouterYes(),
			Provider: func(server internal.Server) websocket.MessageHandler {
				return websocket.MessageHandlerFunc(func(socket *websocket.T, m *websocket.Message) {
					user := server.Users().GetWebsocket(socket)
					server.Log().With(map[string]any{
						"URI":      m.URI,
						"SocketID": socket.ID(),
						"User":     user,
					}).Warn("route unknown")
				})
			},
		},
	}
}
