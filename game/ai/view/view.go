package view

import "github.com/zachtaylor/7elements/game"

type T struct {
	Game  *game.G
	Self  *game.Player
	Enemy *game.Player

	State    *game.State
	IsMyMain bool
}

func (t *T) Update(state *game.State) {
	t.State = state
	t.IsMyMain = state.T.Phase.Type() == "main" && state.Player() == t.Self.ID()
}
