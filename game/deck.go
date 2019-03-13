package game

import (
	"github.com/cznic/mathutil"
)

type Deck struct {
	Username      string
	AccountDeckID int
	Cards         []*Card
}

func NewDeck() *Deck {
	return &Deck{
		Cards: make([]*Card, 0),
	}
}

func (deck *Deck) Draw() *Card {
	if len(deck.Cards) < 1 {
		return nil
	}
	card := deck.Cards[0]
	deck.Cards = deck.Cards[1:]
	return card
}

func (deck *Deck) Prepend(card *Card) {
	deck.Cards = append([]*Card{card}, deck.Cards...)
}

func (deck *Deck) Append(card *Card) {
	deck.Cards = append(deck.Cards, card)
}

func (deck *Deck) Shuffle() {
	shuffleRandom, _ := mathutil.NewFC32(0, len(deck.Cards)-1, true)
	cp := make([]*Card, len(deck.Cards))
	for i := 0; i < len(deck.Cards); i++ {
		rand := shuffleRandom.Next()
		cp[rand] = deck.Cards[i]
	}
	deck.Cards = cp
}
