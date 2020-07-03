package engine

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/state"
)

func NewStart(seat string) game.Stater {
	return state.NewStart(seat)
}
