package scripts

import (
	"reflect"

	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/checktarget"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/phase"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/wsout"
)

const noviceseerId = "novice-seer"

func init() {
	script.Scripts[noviceseerId] = NoviceSeer
}

func NoviceSeer(game *game.T, seat *seat.T, me interface{}, args []string) (rs []game.Phaser, err error) {
	if !checktarget.IsToken(me) {
		err = ErrMeToken
	} else if len(seat.Deck.Cards) < 1 {
		err = ErrNoTarget
	} else if card := seat.Deck.Draw(); card == nil {
		err = ErrBadTarget
	} else {
		rs = append(rs, phase.NewChoice(
			seat.Username,
			"Novice Seer",
			map[string]interface{}{
				"card": card.Data(),
			},
			[]map[string]interface{}{
				map[string]interface{}{
					"choice":  "false",
					"display": "Put on top of your Future",
				},
				map[string]interface{}{
					"choice":  "true",
					"display": "Put into your Past",
				},
			},
			func(val interface{}) {
				game.Log().Add("val", val).Add("type", reflect.TypeOf(val)).Info()
				if destroy, _ := val.(bool); destroy {
					seat.Past[card.ID] = card
				} else {
					seat.Deck.Prepend(card)
				}
				game.Seats.Write(wsout.GameSeatJSON(seat.Data()))
			},
		))
	}

	return
}
