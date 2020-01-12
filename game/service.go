package game

import vii "github.com/zachtaylor/7elements"

// Service defines a game engine
type Service interface {
	// New creates a Game between 2 players
	New(*vii.AccountDeck, *vii.AccountDeck) *T
	// Get returns a Game by ID
	Get(string) *T
	// FindUsername returns Game containing the username
	FindUsername(string) *T
	// Search starts a PVP game search
	Search(deck *vii.AccountDeck) *Search

	Trigger(g *T, seat *Seat, token *Token, name string, arg interface{}) []Stater
}
