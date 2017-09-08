package cards

import (
	"7elements.ztaylor.me/cards"
	"ztaylor.me/json"
)

func MakeCardJson(card *cards.Card, body *cards.Body, texts *cards.Texts) json.Json {
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
