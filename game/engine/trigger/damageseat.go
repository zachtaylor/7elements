package trigger

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

func DamageSeat(game *game.T, seat *seat.T, n int) (rs []game.Phaser) {
	game.Log().With(websocket.MsgData{
		"Target": seat.Username,
		"Life":   seat.Life,
		"Damage": n,
	}).Trace(n)
	seat.Life -= n
	game.Seats.Write(wsout.GameSeatJSON(seat.Data()))
	if seat.Life < 1 {
		rs = append(rs, game.Engine().NewEnding(game, game.NewWinLoss(game.Seats.GetOpponent(seat.Username).Username, seat.Username)))
		return
	}
	// game.Chat(card.Proto.Name, strconv.FormatInt(int64(n), 10)+" damage to "+seat.Username)
	// if rs == nil {
	// 	rs = newphasers()
	// }
	return
}
