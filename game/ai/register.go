package ai

import (
	"math/rand"
	"time"

	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/deck"
)

// GetAccountDeck creates a deck for the AI to play
func GetAccountDeck(service deck.PrototypeService) *vii.AccountDeck {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	decks, err := service.GetAll()
	if err != nil {
		return nil
	}
	i := (r.Int() % len(decks)) + 1
	return vii.NewAccountDeckWith(decks[i], Username)
}
