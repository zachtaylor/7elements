package server

import (
	"github.com/zachtaylor/7elements/server/apihttp"
	"github.com/zachtaylor/7elements/server/apiws"
	"github.com/zachtaylor/7elements/server/internal"
	"taylz.io/http"
	"taylz.io/http/handler"
	"taylz.io/http/router"
	"taylz.io/http/websocket"
)

func GetFork(server internal.Server) *http.Fork {
	fork := &http.Fork{}

	for _, route := range Routes() {
		fork.Path(route.Router, route.Provider(server))
	}

	return fork
}

func GetWSFork(server internal.Server) *websocket.Fork {
	fork := &websocket.Fork{}

	for _, route := range WSRoutes() {
		fork.Path(route.Router, route.Provider(server))
	}

	return fork
}

// Route builds the actual http.Handler
type Route struct {
	Router   http.Router
	Provider func(internal.Server) http.Handler
}

func NewRoute(r http.Router, provider func(internal.Server) http.Handler) Route {
	return Route{
		Router:   r,
		Provider: provider,
	}
}

// WSRoute builds the actual websocket.Handler
type WSRoute struct {
	Router   websocket.MessageRouter
	Provider func(internal.Server) websocket.MessageHandler
}

func NewWSRoute(path string, provider func(internal.Server) websocket.MessageHandler) WSRoute {
	return WSRoute{
		Router:   websocket.RouterURI(path),
		Provider: provider,
	}
}

// Routes returns the default routing configuration
func Routes() []Route {
	return []Route{
		NewRoute(router.Path(`/api/global.json`), apihttp.GlobalDataHandler),
		NewRoute(router.Path(`/api/myaccount.json`), apihttp.MyAccountHandler),
		// NewRoute(router.Path(`/api/newgame.json`), apihttp.NewGameHandler),
		NewRoute(router.Path(`/api/login`), apihttp.LoginHandler),
		NewRoute(router.Path(`/api/signup`), apihttp.SignupHandler),
		NewRoute(router.Path(`/api/username`), apihttp.UsernameHandler),
		NewRoute(router.Path(`/api/logout`), apihttp.LogoutHandler),
		NewRoute(router.Path(`/api/websocket`), func(server internal.Server) http.Handler {
			return server.GetWebsocketManager().NewUpgrader()
		}),
		NewRoute(router.SinglePage, func(server internal.Server) http.Handler { // SPA override
			return handler.Index(server.GetFileSystem())
		}),
		NewRoute(router.Yes(), func(server internal.Server) http.Handler { // assets
			return http.FileServer(server.GetFileSystem())
		}),
	}
}

// WSRoutes returns the websocket api
func WSRoutes() []WSRoute {

	return []WSRoute{
		// NewWSRoute(`/ping`, func(internal.Server) websocket.MessageHandler {
		// 	return websocket.MessageHandlerFunc(func(*websocket.T, *websocket.Message) {})
		// }),
		NewWSRoute("/signup", apiws.Signup),
		NewWSRoute("/login", apiws.Login),
		NewWSRoute("/email", apiws.Email),
		NewWSRoute("/password", apiws.Password),
		NewWSRoute("/logout", apiws.Logout),
		// NewWSRoute("/chat", apiws.Chat),
		// NewWSRoute("/chat/join", apiws.ChatJoin),
		NewWSRoute("/deck", apiws.UpdateDeck),
		NewWSRoute("/game", apiws.Game),
		NewWSRoute("/game/new", apiws.GameNew),
		NewWSRoute("/game/cancel", apiws.GameCancel),
		NewWSRoute("/packs/buy", apiws.PacksBuy),
		WSRoute{ // route 404
			Router: websocket.MessageRouterYes(),
			Provider: func(server internal.Server) websocket.MessageHandler {
				return websocket.MessageHandlerFunc(func(socket *websocket.T, m *websocket.Message) {
					server.Log().With(map[string]any{
						"URI":       m.URI,
						"SocketID":  socket.ID(),
						"SessionID": socket.SessionID(),
					}).Warn("route unknown")
				})
			},
		},
	}
}
