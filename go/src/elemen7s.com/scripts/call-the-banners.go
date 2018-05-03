package scripts

import (
	"elemen7s.com"
	"elemen7s.com/games"
)

func init() {
	games.Scripts["call-the-banners"] = CallTheBanners
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

func CallTheBanners(g *games.Game, s *games.Seat, target interface{}) {
	for i := 0; i < 3; i++ {
		g.RegisterToken(s.Username, vii.NewGameCard(ctbCard, ctbCardText))
	}
	games.BroadcastAnimateSeatUpdate(g, s)
	g.Active.Activate(g)
}
