package cards

import (
	"7elements.ztaylor.me"
	"7elements.ztaylor.me/server/json"
)

func MakeCardJson(card *SE.Card, cardtext *SE.CardText) json.Json {
	j := json.Json{}
	j["id"] = card.Id
	j["image"] = card.Image

	j["name"] = cardtext.Name
	j["description"] = cardtext.Description
	j["flavor"] = cardtext.Flavor

	elementcosts := make([]json.Json, 0)
	totalcost := 0
	for _, elementcost := range card.ElementCosts {
		elementcostJ := json.Json{
			"element": *elementcost.Element,
			"count":   elementcost.Count,
		}
		totalcost += int(elementcost.Count)
		elementcosts = append(elementcosts, elementcostJ)
	}

	j["totalcost"] = totalcost
	j["elementcosts"] = elementcosts

	return j
}
