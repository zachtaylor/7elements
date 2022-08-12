package internal

import (
	"taylz.io/http"
	"taylz.io/http/websocket"
)

// RouteBuilder builds the actual http.Handler
type RouteBuilder struct {
	Router   http.Router
	Provider func(Server) http.Handler
}

func NewRouteBuilder(r http.Router, provider func(Server) http.Handler) RouteBuilder {
	return RouteBuilder{
		Router:   r,
		Provider: provider,
	}
}

// WSRouteBuilder builds the actual websocket.MessageHandler
type WSRouteBuilder struct {
	Router   websocket.MessageRouter
	Provider func(Server) websocket.MessageHandler
}

func NewWSRouteBuilder(path string, provider func(Server) websocket.MessageHandler) WSRouteBuilder {
	return WSRouteBuilder{
		Router:   websocket.MessageRouterURI(path),
		Provider: provider,
	}
}
