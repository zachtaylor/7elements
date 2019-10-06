package game

func TargetPlayerBeing(g *T, seat *Seat, target string) interface{} {
	if target == "self" {
		return seat
	} else if target == "enemy" {
		return g.GetOpponentSeat(seat.Username)
	}
	return TargetBeing(g, target)
}

func TargetBeing(g *T, target interface{}) *Card {
	log := g.Log().Add("Target", target).Tag("game/target-being")
	if gcid, ok := target.(string); !ok {
		log.Error("no gcid")
	} else if card := g.Cards[gcid]; card == nil {
		log.Error("no card")
	} else if s := g.GetSeat(card.Username); s == nil {
		log.Add("Card", card).Error("no seat")
	} else if !s.HasPresentCard(card.Id) {
		log.Add("Card", card).Error("not present")
	} else {
		log.Add("Card", card).Debug()
		return card
	}
	return nil
}
