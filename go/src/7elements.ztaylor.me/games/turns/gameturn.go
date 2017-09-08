package gameturns

import (
	"7elements.ztaylor.me/elements"
	"7elements.ztaylor.me/games/cards"
)

type GameTurn struct {
	GameId   int
	Id       int
	Username string
	elements.Element
	Cards []*gamecards.GameCard
}

func NewGameTurn() *GameTurn {
	return &GameTurn{
		Cards: make([]*gamecards.GameCard, 0),
	}
}
