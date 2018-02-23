package games

import (
	"errors"
	"ztaylor.me/events"
	"ztaylor.me/http"
	"ztaylor.me/js"
)

type playerAgent struct {
	*http.Socket
}

var playerAgents = make(map[string]*Seat)

func init() {
	events.On(http.EVTsocket_close, func(args ...interface{}) {
		socket := args[0].(*http.Socket)
		if seat := playerAgents[socket.Name()]; seat != nil {
			seat.Player = nil
			delete(playerAgents, socket.Name())
		}
	})
}

func ConnectPlayerAgent(seat *Seat, a http.Agent) error {
	if a.Name()[:5] != "ws://" {
		return errors.New("games: agent unsupported")
	} else if socket, ok := a.(*http.Socket); !ok {
		return errors.New("games: agent unsupported type")
	} else {
		seat.Player = &playerAgent{socket}
		playerAgents[socket.Name()] = seat
		return nil
	}
}

func (a playerAgent) Send(uri string, json js.Object) {
	a.Socket.WriteJson(js.Object{
		"uri":  uri,
		"data": json,
	})
}