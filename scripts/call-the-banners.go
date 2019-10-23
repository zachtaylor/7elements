package scripts

import (
	vii "github.com/zachtaylor/7elements"

	"github.com/zachtaylor/7elements/game"
)

func init() {
	game.Scripts["call-the-banners"] = CallTheBanners
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

func CallTheBanners(g *game.T, seat *game.Seat, target interface{}) []game.Event {
	for i := 0; i < 3; i++ {
		card := game.NewCard(ctbCard)
		card.Username = seat.Username
		g.RegisterCard(card)
	}
	g.SendSeatUpdate(seat)
	return nil
}
