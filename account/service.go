package account

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/deck"
)

// Service provides Accounts
type Service interface {
	// Get loads an account from back end
	Get(string) (*T, error)
	// Count returns a number of registered accounts from back end
	Count() (int, error)
	// Insert creates an account on back end
	Insert(*T) error
	// UpdateCoins updates an accounts coin count on back end
	UpdateCoins(*T) error
	// UpdateEmail updates an accounts email
	UpdateEmail(*T) error
	// UpdateLogin updates an accounts login time on back end
	UpdateLogin(*T) error
	// UpdatePassword updates an accounts password on back end
	UpdatePassword(*T) error
	// Delete removes an account
	Delete(string) error
	// GetCards performs a get operation to the cards associated with the username
	GetCards(username string) (card.Count, error)
	// DeleteDecks removes all decks for a username
	GetDecks(username string) (deck.Prototypes, error)
}
