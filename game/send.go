package game

import (
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

// SendData populates game data in multiple writes to stay under ws frame limit
func (game *T) SendData(username string) {
	seat := game.Seats.Get(username)
	if seat == nil {
		return
	}

	seat.Writer.Write(websocket.NewMessage("/game", game.Data(seat)).EncodeToJSON())
	seat.Writer.Write(wsout.GameState(game.State.MessageData()).EncodeToJSON())

	seat.Writer.Write(wsout.GameHand(seat.Hand.Keys()).EncodeToJSON())
	for _, c := range seat.Hand {
		seat.Writer.Write(wsout.GameCardJSON(c.Data()))
	}

	for _, name := range game.Seats.Keys() {
		s := game.Seats.Get(name)
		seat.Writer.Write(wsout.GameSeatJSON(s.Data()))
		for _, t := range s.Present {
			seat.Writer.Write(wsout.GameTokenJSON(t.Data()))
		}
		for _, c := range s.Past {
			seat.Writer.Write(wsout.GameCardJSON(c.Data()))
		}
	}
}
