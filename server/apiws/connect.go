package apiws

import (
	"github.com/zachtaylor/7elements/runtime"
	"ztaylor.me/http/websocket"
)

func Connect(rt *runtime.T) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, _ *websocket.Message) {
		rt.Logger.New().Add("Socket", socket).Info()
		connect(rt, socket)
		go rt.Ping()
	})
}

func connect(rt *runtime.T, socket *websocket.T) {
	if socket.Session == nil {
		rt.Logger.New().Debug("no session")
		return
	}

	player := rt.Players.Get(socket.Session.Name())
	if player == nil {
		rt.Log().Error("session with no player")
		return
	}

	socket.Send("/myaccount", rt.AccountJSON(player.Account))
	go connectWaiter(rt, socket)
	connectgame(rt, socket)
}

func connectWaiter(rt *runtime.T, socket *websocket.T) {
	for socketDone, sessionDone := socket.DoneChan(), socket.Session.Done(); ; {
		log := rt.Logger.New().Add("Socket", socket)
		select {
		case <-socketDone:
			log.Debug("done")
			return
		case <-sessionDone:
			log.Warn("session")
			socket.Send("/myaccount", nil)
			socket.Session = nil
			return
		}
	}
}
