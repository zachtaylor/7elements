package scripts

import (
	"elemen7s.com"
	"elemen7s.com/engine"
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
