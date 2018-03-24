package cards

import (
	"elemen7s.com"
	"ztaylor.me/js"
)

type Card struct {
	Id    int
	Image string
	vii.CardType
	Costs vii.ElementMap
	Powers
	*Body
}

func NewCard() *Card {
	return &Card{
		Costs:  vii.ElementMap{},
		Powers: NewPowers(),
	}
}

func JsonWithText(card *Card, text *vii.CardText) js.Object {
	json := js.Object{
		"id":          card.Id,
		"image":       card.Image,
		"name":        text.Name,
		"type":        card.CardType.String(),
		"description": text.Description,
		"powers":      card.Powers.JsonWithText(text),
		"flavor":      text.Flavor,
		"costs":       card.Costs.Copy(),
		"body":        card.Body.Json(),
	}

	return json
}
