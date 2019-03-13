package engine

import (
	"os"
	"sync"
	"time"

	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/keygen"
	"ztaylor.me/log"
)

func init() {
	game.Service = NewService()
}

type Service struct {
	games     map[string]*game.T
	gamesLock sync.Mutex
	queue     []*game.Search
	queueLock sync.Mutex
}

func NewService() *Service {
	service := &Service{
		games: make(map[string]*game.T),
		queue: make([]*game.Search, 0),
	}
	go service.watchQueue()
	return service
}

func (service *Service) New() *game.T {
	loglvl := "debug"
	logger := log.NewLogger()
	logger.SetLevel(loglvl)
	tStart := time.Now()
	key := service.keygen()
	service.gamesLock.Lock()
	for ; service.games[key] != nil; key = service.keygen() {
	}
	logger.SetFile("log/game/" + key + ".log")
	g := game.New(key, logger, game.NewDefaultSettings())
	service.games[key] = g
	service.gamesLock.Unlock()
	log.WithFields(log.Fields{
		"Key":  key,
		"Time": time.Now().Sub(tStart),
	}).Info("engine/game_service: game added")
	return g
}

func (service *Service) Get(id string) *game.T {
	return service.games[id]
}

func (service *Service) Forget(id string) {
	service.gamesLock.Lock()
	delete(service.games, id)
	service.gamesLock.Unlock()
}

func (service *Service) GetPlayerGames(name string) []string {
	games := make([]string, 0)
	service.gamesLock.Lock()
	for _, g := range service.games {
		for sn, _ := range g.Seats {
			if name == sn {
				games = append(games, g.ID())
			}
		}
	}
	service.gamesLock.Unlock()
	return games
}

func (service *Service) GetPlayerSearch(name string) *game.Search {
	service.queueLock.Lock()
	defer service.queueLock.Unlock()
	for _, search := range service.queue {
		if search.Deck.Username == name {
			return search
		}
	}
	return nil
}

func (service *Service) StartPlayerSearch(deck *vii.AccountDeck) *game.Search {
	service.queueLock.Lock()
	search := game.NewSearch(deck)
	service.queue = append(service.queue, search)
	service.queueLock.Unlock()
	return search
}

func (service *Service) keygen() string {
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

func (service *Service) match(search0, search1 *game.Search) bool {
	return true
}

func (service *Service) watchQueue() {
	i := 0
	var search0, searchi *game.Search
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
				Watch(game)
				break
			}
		}
		service.queueLock.Unlock()
	}
}
