package ai

import (
	"math/rand"
	"time"

	"github.com/zachtaylor/7elements/deck"
	"github.com/zachtaylor/7elements/game"
)

// GetDeck creates a deck for the AI to play
func GetDeck(version *game.Version) *deck.Prototype {
	decks := version.GetDecks()
	r := rand.New(rand.NewSource(time.Now().Unix()))
	i := (r.Int() % len(decks)) + 1
	return decks[i]
}
