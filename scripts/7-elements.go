package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/event"
)

func init() {
	game.Scripts["7-elements"] = Elemen7s
}

func Elemen7s(g *game.T, seat *game.Seat, target interface{}) []game.Event {
	return []game.Event{event.NewEndEvent(seat.Username, g.GetOpponentSeat(seat.Username).Username)}
}
