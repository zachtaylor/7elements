package cards

import (
	"7elements.ztaylor.me/cards/types"
	"7elements.ztaylor.me/elements"
	"ztaylor.me/json"
)

type Card struct {
	Id    int
	Image string
	ctypes.CardType
	Costs elements.Stack
}

func NewCard() *Card {
	return &Card{
		Costs: elements.Stack{},
	}
}

func Json(card *Card, body *Body, texts *Texts) json.Json {
	j := json.Json{
		"id":          card.Id,
		"image":       card.Image,
		"name":        texts.Name,
		"type":        card.CardType.String(),
		"description": texts.Description,
		"flavor":      texts.Flavor,
		"costs":       card.Costs.Copy(),
	}

	if body != nil {
		j["body"] = "true"
		j["attack"] = body.Attack
		j["health"] = body.Health
	}

	return j
}
