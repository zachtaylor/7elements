package games

import (
	"sync"
)

var Cache = &gamecache{
	games: make(map[int]*Game),
}

type gamecache struct {
	sync.Mutex
	games map[int]*Game
}

func (cache *gamecache) Add(g *Game) {
	cache.Lock()
	cache.games[g.Id] = g
	cache.Unlock()
}

func (cache *gamecache) Get(id int) *Game {
	cache.Lock()
	defer cache.Unlock()
	return cache.games[id]
}

func (cache *gamecache) GetPlayerGames(name string) []int {
	games := make([]int, 0)
	cache.Lock()
	for _, g := range cache.games {
		for _, s := range g.Seats {
			if s.Username == name {
				games = append(games, g.Id)
				break
			}
		}
	}
	cache.Unlock()
	return games
}

func (cache *gamecache) Remove(g *Game) {
	cache.Lock()
	delete(cache.games, g.Id)
	cache.Unlock()
}
