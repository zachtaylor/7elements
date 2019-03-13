package game

import (
	"fmt"

	"github.com/zachtaylor/7elements"
)

type Cards map[string]*Card

func (cards Cards) String() string {
	collapse := make([]int, len(cards))
	i := 0
	for _, c := range cards {
		collapse[i] = c.Card.Id
		i++
	}
	return fmt.Sprintf("games.Cards%v", collapse)
}

func (cards Cards) Json() vii.Json {
	data := vii.Json{}
	for gcid, c := range cards {
		data[gcid] = c.Json()
	}
	return data
}

type Card struct {
	Id       string
	Username string
	IsAwake  bool
	IsToken  bool
	Card     *vii.Card
	Body     *vii.Body
	vii.Powers
}

func NewCard(c *vii.Card) *Card {
	return &Card{
		Card:   c,
		Body:   c.Body.Copy(),
		Powers: c.Powers.Copy(),
	}
}

func (card Card) String() string {
	return card.Id + ":" + card.Card.Name
}

func (c *Card) IsRegistered() bool {
	return len(c.Id) < 1
}

func (c *Card) Json() vii.Json {
	return vii.Json{
		"gcid":     c.Id,
		"cardid":   c.Card.Id,
		"name":     c.Card.Name,
		"costs":    c.Card.Costs,
		"text":     c.Card.Text,
		"username": c.Username,
		"image":    c.Card.Image,
		"awake":    c.IsAwake,
		"powers":   c.Powers.Json(),
		"body":     c.Body.Json(),
	}
}
