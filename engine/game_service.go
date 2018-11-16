package engine

import (
	"os"
	"sync"
	"time"

	"github.com/zachtaylor/7elements"
	"ztaylor.me/keygen"
	"ztaylor.me/log"
)

func init() {
	vii.GameService = NewGameService()
}

type GameService struct {
	games     map[string]*vii.Game
	gamesLock sync.Mutex
	queue     []*vii.GameSearch
	queueLock sync.Mutex
}

func NewGameService() *GameService {
	service := &GameService{
		games: make(map[string]*vii.Game),
		queue: make([]*vii.GameSearch, 0),
	}
	go service.watchQueue()
	return service
}

func (service *GameService) New() *vii.Game {
	game := vii.NewGame(vii.NewDefaultGameSettings())
	service.Add(game)
	game.Logger.SetLevel("debug")
	game.Logger.SetFile("log/game/" + game.Key + ".log")
	return game
}

func (service *GameService) Add(game *vii.Game) {
	tStart := time.Now()
	service.gamesLock.Lock()
	for ; service.games[game.Key] != nil; game.Key = service.keygen() {
	}
	service.games[game.Key] = game
	service.gamesLock.Unlock()
	log.WithFields(log.Fields{
		"Key":  game.Key,
		"Time": time.Now().Sub(tStart),
	}).Info("engine/game_service: game added")
}

func (service *GameService) Get(id string) *vii.Game {
	return service.games[id]
}

func (service *GameService) Forget(id string) {
	service.gamesLock.Lock()
	delete(service.games, id)
	service.gamesLock.Unlock()
}

func (service *GameService) Watch(game *vii.Game) {
	go game.Log().Protect(func() {
		Run(game)
	})
}

func (service *GameService) GetPlayerGames(name string) []string {
	games := make([]string, 0)
	service.gamesLock.Lock()
	for _, g := range service.games {
		for sn, _ := range g.Seats {
			if name == sn {
				games = append(games, g.Key)
			}
		}
	}
	service.gamesLock.Unlock()
	return games
}

func (service *GameService) GetPlayerSearch(name string) *vii.GameSearch {
	service.queueLock.Lock()
	defer service.queueLock.Unlock()
	for _, search := range service.queue {
		if search.Deck.Username == name {
			return search
		}
	}
	return nil
}

func (service *GameService) StartPlayerSearch(deck *vii.AccountDeck) *vii.GameSearch {
	service.queueLock.Lock()
	search := vii.NewGameSearch(deck)
	service.queue = append(service.queue, search)
	service.queueLock.Unlock()
	return search
}

func (service *GameService) keygen() string {
	for key := keygen.NewVal(); ; key = keygen.NewVal() {
		if stat, err := os.Stat("log/game/" + key + ".log"); err == nil {
			log.WithFields(log.Fields{
				"Stat": stat,
			}).Warn("engine/game_service: proposed key already exists")
		} else {
			return key
		}
	}
}

func (service *GameService) match(search0, search1 *vii.GameSearch) bool {
	return true
}

func (service *GameService) watchQueue() {
	i := 0
	var search0, searchi *vii.GameSearch
	for {
		<-time.After(time.Second)
		if len(service.queue) < 2 {
			continue
		}
		service.queueLock.Lock()
		search0 = service.queue[0]
		for i = 1; i < len(service.queue); i++ {
			searchi = service.queue[i]
			if service.match(search0, searchi) {
				game := service.New()
				game.Register(search0.Deck)
				game.Register(searchi.Deck)
				service.Add(game)
				service.Watch(game)
				break
			}
		}
		service.queueLock.Unlock()
	}
}
