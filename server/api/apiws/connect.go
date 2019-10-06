package apiws

import (
	"github.com/zachtaylor/7elements/server/api"
	"ztaylor.me/cast"
	"ztaylor.me/http/websocket"
)

func Connect(rt *api.Runtime) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := rt.Root.Logger.New().Tag("apiws/connect")
		pushJSON(socket, "/data/ping", cast.JSON{
			"online": rt.Sessions.Count(),
		})
		if socket.Session == nil {
			log.Debug("no session")
			return
		}
		connectAccount(rt, log, socket)
		connectGame(rt, log, socket)
	})
}
