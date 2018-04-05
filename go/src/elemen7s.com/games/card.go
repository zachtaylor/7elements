package games

import (
	"elemen7s.com"
	"fmt"
	"ztaylor.me/js"
)

type Cards map[int]*Card

type Card struct {
	Id       int
	Username string
	Awake    bool
	*vii.Card
	*vii.CardText
	*vii.CardBody
	vii.Powers
}

func NewCard(card *vii.Card, text *vii.CardText) *Card {
	return &Card{
		Card:     card,
		CardText: text,
		CardBody: card.CardBody.Copy(),
		Powers:   card.Powers.Copy(),
	}
}

func (c *Card) IsRegistered() bool {
	return c.Id > 0
}

func (cards Cards) Json() js.Object {
	data := js.Object{}
	for gcid, card := range cards {
		data[fmt.Sprintf(`%d`, gcid)] = card.Json()
	}
	return data
}

func (card *Card) Json() js.Object {
	return js.Object{
		"gcid":     card.Id,
		"cardid":   card.Card.Id,
		"name":     card.CardText.Name,
		"username": card.Username,
		"image":    card.Card.Image,
		"awake":    card.Awake,
		"powers":   card.Powers.JsonWithText(card.CardText),
		"body":     card.CardBody.Json(),
	}
}

func (cards Cards) String() string {
	collapse := make([]int, len(cards))
	for i, card := range cards {
		collapse[i] = card.Card.Id
	}
	return fmt.Sprintf("games.Cards%v", collapse)
}
