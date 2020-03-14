package scripts

import (
	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/element"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/trigger"
	"github.com/zachtaylor/7elements/power"
)

func init() {
	game.Scripts["call-banners"] = CallTheBanners
}

var ctbCard = &vii.Card{
	Type:  card.BodyType,
	Image: "/img/card/4.jpg",
	Name:  "Bannerman",
	Text:  "At your call",
	Costs: element.Count{},
	Body: &vii.Body{
		Attack: 2,
		Health: 2,
	},
	Powers: power.NewSet(),
}

func CallTheBanners(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	card := game.NewCard(ctbCard)
	card.Username = s.Username
	for i := 0; i < 3; i++ {
		_, e := trigger.Spawn(g, s, card)
		events = append(events, e...)
	}
	return
}
