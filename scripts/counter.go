package scripts

import (
	"errors"

	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/checktarget"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/phase"
	"github.com/zachtaylor/7elements/game/seat"
)

const CounterID = "counter"

func init() {
	script.Scripts[CounterID] = Counter
}

func Counter(game *game.T, seat *seat.T, me interface{}, args []string) (rs []game.Phaser, err error) {
	if !checktarget.IsCard(me) {
		err = ErrMeCard
	} else if len(args) < 1 {
		err = ErrNoTarget
	} else if state, e := checktarget.PlayState(game, seat, args[0]); e != nil {
		err = e
	} else if play, _ := state.Phase.(*phase.Play); play == nil {
		err = ErrBadTarget
	} else if play.IsCancelled {
		err = errors.New("already cancelled")
	} else {
		play.IsCancelled = true
	}
	return
}
