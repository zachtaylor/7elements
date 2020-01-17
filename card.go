package vii

import (
	"fmt"
	"sort"
	"strings"

	"ztaylor.me/cast"
)

type Card struct {
	Id     int
	Name   string
	Text   string
	Type   CardType
	Image  string
	Costs  ElementMap
	Body   *Body
	Powers Powers
}

func NewCard() *Card {
	return &Card{
		Costs:  ElementMap{},
		Powers: NewPowers(),
	}
}

func (c *Card) JSON() cast.JSON {
	return cast.JSON{
		"id":     c.Id,
		"image":  c.Image,
		"name":   c.Name,
		"text":   c.Text,
		"type":   c.Type.String(),
		"powers": c.Powers.JSON(),
		"costs":  c.Costs.JSON(),
		"body":   c.Body.JSON(),
	}
}

func (c *Card) String() string {
	return cast.StringN(
		`{`,
		c.Id,
		` `, c.Name,
		`}`,
	)
}

type Cards map[int]*Card

func (cards Cards) JSON() fmt.Stringer {
	json := make([]string, 0)
	keys := make([]int, len(cards))

	var i int
	for k := range cards {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	for _, id := range keys {
		json = append(json, cards[id].JSON().String())
	}
	return cast.Stringer(`[` + strings.Join(json, ",") + `]`)
}

type CardService interface {
	Start() error
	Get(cardid int) (*Card, error)
	GetAll() Cards
}
