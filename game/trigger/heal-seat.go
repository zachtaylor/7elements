package trigger

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/out"
	"ztaylor.me/cast"
)

func HealSeat(g *game.T, card *card.T, seat *game.Seat, n int) []game.Stater {
	seat.Life += n
	g.Settings.Chat.AddMessage(chat.NewMessage(seat.Username, "gain "+cast.StringI(n)+" Life ("+card.Proto.Name+")"))
	out.GameSeat(g, seat.JSON())
	return nil
}
