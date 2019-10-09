package apiws

import (
	"github.com/zachtaylor/7elements/server/api"
	"ztaylor.me/cast"
	"ztaylor.me/http/websocket"
)

func Connect(rt *api.Runtime) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := rt.Root.Logger.New().Tag("apiws/connect")
		rt.Ping.Add()
		pushPingJSON(rt, socket)
		if socket.Session == nil {
			log.Debug("no session")
			return
		}
		connectAccount(rt, log, socket)
		connectGame(rt, log, socket)
	})
}

func pushPingJSON(rt *api.Runtime, socket *websocket.T) {
	pushJSON(socket, "/data/ping", cast.JSON{
		"ping":   rt.Ping.Get(),
		"online": rt.Sessions.Count(),
	})
}
