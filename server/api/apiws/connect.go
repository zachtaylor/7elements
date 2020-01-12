package apiws

import "ztaylor.me/http/websocket"

func Connect(rt *Runtime) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, _ *websocket.Message) {
		rt.Runtime.Root.Logger.New().Add("Socket", socket).Source().Info()
		rt.Runtime.Ping.Add()
		go ping(rt)
		connect(rt, socket)
	})
}

func connect(rt *Runtime, socket *websocket.T) {
	if socket.Session == nil {
		rt.Runtime.Root.Logger.New().Source().Debug("no session")
		return
	}
	socket.Message("/myaccount", rt.Runtime.Root.AccountJSON(socket.Session.Name()))
	go _connectSessionWaiter(rt, socket)
	connectgame(rt, socket)
}
