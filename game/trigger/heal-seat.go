package trigger

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/update"
	"ztaylor.me/cast"
)

func HealSeat(g *game.T, card *game.Card, seat *game.Seat, n int) []game.Stater {
	seat.Life += n
	update.GameChat(g, seat.Username, "gain "+cast.StringI(n)+" Life ("+card.Proto.Name+")")
	update.Seat(g, seat)
	return nil
}
