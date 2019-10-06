package engine

import (
	"os"
	"sync"
	"time"

	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/event"
	"ztaylor.me/keygen"
	"ztaylor.me/log"
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
		log.NewService(log.LevelInfo, log.DefaultFormatter(false), f),
		service.chat,
	))
	g.Register(a)
	g.Register(b)
	g.State = g.NewState(event.NewStartEvent(a.Username))
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

func (service *Service) CardTriggeredEvents(g *game.T, seat *game.Seat, card *game.Card, trigger string, origTarget interface{}) []game.Event {
	log := g.Log().With(log.Fields{
		"Seat":    seat.Username,
		"Card":    card.Card.Name,
		"Trigger": trigger,
	}).Tag("engine/card-trigger")
	powers := card.Powers.GetTrigger(trigger)
	if len(powers) < 1 {
		log.Debug("empty")
		return nil
	}
	events := make([]game.Event, 0)
	for _, p := range powers {
		var target interface{}
		if p.Target == "self" {
			target = card
		} else if origTarget != nil {
			target = origTarget
		}

		if target != nil {
			events = append(events, event.NewTriggerEvent(seat.Username, card, p, target))
		} else {
			script := game.Scripts[p.Script]
			events = append(events, event.NewTargetEvent(
				seat.Username,
				p.Target,
				p.Text,
				func(val string) []game.Event {
					return script(g, seat, val)
				},
			))
		}
	}
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
