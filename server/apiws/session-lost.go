package apiws

import "ztaylor.me/http/websocket"

func _connectSessionWaiter(rt *Runtime, socket *websocket.T) {
	for socketDone, sessionDone := socket.DoneChan(), socket.Session.Done(); ; {
		select {
		case <-socketDone:
		case <-sessionDone:
			rt.Runtime.Root.Logger.New().Add("Socket", socket).Source().Debug()
			socket.Message("/myaccount", nil)
			socket.Session = nil
			return
		}
	}
}
