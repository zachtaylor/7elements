package trigger

import (
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/event/end"
	"ztaylor.me/cast"
)

func DamageSeat(g *game.T, card *game.Card, seat *game.Seat, n int) []game.Event {
	if n >= seat.Life {
		return []game.Event{end.New(card.Username, seat.Username)}
	}
	seat.Life -= n
	go g.GetChat().AddMessage(chat.NewMessage(card.Card.Name, cast.StringI(n)+" damage to "+seat.Username))
	g.SendSeatUpdate(seat)
	return nil
}
