package deck

import (
	"github.com/cznic/mathutil"
	"github.com/zachtaylor/7elements/card"
)

// T is a Deck
type T struct {
	Username      string
	AccountDeckID int
	Cards         []*card.T
}

// New returns a new Deck
func New() *T {
	return &T{
		Cards: make([]*card.T, 0),
	}
}

// Draw removes the top card of the Deck and returns it
func (deck *T) Draw() *card.T {
	if len(deck.Cards) < 1 {
		return nil
	}
	card := deck.Cards[0]
	deck.Cards = deck.Cards[1:]
	return card
}

// Prepend places a Card on top of the Deck
func (deck *T) Prepend(c *card.T) {
	deck.Cards = append([]*card.T{c}, deck.Cards...)
}

// Append places a Card on bottom of the Deck
func (deck *T) Append(c *card.T) {
	deck.Cards = append(deck.Cards, c)
}

// Shuffle randomizes the order of Cards in the Deck
func (deck *T) Shuffle() {
	shuffleRandom, _ := mathutil.NewFC32(0, len(deck.Cards)-1, true)
	cp := make([]*card.T, len(deck.Cards))
	for i := 0; i < len(deck.Cards); i++ {
		rand := shuffleRandom.Next()
		cp[rand] = deck.Cards[i]
	}
	deck.Cards = cp
}
