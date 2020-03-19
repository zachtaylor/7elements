package game

import (
	"github.com/cznic/mathutil"
	"github.com/zachtaylor/7elements/card"
)

type Deck struct {
	Username      string
	AccountDeckID int
	Cards         []*card.T
}

func NewDeck() *Deck {
	return &Deck{
		Cards: make([]*card.T, 0),
	}
}

func (deck *Deck) Draw() *card.T {
	if len(deck.Cards) < 1 {
		return nil
	}
	card := deck.Cards[0]
	deck.Cards = deck.Cards[1:]
	return card
}

func (deck *Deck) Prepend(c *card.T) {
	deck.Cards = append([]*card.T{c}, deck.Cards...)
}

func (deck *Deck) Append(c *card.T) {
	deck.Cards = append(deck.Cards, c)
}

func (deck *Deck) Shuffle() {
	shuffleRandom, _ := mathutil.NewFC32(0, len(deck.Cards)-1, true)
	cp := make([]*card.T, len(deck.Cards))
	for i := 0; i < len(deck.Cards); i++ {
		rand := shuffleRandom.Next()
		cp[rand] = deck.Cards[i]
	}
	deck.Cards = cp
}
