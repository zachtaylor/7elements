package vii

var GameService interface {
	New() *Game
	Get(id string) *Game
	Forget(id string)
	GetPlayerGames(name string) []string
}
