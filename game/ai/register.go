package ai

import (
	"math/rand"
	"time"

	"github.com/zachtaylor/7elements/account"
	"github.com/zachtaylor/7elements/deck"
)

// GetAccountDeck creates a deck for the AI to play
func GetAccountDeck(service deck.PrototypeService) *account.Deck {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	decks, err := service.GetAll()
	if err != nil {
		return nil
	}
	i := (r.Int() % len(decks)) + 1
	return account.NewDeckWith(decks[i], Username)
}
