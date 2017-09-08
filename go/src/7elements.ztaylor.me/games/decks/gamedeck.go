package gamedecks

import (
	"7elements.ztaylor.me/games/cards"
)

type GameDeck struct {
	Cards []*gamecards.GameCard
}

func New() *GameDeck {
	return &GameDeck{
		Cards: make([]*gamecards.GameCard, 0),
	}
}

func (deck *GameDeck) Draw() *gamecards.GameCard {
	card := deck.Cards[0]
	deck.Cards = deck.Cards[1:]
	return card
}

func (deck *GameDeck) Append(card *gamecards.GameCard) {
	deck.Cards = append(deck.Cards, card)
}

func (deck *GameDeck) Shuffle() {
}
