package apiws

import (
	"github.com/zachtaylor/7elements/runtime"
	"ztaylor.me/http/websocket"
)

func connectgame(rt *runtime.T, socket *websocket.T) {
	log := rt.Logger.New().Add("Session", socket.Session)
	if socket.Session == nil {
		log.Warn("no session")
	} else if g := rt.Games.FindUsername(socket.Session.Name()); g == nil {
		// log.Source().Warn("no game")
	} else if seat := g.GetSeat(socket.Session.Name()); seat == nil || seat.Player != nil {
		log.Warn("no game")
	} else {
		log.Add("GameID", g.ID()).Debug()
		seat.Player = &WebsocketReceiver{socket}
		go g.Request(seat.Username, "connect", nil)
		g.Settings.Chat.AddUser(socket)
		go _connectGameWaiter(socket, g, seat.Username, log)
	}
}
