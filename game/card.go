package game

import (
	"fmt"

	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/element"
	"ztaylor.me/cast"
)

type Card struct {
	ID       string
	Username string
	Proto    *card.Prototype
}

func NewCard(c *card.Prototype) *Card {
	return &Card{
		Proto: c,
	}
}

func (c *Card) IsRegistered() bool {
	return len(c.ID) < 1
}

func (c *Card) String() string {
	if c == nil {
		return `<nil>`
	}
	return cast.StringN(
		`{`,
		c.ID,
		` user:`, c.Username,
		` card:`, c.Proto.String(),
		`}`,
	)
}

// JSON returns a representation of a game card
func (c *Card) JSON() cast.JSON {
	if c == nil {
		return nil
	}
	return cast.JSON{
		"id":       c.ID,
		"cardid":   c.Proto.ID,
		"name":     c.Proto.Name,
		"costs":    c.Proto.Costs.JSON(),
		"text":     c.Proto.Text,
		"username": c.Username,
		"image":    c.Proto.Image,
		"type":     c.Proto.Type.String(),
		"powers":   c.Proto.Powers.JSON(),
	}
}

// Cards is a map of cid->Card
type Cards map[string]*Card

// Devotion returns the ElementMap describing the devotion of this group of cards
func (cards Cards) Devotion() element.Count {
	devo := element.Count{}
	for _, c := range cards {
		for e, count := range c.Proto.Costs {
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
		collapse[i] = c.Proto.ID
		i++
	}
	return fmt.Sprintf("%v", collapse)
}

// JSON returns a representation of a set of game cards
func (cards Cards) JSON() cast.JSON {
	data := cast.JSON{}
	for cid, c := range cards {
		data[cid] = c.JSON()
	}
	return data
}
