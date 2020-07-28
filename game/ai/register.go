package ai

import (
	"math/rand"
	"time"

	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/deck"
	"ztaylor.me/log"
)

// GetDeck creates a deck for the AI to play
func GetDeck(log log.Service, cards card.PrototypeService, decks deck.PrototypeService) *deck.T {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	set, err := decks.GetUser(Username)
	if err != nil {
		return nil
	}
	i := (r.Int() % len(set)) + 1
	return deck.New(log, cards, set[i], Username)
}
