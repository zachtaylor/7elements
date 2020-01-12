package scripts

import (
	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/game/trigger"

	"github.com/zachtaylor/7elements/game"
)

func init() {
	game.Scripts["call-banners"] = CallTheBanners
}

var ctbCard = &vii.Card{
	Type:  vii.CTYPbody,
	Image: "/img/card/4.jpg",
	Name:  "Bannerman",
	Text:  "At your call",
	Costs: vii.ElementMap{},
	Body: &vii.Body{
		Attack: 2,
		Health: 2,
	},
	Powers: vii.NewPowers(),
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
