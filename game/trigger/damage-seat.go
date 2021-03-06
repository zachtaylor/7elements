package trigger

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/state/end"
	"github.com/zachtaylor/7elements/game/update"
	"ztaylor.me/cast"
)

func DamageSeat(g *game.T, card *card.T, seat *game.Seat, n int) []game.Stater {
	if n >= seat.Life {
		return []game.Stater{end.New(card.Username, seat.Username)}
	}
	seat.Life -= n
	update.GameChat(g, card.Proto.Name, cast.StringI(n)+" damage to "+seat.Username)
	update.Seat(g, seat)
	return nil
}
