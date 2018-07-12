package vii

import (
	"fmt"
)

type GameCards map[string]*GameCard

type GameCard struct {
	Id       string
	Username string
	IsAwake  bool
	IsToken  bool
	*Card
	*CardText
	*CardBody
	Powers
}

func NewGameCard(c *Card, t *CardText) *GameCard {
	return &GameCard{
		Card:     c,
		CardText: t,
		CardBody: c.CardBody.Copy(),
		Powers:   c.Powers.Copy(),
	}
}

func (card GameCard) String() string {
	return card.Id + ":" + card.CardText.Name
}

func (cards GameCards) String() string {
	collapse := make([]int, len(cards))
	i := 0
	for _, c := range cards {
		collapse[i] = c.Card.Id
		i++
	}
	return fmt.Sprintf("games.GameCards%v", collapse)
}

func (c *GameCard) IsRegistered() bool {
	return len(c.Id) < 1
}

func (cards GameCards) Json() Json {
	data := Json{}
	for gcid, c := range cards {
		data[gcid] = c.Json()
	}
	return data
}

func (c *GameCard) Json() Json {
	return Json{
		"gcid":     c.Id,
		"cardid":   c.Card.Id,
		"name":     c.CardText.Name,
		"username": c.Username,
		"image":    c.Card.Image,
		"awake":    c.IsAwake,
		"powers":   c.Powers.JsonWithText(c.CardText),
		"body":     c.CardBody.Json(),
	}
}
