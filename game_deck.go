package vii

import (
	"github.com/cznic/mathutil"
)

type GameDeck struct {
	Username      string
	AccountDeckID int
	Cards         []*GameCard
}

func NewGameDeck() *GameDeck {
	return &GameDeck{
		Cards: make([]*GameCard, 0),
	}
}

func (deck *GameDeck) Draw() *GameCard {
	if len(deck.Cards) < 1 {
		return nil
	}
	card := deck.Cards[0]
	deck.Cards = deck.Cards[1:]
	return card
}

func (deck *GameDeck) Prepend(card *GameCard) {
	deck.Cards = append([]*GameCard{card}, deck.Cards...)
}

func (deck *GameDeck) Append(card *GameCard) {
	deck.Cards = append(deck.Cards, card)
}

func (deck *GameDeck) Shuffle() {
	shuffleRandom, _ := mathutil.NewFC32(0, len(deck.Cards)-1, true)
	cp := make([]*GameCard, len(deck.Cards))
	for i := 0; i < len(deck.Cards); i++ {
		rand := shuffleRandom.Next()
		cp[rand] = deck.Cards[i]
	}
	deck.Cards = cp
}
