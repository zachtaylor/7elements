package apiws

import "ztaylor.me/http/websocket"

func Connect(rt *Runtime) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, _ *websocket.Message) {
		rt.Runtime.Root.Logger.New().Add("Socket", socket).Source().Info()
		rt.Runtime.Ping.Add()
		connect(rt, socket)
		go rt.SendPing()
	})
}

func connect(rt *Runtime, socket *websocket.T) {
	if socket.Session == nil {
		rt.Runtime.Root.Logger.New().Source().Debug("no session")
		return
	}

	// introduce connection service

	socket.Message("/myaccount", rt.Runtime.Root.FindAccountJSON(socket.Session.Name()))
	go connectWaiter(rt, socket)
	connectgame(rt, socket)
}

func connectWaiter(rt *Runtime, socket *websocket.T) {
	for socketDone, sessionDone := socket.DoneChan(), socket.Session.Done(); ; {
		log := rt.Runtime.Root.Logger.New().Add("Socket", socket)
		select {
		case <-socketDone:
			log.Source().Debug("done")
			return
		case <-sessionDone:
			log.Source().Warn("session")
			socket.Message("/myaccount", nil)
			socket.Session = nil
			return
		}
	}
}
