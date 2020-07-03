package apiws

import "ztaylor.me/http/websocket"

func Logout(rt *Runtime) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		if socket.Session != nil {
			rt.Runtime.Root.Logger.New().Add("Socket", socket).Source().Info()
			socket.Session.Close()
		} else {
			rt.Runtime.Root.Logger.New().Add("Socket", socket).Source().Warn("cookie missing")
		}
		socket.Message("/myaccount", nil)
		rt.SendPing()
	})
}
