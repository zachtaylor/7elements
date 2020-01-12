package apiws

import "ztaylor.me/http/websocket"

func Disconnect(rt *Runtime) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, _ *websocket.Message) {
		rt.Runtime.Root.Logger.New().Add("Socket", socket).Source().Info()
		rt.Runtime.Ping.Remove()
		go ping(rt)
	})
}
