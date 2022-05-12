package request

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/phase"
	"github.com/zachtaylor/7elements/game/seat"
)

func Phase(game *game.T, seat *seat.T, json map[string]any) {
	phase.TryOnRequest(game, seat, json)
}
