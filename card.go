package vii

import (
	"strings"

	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/element"
	"github.com/zachtaylor/7elements/power"
	"ztaylor.me/cast"
)

type Card struct {
	ID     int
	Name   string
	Text   string
	Type   card.Type
	Image  string
	Costs  element.Count
	Body   *card.Body
	Powers power.Set
}

func NewCard() *Card {
	return &Card{
		Costs:  element.Count{},
		Powers: power.NewSet(),
	}
}

func (c *Card) JSON() cast.JSON {
	return cast.JSON{
		"id":     c.ID,
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
	return cast.StringN(`{`, c.ID, ` `, c.Name, `}`)
}

type Cards map[int]*Card

func (cards Cards) JSON() cast.IStringer {
	json := make([]string, 0)
	keys := make([]int, len(cards))

	var i int
	for k := range cards {
		keys[i] = k
		i++
	}
	cast.SortInts(keys)
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
