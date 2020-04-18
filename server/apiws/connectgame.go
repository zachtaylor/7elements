package apiws

import "ztaylor.me/http/websocket"

func connectgame(rt *Runtime, socket *websocket.T) {
	log := rt.Runtime.Root.Logger.New().Add("Session", socket.Session)
	if socket.Session == nil {
		log.Source().Warn("no session")
	} else if g := rt.Runtime.Games.FindUsername(socket.Session.Name()); g == nil {
		// log.Source().Warn("no game")
	} else if seat := g.GetSeat(socket.Session.Name()); seat == nil || seat.Receiver != nil {
		log.Source().Warn("no game")
	} else {
		log.Add("GameID", g.ID()).Source().Debug()
		seat.Receiver = &WebsocketReceiver{socket}
		chat := g.GetChat()
		go g.Request(seat.Username, "connect", nil)
		chat.AddUser(socket)
		go _connectGameWaiter(socket, g, seat.Username, log)
	}
}
