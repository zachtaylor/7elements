package apiws

import "ztaylor.me/http/websocket"

func Logout(rt *Runtime) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		socket.Message("/myaccount", nil)
		if socket.Session != nil {
			rt.Runtime.Root.Logger.New().Add("Socket", socket).Source().Debug("close")
			rt.Runtime.Sessions.Remove(socket.Session)
			socket.Session = nil
			go rt.SendPing()
		} else {
			rt.Runtime.Root.Logger.New().Add("Socket", socket).Source().Warn("cookie missing")
		}
	})
}
