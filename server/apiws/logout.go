package apiws

import (
	"github.com/zachtaylor/7elements/runtime"
	"ztaylor.me/http/websocket"
)

func Logout(rt *runtime.T) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		if socket.Session != nil {
			rt.Log().Add("Socket", socket).Info()
			socket.Session.Close()
		} else {
			rt.Log().Add("Socket", socket).Warn("cookie missing")
		}
		socket.Send("/myaccount", nil)
		go rt.Ping()
	})
}
