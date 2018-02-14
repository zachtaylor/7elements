package games

import (
	"github.com/cznic/mathutil"
)

var gameIdGen, _ = mathutil.NewFC32(0, 999999999, true)

func NewGameId() int {
	return int(gameIdGen.Next())
}
