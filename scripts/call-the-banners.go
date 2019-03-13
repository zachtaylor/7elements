package scripts

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/engine"
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
	Body: &vii.Body{
		Attack: 2,
		Health: 2,
	},
	Powers: vii.NewPowers(),
}

func CallTheBanners(g *game.T, seat *game.Seat, target interface{}) game.Event {
	for i := 0; i < 3; i++ {
		card := game.NewCard(ctbCard)
		card.Username = seat.Username
		card.IsToken = true
		g.RegisterCard(card)
	}
	animate.GameSeat(g, seat)
	return nil
}
