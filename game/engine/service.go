package engine

import (
	"os"
	"sync"
	"time"

	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/state"
	"github.com/zachtaylor/7elements/game/trigger"
	"ztaylor.me/cast"
	"ztaylor.me/keygen"
)

type Service struct {
	games     map[string]*game.T
	gamesLock sync.Mutex
	queue     []*game.Search
	queueLock sync.Mutex
	runtime   *vii.Runtime
	chat      chat.Service
}

func NewService(rt *vii.Runtime, chat chat.Service) *Service {
	service := &Service{
		games:   make(map[string]*game.T),
		queue:   make([]*game.Search, 0),
		runtime: rt,
		chat:    chat,
	}
	go service.watchQueue()
	return service
}

func (service *Service) New(a, b *vii.AccountDeck) *game.T {
	key := service.keygen()
	f, _ := os.OpenFile("log/game/"+key+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	g := game.New(key, game.NewRuntime(
		service.runtime,
		service,
		70*time.Second,
		f,
		service.chat,
	))
	g.Register(a)
	g.Register(b)
	g.State = g.NewState(state.NewStart(a.Username))
	service.set(key, g)
	go service.runner(g)
	service.runtime.Logger.New().Add("Key", key).Info("engine: new game started")
	return g
}
func (service *Service) runner(g *game.T) {
	Run(g)
	service.set(g.ID(), nil)
}

func (service *Service) Get(id string) *game.T {
	return service.games[id]
}

func (service *Service) set(id string, g *game.T) {
	service.gamesLock.Lock()
	if g == nil {
		delete(service.games, id)
	} else {
		service.games[id] = g
	}
	service.gamesLock.Unlock()
}

func (service *Service) Forget(id string) {
	service.gamesLock.Lock()
	delete(service.games, id)
	service.gamesLock.Unlock()
}

func (service *Service) FindUsername(name string) *game.T {
	service.gamesLock.Lock()
	defer service.gamesLock.Unlock()
	for _, g := range service.games {
		for _, s := range g.Seats {
			if name == s.Username {
				return g
			}
		}
	}
	return nil
}

func (service *Service) Search(deck *vii.AccountDeck) *game.Search {
	if service.FindUsername(deck.Username) != nil {
		return nil
	}
	service.queueLock.Lock()
	defer service.queueLock.Unlock()
	for _, search := range service.queue {
		if search.Deck.Username == deck.Username {
			return search
		}
	}
	search := game.NewSearch(deck)
	service.queue = append(service.queue, search)
	return search
}

func (service *Service) Trigger(g *game.T, seat *game.Seat, token *game.Token, name string, arg interface{}) []game.Stater {
	log := g.Log().With(cast.JSON{
		"Seat":    seat.Username,
		"Token":   token.String(),
		"Trigger": name,
	})
	powers := token.Powers.GetTrigger(name)
	if len(powers) < 1 {
		log.Debug("empty")
		return nil
	}
	events := make([]game.Stater, 0)
	for _, p := range powers {
		if p.Target != "self" && arg == nil {
			power := p
			events = append(events, state.NewTarget(
				seat.Username,
				p.Target,
				p.Text,
				func(val string) []game.Stater {
					return trigger.Power(g, seat, power, token, cast.NewArray(val))
				},
			))
		} else {
			events = append(events, state.NewTrigger(seat.Username, token, p, arg))
		}
	}
	log.Add("Events", events).Source().Debug()
	return events
}

func (service *Service) keygen() (key string) {
	service.gamesLock.Lock()
	for key = keygen.NewVal(); ; key = keygen.NewVal() {
		if _, ok := service.games[key]; ok { // respect key guard
			// continue
		} else if _, err := os.Stat("log/game/" + key + ".log"); err != nil {
			break
		}
	}
	service.games[key] = nil // key guard
	service.gamesLock.Unlock()
	return
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
				service.New(search0.Deck, searchi.Deck)
				break
			}
		}
		service.queueLock.Unlock()
	}
}
