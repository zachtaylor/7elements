package scripts

import (
	"elemen7s.com/games"
)

func init() {
	games.Scripts["7-elements"] = Elemen7s
}

func Elemen7s(g *games.Game, s *games.Seat, target interface{}) {
	g.Win(s)
}
