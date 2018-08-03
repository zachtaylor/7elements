package api

import (
	"github.com/zachtaylor/7tcg"
	"ztaylor.me/events"
	"ztaylor.me/http"
	"ztaylor.me/log"
)

var playerAgents = make(map[string]*vii.GameSeat)

func init() {
	events.On(http.EVTsocket_close, func(args ...interface{}) {
		socket := args[0].(*http.Socket)
		if seat := playerAgents[socket.Name()]; seat != nil {
			seat.Receiver = nil
			delete(playerAgents, socket.Name())
		}
	})
}

func JoinHandler(r *http.Request) error {
	if gameid := r.Data.Sval("gameid"); gameid == "" {
		return ErrGameIdRequired
	} else if game := vii.GameService.Get(gameid); game == nil {
		return ErrGameMissing
	} else if seat := game.GetSeat(r.Username); seat == nil {
		log.WithFields(log.Fields{
			"Game":     game,
			"Username": r.Username,
		}).Warn("/api/join: not participating in game")
	} else if seat.Receiver != nil {
		log.WithFields(log.Fields{
			"Game": game,
			"Seat": seat,
		}).Warn("/api/join: seat already occupied")
	} else if socket, ok := r.Agent.(*http.Socket); !ok {
		log.WithFields(log.Fields{
			"Game":  game,
			"Seat":  seat,
			"Agent": r.Agent,
		}).Warn("/api/join: request agent is not socket")
	} else {
		playerAgents[socket.Name()] = seat
		seat.Receiver = &SocketReceiver{socket}
		go game.SendCatchup(seat.Username)
	}
	return nil
}
