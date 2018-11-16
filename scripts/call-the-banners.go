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
	Type:  vii.CTYPbody,
	Image: "/img/cards/zealot-0.jpg",
	Name:  "Bannerman",
	Text:  "At your call",
	Costs: vii.ElementMap{},
	Body: &vii.CardBody{
		Attack: 2,
		Health: 2,
	},
	Powers: vii.NewPowers(),
}

func CallTheBanners(game *vii.Game, seat *vii.GameSeat, target interface{}) vii.GameEvent {
	for i := 0; i < 3; i++ {
		card := vii.NewGameCard(ctbCard)
		card.Username = seat.Username
		card.IsToken = true
		game.RegisterCard(card)
	}
	animate.GameSeat(game, seat)
	return nil
}
