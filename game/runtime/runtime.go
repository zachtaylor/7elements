package runtime

// import (
// 	"os"
// 	"sync"
// 	"time"

// 	"github.com/zachtaylor/7elements/account"
// 	"github.com/zachtaylor/7elements/card"
// 	"github.com/zachtaylor/7elements/chat"
// 	"github.com/zachtaylor/7elements/deck"
// 	"github.com/zachtaylor/7elements/game"
// 	"github.com/zachtaylor/7elements/game/engine"
// 	"ztaylor.me/keygen"
// 	"ztaylor.me/log"
// )

// type Service = game.Service

// type T struct {
// 	log       log.Service
// 	accounts  account.Service
// 	cards     card.PrototypeService
// 	games     map[string]*game.T
// 	gamesLock sync.Mutex
// 	queue     []*Search
// 	queueLock sync.Mutex
// 	chat      chat.Service
// }

// func New(logger log.Service, accounts account.Service, cards card.PrototypeService, chat chat.Service) *T {
// 	service := &T{
// 		log:      logger,
// 		accounts: accounts,
// 		cards:    cards,
// 		games:    make(map[string]*game.T),
// 		queue:    make([]*Search, 0),
// 		chat:     chat,
// 	}
// 	go service.watchQueue()
// 	return service
// }
// func (t *T) _isService() Service {
// 	return t
// }

// func (t *T) New(a, b *deck.T) *game.T {
// 	key := t.keygen()
// 	f, _ := os.OpenFile("log/game/"+key+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
// 	g := game.New(key, game.NewRuntime(
// 		engine.NewService(),
// 		t.accounts,
// 		t.cards,
// 		70*time.Second,
// 		f,
// 		t.chat.New("game#"+key, 8),
// 	))
// 	g.Register(a)
// 	g.Register(b)
// 	g.State = g.NewState(engine.NewStart(a.User), g.Runtime.Timeout)
// 	t.set(key, g)
// 	go t.runGame(g)
// 	g.Runtime.Logger.New().Add("Key", key).Info("new game started")
// 	return g
// }
// func (t *T) runGame(g *game.T) {
// 	engine.Run(g)
// 	t.gamesLock.Lock()
// 	delete(t.games, g.ID())
// 	t.gamesLock.Unlock()
// }

// func (t *T) Get(id string) *game.T {
// 	return t.games[id]
// }

// func (t *T) set(id string, g *game.T) {
// 	t.gamesLock.Lock()
// 	if g == nil {
// 		delete(t.games, id)
// 	} else {
// 		t.games[id] = g
// 	}
// 	t.gamesLock.Unlock()
// }

// func (t *T) Forget(id string) {
// 	t.gamesLock.Lock()
// 	delete(t.games, id)
// 	t.gamesLock.Unlock()
// }

// func (t *T) FindUsername(name string) *game.T {
// 	t.gamesLock.Lock()
// 	defer t.gamesLock.Unlock()
// 	for _, g := range t.games {
// 		for _, s := range g.Seats {
// 			if name == s.Username {
// 				return g
// 			}
// 		}
// 	}
// 	return nil
// }

// func (t *T) Search(deck *deck.T) *Search {
// 	if t.FindUsername(deck.User) != nil {
// 		return nil
// 	}
// 	t.queueLock.Lock()
// 	defer t.queueLock.Unlock()
// 	for _, search := range t.queue {
// 		if search.Deck.User == deck.User {
// 			return search
// 		}
// 	}
// 	search := NewSearch(deck)
// 	t.queue = append(t.queue, search)
// 	return search
// }

// func (t *T) keygen() (id string) {
// 	t.gamesLock.Lock()
// 	for {
// 		id = keygen.NewVal()
// 		if _, err := os.Stat("log/game/" + id + ".log"); err != nil {
// 			break
// 		}
// 	}
// 	if file, err := os.Create("log/game/" + id + ".log"); err == nil {
// 		file.Close()
// 	}
// 	t.gamesLock.Unlock()
// 	return
// }

// func (t *T) match(search0, search1 *Search) bool {
// 	return true
// }

// func (t *T) watchQueue() {
// 	i := 0
// 	var search0, searchi *Search
// 	for {
// 		<-time.After(time.Second)
// 		if len(t.queue) < 2 {
// 			continue
// 		}
// 		t.queueLock.Lock()
// 		search0 = t.queue[0]
// 		for i = 1; i < len(t.queue); i++ {
// 			searchi = t.queue[i]
// 			if t.match(search0, searchi) {
// 				t.New(search0.Deck, searchi.Deck)
// 				break
// 			}
// 		}
// 		t.queueLock.Unlock()
// 	}
// }
