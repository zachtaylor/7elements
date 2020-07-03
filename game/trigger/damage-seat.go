package trigger

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/out"
	"ztaylor.me/cast"
)

func DamageSeat(g *game.T, card *card.T, seat *game.Seat, n int) []game.Stater {
	if n >= seat.Life {
		return []game.Stater{
			g.Settings.Engine.End(card.Username, seat.Username),
		}
	}
	seat.Life -= n
	g.Settings.Chat.AddMessage(chat.NewMessage(card.Proto.Name, cast.StringI(n)+" damage to "+seat.Username))
	out.GameSeat(g, seat.JSON())
	return nil
}
