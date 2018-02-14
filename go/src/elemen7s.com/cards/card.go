package cards

import (
	"elemen7s.com/cards/texts"
	"elemen7s.com/cards/types"
	"elemen7s.com/elements"
	"ztaylor.me/js"
)

type Card struct {
	Id    int
	Image string
	ctypes.CardType
	Costs elements.Stack
	Powers
	*Body
}

func NewCard() *Card {
	return &Card{
		Costs:  elements.Stack{},
		Powers: NewPowers(),
	}
}

func JsonWithText(card *Card, text *texts.Text) js.Object {
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
