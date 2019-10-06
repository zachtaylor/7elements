package apiws

import (
	"github.com/zachtaylor/7elements/server/api"
	"ztaylor.me/http/websocket"
	"ztaylor.me/log"
)

func connectGame(rt *api.Runtime, log *log.Entry, socket *websocket.T) {
	if socket.Session == nil {
		log.Debug("no session")
	} else if g := rt.Games.FindUsername(socket.Session.Name()); g == nil {
		log.Debug("no game")
	} else if seat := g.GetSeat(socket.Session.Name()); seat == nil || seat.Receiver != nil {
		log.Add("Seat", seat).Warn("seat?")
	} else {
		log.Add("Name", socket.Session.Name()).Add("GameID", g.ID()).Debug("game data")
		pushJSON(socket, "/game", g.PerspectiveJSON(socket.Session.Name()))
		seat.Receiver = &WebsocketReceiver{socket}
		go g.Request(seat.Username, "connect", nil)
		go g.GetChat().User(newChatUser("game#"+g.ID(), socket))
		select {
		case <-g.Done():
		case <-socket.Done():
		}
		log.Debug("vacate game")
		seat.Receiver = nil
	}
}
