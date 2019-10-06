package apiws

import (
	"github.com/zachtaylor/7elements/server/api"
	"ztaylor.me/http/websocket"
)

func Logout(rt *api.Runtime) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		pushJSON(socket, "/data/myaccount", nil)
		log := rt.Root.Logger.New().Tag("apiws/logout").Add("Socket", socket.String())
		if socket.Session != nil {
			rt.Sessions.Remove(socket.Session)
			socket.Session = nil
			pushJSON(socket, "/data/ping", api.PingData(rt))
			log.Debug("close")
		} else {
			log.Warn("cookie missing")
		}
	})
}
