package ai

import (
	"math/rand"
	"time"

	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/deck"
	"taylz.io/log"
)

// GetDeck creates a deck for the AI to play
func GetDeck(log *log.T, cards card.Prototypes, decks deck.Prototypes, username string) *deck.T {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	i := (r.Int() % len(decks)) + 1
	return deck.New(log, cards, decks[i], username)
}
