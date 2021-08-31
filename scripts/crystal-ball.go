package scripts

import (
	"reflect"

	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/checktarget"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/phase"
	"github.com/zachtaylor/7elements/game/seat"
)

const crystalballID = "crystal-ball"

func init() {
	script.Scripts[crystalballID] = CrystalBall
}

func CrystalBall(game *game.T, seat *seat.T, me interface{}, args []string) (rs []game.Phaser, err error) {
	if !checktarget.IsToken(me) {
		err = ErrMeToken
	} else if len(args) < 1 {
		err = ErrNoTarget
	} else {
		card := seat.Deck.Cards[0]
		rs = append(rs, phase.NewChoice(
			seat.Username,
			"Shuffle your Future?",
			map[string]interface{}{
				"card": card.Data(),
			},
			[]map[string]interface{}{
				map[string]interface{}{"choice": true, "display": `<a>Yes</a>`},
				map[string]interface{}{"choice": false, "display": `<a>No</a>`},
			},
			func(val interface{}) {
				game.Log().Add("val", val).Add("type", reflect.TypeOf(val)).Info()
				if choice, ok := val.(bool); ok {
					if choice {
						seat.Deck.Shuffle()
					}
				}
			},
		))
	}
	return
}
