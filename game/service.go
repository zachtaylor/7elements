package game

import "github.com/zachtaylor/7elements/account"

// Service defines a game engine
type Service interface {
	// New creates a Game between 2 players
	New(*account.Deck, *account.Deck) *T
	// Get returns a Game by ID
	Get(string) *T
	// FindUsername returns Game containing the username
	FindUsername(string) *T
	// Search starts a PVP game search
	Search(deck *account.Deck) *Search

	Trigger(g *T, seat *Seat, token *Token, name string, arg interface{}) []Stater
}
