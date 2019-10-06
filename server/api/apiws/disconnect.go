package apiws

import (
	"github.com/zachtaylor/7elements/server/api"
	"ztaylor.me/http/websocket"
)

func Disconnect(rt *api.Runtime) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		rt.Root.Logger.New().Add("Username", socket.GetUser()).Tag("apiws/disconnect").Info()
	})
}
