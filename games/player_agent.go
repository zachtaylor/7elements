package games

// import (
// 	"github.com/zachtaylor/7elements"
// 	"errors"
// 	"ztaylor.me/events"
// 	"ztaylor.me/http"
// 	"ztaylor.me/js"
// )

// type playerAgent struct {
// 	*http.Socket
// }

// var playerAgents = make(map[string]*vii.GameSeat)

// func init() {
// 	events.On(http.EVTsocket_close, func(args ...interface{}) {
// 		socket := args[0].(*http.Socket)
// 		if seat := playerAgents[socket.Name()]; seat != nil {
// 			seat.Receiver = nil
// 			delete(playerAgents, socket.Name())
// 		}
// 	})
// }

// func ConnectPlayerAgent(seat *vii.GameSeat, a http.Agent) error {
// 	if a.Name()[:5] != "ws://" {
// 		return errors.New("games: agent unsupported")
// 	} else if socket, ok := a.(*http.Socket); !ok {
// 		return errors.New("games: agent unsupported type")
// 	} else {
// 		seat.Receiver = &playerAgent{socket}
// 		playerAgents[socket.Name()] = seat
// 		return nil
// 	}
// }

// func (a playerAgent) Send(uri string, json js.Object) {
// 	a.Socket.WriteJson(js.Object{
// 		"uri":  uri,
// 		"data": json,
// 	})
// }
