package games

// import (
// 	"sync"
// )

// var Cache = &gamecache{
// 	games: make(map[int]*vii.Game),
// }

// type gamecache struct {
// 	sync.Mutex
// 	games map[int]*vii.Game
// }

// func (cache *gamecache) Add(g *vii.Game) {
// 	cache.Lock()
// 	cache.games[g.Id] = g
// 	cache.Unlock()
// }

// func (cache *gamecache) Get(id int) *vii.Game {
// 	cache.Lock()
// 	defer cache.Unlock()
// 	return cache.games[id]
// }

// func (cache *gamecache) GetPlayerGames(name string) []int {
// 	games := make([]int, 0)
// 	cache.Lock()
// 	for _, g := range cache.games {
// 		for _, s := range g.Seats {
// 			if s.Username == name {
// 				games = append(games, g.Id)
// 				break
// 			}
// 		}
// 	}
// 	cache.Unlock()
// 	return games
// }

// func (cache *gamecache) Remove(g *vii.Game) {
// 	cache.Lock()
// 	delete(cache.games, g.Id)
// 	cache.Unlock()
// }