package game

import (
	"fmt"

	vii "github.com/zachtaylor/7elements"
	"ztaylor.me/cast"
)

type Card struct {
	Id       string
	Username string
	IsAwake  bool
	Card     *vii.Card
	Body     *vii.Body
	Powers   vii.Powers
}

func NewCard(c *vii.Card) *Card {
	return &Card{
		Card:   c,
		Body:   c.Body.Copy(),
		Powers: c.Powers.Copy(),
	}
}

func (c *Card) IsRegistered() bool {
	return len(c.Id) < 1
}

func (c *Card) String() string {
	if c == nil {
		return `<nil>`
	}
	return `game.Card{` + c.Print() + `}`
}

// Print returns a detailed compressed string representation
func (c *Card) Print() string {
	return c.Id + ":" + c.Card.Name
}

// JSON returns a representation of a game card
func (c *Card) JSON() cast.JSON {
	if c == nil {
		return nil
	}

	return cast.JSON{
		"gcid":     c.Id,
		"cardid":   c.Card.Id,
		"name":     c.Card.Name,
		"costs":    c.Card.Costs.JSON(),
		"text":     c.Card.Text,
		"username": c.Username,
		"image":    c.Card.Image,
		"awake":    c.IsAwake,
		"powers":   c.Powers.JSON(),
		"body":     c.Body.JSON(),
		"type":     c.Card.Type.String(),
	}
}

// Cards is a map of gcid->Card
type Cards map[string]*Card

// Devotion returns the ElementMap describing the devotion of this group of cards
func (cards Cards) Devotion() vii.ElementMap {
	devo := vii.ElementMap{}
	for _, c := range cards {
		for e, count := range c.Card.Costs {
			devo[e] += count
		}
	}
	return devo
}

func (cards Cards) String() string {
	return `game.Cards` + cards.Print()
}

func (cards Cards) Print() string {
	collapse := make([]int, len(cards))
	i := 0
	for _, c := range cards {
		collapse[i] = c.Card.Id
		i++
	}
	return fmt.Sprintf("%v", collapse)
}

// JSON returns a representation of a set of game cards
func (cards Cards) JSON() cast.JSON {
	data := cast.JSON{}
	for gcid, c := range cards {
		data[gcid] = c.JSON()
	}
	return data
}
