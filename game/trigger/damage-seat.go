package trigger

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/event/end"
)

func DamageSeat(g *game.T, card *game.Card, seat *game.Seat, n int) []game.Event {
	if n > seat.Life {
		return []game.Event{end.New(card.Username, seat.Username)}
	}
	seat.Life -= n
	g.SendSeatUpdate(seat)
	return nil
}
