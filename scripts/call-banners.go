package scripts

import (
	vii "github.com/zachtaylor/7elements"

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

func CallTheBanners(g *game.T, seat *game.Seat, target interface{}) []game.Event {
	for i := 0; i < 3; i++ {
		card := game.NewCard(ctbCard)
		card.Username = seat.Username
		g.RegisterCard(card)
		seat.Present[card.Id] = card
		g.SendCardUpdate(card)
	}
	return nil
}
