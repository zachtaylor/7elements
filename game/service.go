package game

import "github.com/zachtaylor/7elements"

var Service interface {
	New() *T
	Get(id string) *T
	Forget(id string)
	GetPlayerGames(name string) []string
	GetPlayerSearch(name string) *Search
	StartPlayerSearch(deck *vii.AccountDeck) *Search
}
