package cards

import (
	"7elements.ztaylor.me/cards/types"
	"7elements.ztaylor.me/elements"
)

type Card struct {
	Id    int
	Image string
	ctypes.CardType
	Costs elements.Stack
}

func NewCard() *Card {
	return &Card{
		Costs: elements.Stack{},
	}
}
