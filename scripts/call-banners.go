package scripts

import (
	"strconv"

	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/element"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/checktarget"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/engine/trigger"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/game/token"
	"github.com/zachtaylor/7elements/power"
)

func init() {
	script.Scripts["call-banners"] = CallTheBanners
}

var callthebannersTokenCardProto = &card.Prototype{
	Type:  card.BodyType,
	Name:  "Bannerman",
	Text:  "At your call",
	Costs: element.Count{},
	Body: &card.Body{
		Attack: 2,
		Health: 2,
	},
	Powers: power.NewSet(),
}

func CallTheBanners(game *game.T, seat *seat.T, me interface{}, args []string) (rs []game.Phaser, err error) {
	if !checktarget.IsCard(me) {
		err = ErrMeCard
	} else {
		card := card.New(callthebannersTokenCardProto, seat.Username)
		for i := 0; i < 3; i++ {
			token := token.New(card, seat.Username)
			token.Image = "17." + strconv.FormatInt(int64(i), 10)
			if _rs := trigger.NewToken(game, seat, token); len(_rs) > 1 {
				rs = append(rs, _rs...)
			}
		}
	}
	return
}
