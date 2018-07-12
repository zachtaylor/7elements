package engine

import (
	"elemen7s.com"
	"os"
	"sync"
	"time"
	"ztaylor.me/keygen"
	"ztaylor.me/log"
)

func init() {
	vii.GameService = NewGameService()
}

type GameService struct {
	Games map[string]*vii.Game
	sync.Mutex
}

func NewGameService() *GameService {
	return &GameService{
		Games: make(map[string]*vii.Game),
	}
}

func (service *GameService) New() *vii.Game {
	game := vii.NewGame()
	service.Add(game)
	game.Logger.SetLevel("debug")
	game.Logger.SetFile("log/game/" + game.Key + ".log")
	return game
}

func (service *GameService) Add(game *vii.Game) {
	tStart := time.Now()
	log.Warn("engine-service: add")
	service.Lock()
	game.Key = service.keygen()
	service.Games[game.Key] = game
	service.Unlock()
	log.WithFields(log.Fields{
		"Key":  game.Key,
		"Time": time.Now().Sub(tStart),
	}).Info("engine-service: game added")
}

func (service *GameService) keygen() string {
	for key := keygen.NewVal(); ; key = keygen.NewVal() {
		if stat, err := os.Stat("log/game/" + key + ".log"); err == nil {
			log.WithFields(log.Fields{
				"Stat": stat,
			}).Warn("engine-service: proposed key has log file")
		} else if service.Games[key] != nil {
			log.WithFields(log.Fields{
				"Stat": stat,
			}).Warn("engine-service: proposed key has log file")
		} else {
			return key
		}
	}
}

func (service *GameService) Get(id string) *vii.Game {
	return service.Games[id]
}

func (service *GameService) Forget(id string) {
	service.Lock()
	delete(service.Games, id)
	service.Unlock()
}

func (service *GameService) GetPlayerGames(name string) []string {
	games := make([]string, 0)
	service.Lock()
	for _, g := range service.Games {
		for sn, _ := range g.Seats {
			if name == sn {
				games = append(games, g.Key)
			}
		}
	}
	service.Unlock()
	return games
}
