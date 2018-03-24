package games

import (
	"elemen7s.com"
	"elemen7s.com/cards"
	"fmt"
	"ztaylor.me/js"
)

type Cards map[int]*Card

type Card struct {
	Id       int
	Username string
	Awake    bool
	*cards.Card
	*vii.CardText
	*cards.Body
	cards.Powers
}

func NewCard(card *cards.Card, text *vii.CardText) *Card {
	return &Card{
		Card:     card,
		CardText: text,
		Body:     copyOrNilBody(card.Body),
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
		"powers":   card.Powers.Json(),
		"body":     card.Body.Json(),
	}
}

func (cards Cards) String() string {
	collapse := make([]int, len(cards))
	for i, card := range cards {
		collapse[i] = card.Card.Id
	}
	return fmt.Sprintf("games.Cards%v", collapse)
}

func copyOrNilBody(body *cards.Body) *cards.Body {
	if body == nil {
		return nil
	}

	return body.Copy()
}
