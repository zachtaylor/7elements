package scripts

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/engine"
)

func init() {
	engine.Scripts["call-the-banners"] = CallTheBanners
}

var ctbCard = &vii.Card{
	Image:    "/img/cards/zealot-0.jpg",
	CardType: vii.CTYPbody,
	Costs:    vii.ElementMap{},
	Powers:   vii.NewPowers(),
	CardBody: &vii.CardBody{
		Attack: 2,
		Health: 2,
	},
}

var ctbCardText = &vii.CardText{
	Name:   "Bannerman",
	Powers: make(map[int]string),
	Flavor: "At your call",
}

func CallTheBanners(game *vii.Game, t *engine.Timeline, seat *vii.GameSeat, target interface{}) *engine.Timeline {
	for i := 0; i < 3; i++ {
		card := vii.NewGameCard(ctbCard, ctbCardText)
		card.Username = seat.Username
		card.IsToken = true
		game.RegisterCard(card)
	}
	animate.BroadcastSeatUpdate(game, seat)
	return nil
}
