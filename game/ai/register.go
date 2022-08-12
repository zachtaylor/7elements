package ai

import (
	"math/rand"
	"time"

	"github.com/zachtaylor/7elements/content"
	"github.com/zachtaylor/7elements/deck"
)

// GetDeck creates a deck for the AI to play
func GetDeck(content content.T) *deck.Prototype {
	decks := content.Decks()
	r := rand.New(rand.NewSource(time.Now().Unix()))
	i := (r.Int() % len(decks)) + 1
	return decks[i]
}
