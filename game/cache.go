package game

import (
	"os"
	"sync"
	"time"

	"github.com/zachtaylor/7elements/deck"
	"ztaylor.me/keygen"
)

type Cache struct {
	Settings CacheSettings
	games    map[string]*T
	gameμ    sync.Mutex
	search   []*Search
	srchμ    sync.Mutex
}

func NewCache(settings CacheSettings) *Cache {
	c := &Cache{
		Settings: settings,
		games:    make(map[string]*T),
		search:   make([]*Search, 0),
	}
	go c.watchQueue()
	return c
}

func (c *Cache) New(a, b *deck.T) *T {
	var key string
	c.gameμ.Lock()
	for ok := true; ok; {
		key = keygen.NewVal()
		if _, ok := c.games[key]; ok { // respect key guard
			// continue
		} else if _, err := os.Stat("log/game/" + key + ".log"); err != nil {
			break
		}
	}
	c.games[key] = nil // key guard
	c.gameμ.Unlock()

	f, _ := os.OpenFile("log/game/"+key+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	g := New(NewSettings(
		c.Settings.Engine,
		70*time.Second,
		c.Settings.Accounts,
		c.Settings.Cards,
		c.Settings.Chats.New("game#"+key, 7),
		f,
		c.Settings.Sockets,
	), key)

	g.Register(a)
	g.Register(b)
	g.State = g.NewState(g.Settings.Engine.Start(a.User), g.Settings.Timeout)
	c.set(key, g)
	go c.runner(g)
	c.Settings.Logger.New().Add("Key", key).Info("engine: new game started")
	return g
}
func (c *Cache) runner(g *T) {
	g.Settings.Engine.Run(g)
	c.set(g.ID(), nil)
}

func (c *Cache) Get(id string) *T {
	return c.games[id]
}

func (c *Cache) set(id string, g *T) {
	c.gameμ.Lock()
	if g == nil {
		delete(c.games, id)
	} else {
		c.games[id] = g
	}
	c.gameμ.Unlock()
}

func (c *Cache) Forget(id string) {
	c.gameμ.Lock()
	delete(c.games, id)
	c.gameμ.Unlock()
}

func (c *Cache) FindUsername(name string) *T {
	c.gameμ.Lock()
	defer c.gameμ.Unlock()
	for _, g := range c.games {
		for _, s := range g.Seats {
			if name == s.Username {
				return g
			}
		}
	}
	return nil
}

func (c *Cache) Search(deck *deck.T) *Search {
	if c.FindUsername(deck.User) != nil {
		return nil
	}
	c.srchμ.Lock()
	defer c.srchμ.Unlock()
	for _, search := range c.search {
		if search.Deck.User == deck.User {
			return search
		}
	}
	search := NewSearch(deck)
	c.search = append(c.search, search)
	return search
}

// func (c *Cache) Trigger(g *T, seat *game.Seat, token *game.Token, name string, arg interface{}) []game.Stater {
// 	log := g.Log().With(cast.JSON{
// 		"Seat":    seat.Username,
// 		"Token":   token.String(),
// 		"Trigger": name,
// 	})
// 	powers := token.Powers.GetTrigger(name)
// 	if len(powers) < 1 {
// 		log.Debug("empty")
// 		return nil
// 	}
// 	events := make([]game.Stater, 0)
// 	for _, p := range powers {
// 		if p.Target != "self" && arg == nil {
// 			power := p
// 			events = append(events, g.Runtime.Engine.Target(
// 				seat,
// 				p.Target,
// 				p.Text,
// 				func(val string) []game.Stater {
// 					return trigger.Power(g, seat, power, token, cast.NewArray(val))
// 				},
// 			))
// 		} else {
// 			events = append(events, state.NewTrigger(seat.Username, token, p, arg))
// 		}
// 	}
// 	log.Add("Events", events).Source().Debug()
// 	return events
// }

func (c *Cache) match(search0, search1 *Search) bool {
	return true
}

func (c *Cache) watchQueue() {
	i := 0
	var search0, searchi *Search
	for {
		<-time.After(time.Second)
		if len(c.search) < 2 {
			continue
		}
		c.srchμ.Lock()
		search0 = c.search[0]
		for i = 1; i < len(c.search); i++ {
			searchi = c.search[i]
			if c.match(search0, searchi) {
				c.New(search0.Deck, searchi.Deck)
				break
			}
		}
		c.srchμ.Unlock()
	}
}
