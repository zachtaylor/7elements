package engine

import (
	"github.com/zachtaylor/7elements/game/seat"
)

func randomfirstturn(seats *seat.List) string { return seats.Keys()[0] }
