package scripts

import (
	"reflect"
	"strconv"

	"github.com/zachtaylor/7elements/element"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/checktarget"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/phase"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/wsout"
)

func init() {
	script.Scripts["new-element"] = NewElement
}

func NewElement(game *game.T, seat *seat.T, me interface{}, args []string) (rs []game.Phaser, err error) {
	if !checktarget.IsCard(me) {
		err = ErrMeCard
	} else if len(args) < 1 {
		err = ErrNoTarget
	} else if meAsCard := game.GetCard(args[0]); meAsCard == nil {
		err = ErrBadTarget
	} else {
		rs = append(rs, phase.NewChoice(
			seat.Username,
			"Create a New Element",
			map[string]interface{}{
				"card": meAsCard.Data(),
			},
			wsout.GameChoiceElementsData,
			func(val interface{}) {
				log := game.Log().Add("val", val)
				if v, _ := val.(string); len(v) < 1 {
					log.Add("type", reflect.TypeOf(val)).Warn()
					seat.Writer.Write(wsout.ErrorJSON("vii", "elementid failed"))
				} else if i, _ := strconv.ParseInt(v, 10, 0); i < 1 || i > 7 {
					seat.Writer.Write(wsout.ErrorJSON("vii", "elementid out of bounds"))
				} else {
					seat.Karma.Append(element.T(i), false)
					game.Seats.Write(wsout.GameSeatJSON(seat.Data()))
				}
			},
		))
	}
	return
}
