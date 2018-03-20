package scripts

import (
	"elemen7s.com"
	"elemen7s.com/cards"
	"elemen7s.com/cards/texts"
	"elemen7s.com/cards/types"
	"elemen7s.com/games"
)

func init() {
	games.Scripts["call-the-banners"] = CallTheBanners
}

var ctbCard = &cards.Card{
	Image:    "/img/cards/zealot-0.jpg",
	CardType: ctypes.Body,
	Costs:    vii.ElementMap{},
	Powers:   cards.NewPowers(),
	Body: &cards.Body{
		Attack: 2,
		Health: 2,
	},
}

var ctbCardText = &texts.Text{
	Language: "en-US",
	Name:     "Bannerman",
	Powers:   make(map[int]string),
	Flavor:   "At your call",
}

func CallTheBanners(g *games.Game, s *games.Seat) {
	for i := 0; i < 3; i++ {
		g.RegisterToken(s.Username, games.NewCard(ctbCard, ctbCardText))
	}
}
