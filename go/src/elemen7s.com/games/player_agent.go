package games

import (
	"errors"
	"ztaylor.me/http"
	"ztaylor.me/js"
)

type playerAgent struct {
	*http.Socket
}

func ConnectPlayerAgent(seat *Seat, a http.Agent) error {
	if a.Name()[:5] != "ws://" {
		return errors.New("games: agent unsupported")
	} else if socket, ok := a.(*http.Socket); !ok {
		return errors.New("games: agent unsupported type")
	} else {
		socket.On(http.EVTsocket_close, func(args ...interface{}) {
			seat.Player = nil
		})
		seat.Player = &playerAgent{socket}
		return nil
	}
}

func (a playerAgent) Send(uri string, json js.Object) {
	a.Socket.WriteJson(js.Object{
		"uri":  uri,
		"data": json,
	})
}
