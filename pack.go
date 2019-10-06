package vii

import (
	"fmt"
	"sort"
	"strings"

	"ztaylor.me/cast"
)

type Pack struct {
	ID    int
	Name  string
	Size  int
	Cost  int
	Image string
	Cards []*PackChance
}

func NewPack() *Pack {
	return &Pack{
		Cards: make([]*PackChance, 0),
	}
}

func (p *Pack) JSON() cast.JSON {
	cards := make([]string, 0)
	for _, card := range p.Cards {
		cards = append(cards, cast.StringI(card.CardID))
	}
	return cast.JSON{
		"id":    p.ID,
		"name":  p.Name,
		"size":  p.Size,
		"cost":  p.Cost,
		"image": p.Image,
		"cards": cast.Stringer(`[` + strings.Join(cards, ",") + `]`),
	}
}

type PackChance struct {
	PackID int
	CardID int
	Weight int
}

type Packs map[int]*Pack

func (packs Packs) JSON() fmt.Stringer {
	json := make([]string, 0)
	keys := make([]int, len(packs))
	var i int
	for k := range packs {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	for _, k := range keys {
		json = append(json, packs[k].JSON().String())
	}
	return cast.Stringer(`[` + strings.Join(json, ",") + `]`)
}

type PackService interface {
	Get(int) (*Pack, error)
	GetAll() (Packs, error)
}
