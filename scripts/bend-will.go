package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/checktarget"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/phase"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/wsout"
)

const bendwillID = "bend-will"

func init() {
	script.Scripts[bendwillID] = BendWill
}

func BendWill(game *game.T, seat *seat.T, me interface{}, args []string) (rs []game.Phaser, err error) {
	if !checktarget.IsCard(me) {
		err = ErrMeCard
	} else if len(args) < 1 {
		err = ErrNoTarget
	} else if token, e := checktarget.PresentBeing(game, seat, args[0]); err != nil {
		err = e
	} else {
		rs = append(rs, phase.NewChoice(
			seat.Username,
			token.Card.Proto.Name+"-> Awake or Asleep?",
			map[string]interface{}{
				"token": token.Data(),
			},
			[]map[string]interface{}{
				map[string]interface{}{"choice": "awake", "display": `<a>Awake</a>`},
				map[string]interface{}{"choice": "asleep", "display": `<a>Asleep</a>`},
			},
			func(val interface{}) {
				if val == "awake" {
					token.IsAwake = true
				} else if val == "asleep" {
					token.IsAwake = false
				} else {
					return
				}
				game.Seats.WriteSync(wsout.GameToken(token.Data()).EncodeToJSON())
			},
		))
	}
	return
}
