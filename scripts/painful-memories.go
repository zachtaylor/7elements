package scripts

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/checktarget"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/engine/trigger"
	"github.com/zachtaylor/7elements/game/seat"
)

const PainfulMemoriesID = "painful-memories"

func init() {
	script.Scripts[PainfulMemoriesID] = PainfulMemories
}

func PainfulMemories(game *game.T, seat *seat.T, me interface{}, args []string) (rs []game.Phaser, err error) {
	if !checktarget.IsCard(me) {
		err = ErrMeCard
	} else if len(args) < 1 {
		err = ErrNoTarget
	} else {
		count := 0
		for _, c := range seat.Past {
			if c.Proto.Type == card.BodyType {
				count++
			}
		}
		for _, name := range game.Seats.Keys() {
			if name == seat.Username {
				continue
			}
			rs = append(rs, trigger.DamageSeat(game, seat, count)...)
		}
	}
	return
}
