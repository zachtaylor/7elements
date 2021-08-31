package checktarget

import (
	"errors"

	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/seat"
)

func PlayState(game *game.T, seat *seat.T, stateid string) (state *game.State, err error) {
	for s := game.State.Stack; s != nil; s = s.Stack {
		if s.ID() == stateid && s.Phase.Name() == "play" {
			state = s
			break
		}
	}
	if state == nil {
		err = errors.New("state not found")
	}
	return
}
