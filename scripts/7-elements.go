package scripts

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/engine"
)

func init() {
	engine.Scripts["7-elements"] = Elemen7s
}

func Elemen7s(game *vii.Game, t *engine.Timeline, seat *vii.GameSeat, target interface{}) *engine.Timeline {
	game.Results = &vii.GameResults{
		Winner: seat.Username,
		Loser:  game.GetOpponentSeat(seat.Username).Username,
	}
	return nil
}
