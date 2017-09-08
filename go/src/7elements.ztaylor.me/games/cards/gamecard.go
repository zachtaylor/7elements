package gamecards

import (
	"ztaylor.me/json"
)

type GameCard struct {
	CardId     int
	Username   string
	GameCardId int
	Active     bool
	Data       map[string]interface{}
}

type Stack []*GameCard

func New() *GameCard {
	return &GameCard{
		Data: make(map[string]interface{}),
	}
}

func Build(cardid int, username string, gcid int) *GameCard {
	return &GameCard{
		CardId:     cardid,
		Username:   username,
		GameCardId: gcid,
		Data:       make(map[string]interface{}),
	}
}

func (card *GameCard) Json() json.Json {
	return json.Json{
		"cardid": card.CardId,
		"gcid":   card.GameCardId,
		"active": card.Active,
		"data":   card.Data,
	}
}

func (stack Stack) Json() []json.Json {
	data := make([]json.Json, 0)
	for _, card := range stack {
		data = append(data, card.Json())
	}
	return data
}
