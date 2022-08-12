package match

import (
	"errors"

	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/ai"
	"github.com/zachtaylor/7elements/game/engine"
)

var (
	ErrDuplicate = errors.New("queue exists")
	ErrNotFound  = errors.New("matching error")
	ErrMatchSync = errors.New("match sync error")
	ErrDeckID    = errors.New("deckid out of range")
	ErrOwner     = errors.New("owner must be self or system")
)

// Maker is a match manager
type Maker struct {
	server Server
	cache  *Cache
	live   map[game.Rules]string
}

func NewMaker(server Server) *Maker {
	return &Maker{
		server: server,
		cache:  NewCache(),
		live:   make(map[game.Rules]string),
	}
}

func (m *Maker) Make(rules game.Rules, q1, q2 *Queue) {
	game := m.server.Games().New(
		rules,
		engine.NewRunner(),
		game.NewEntry(q1.writer, q1.cards),
		game.NewEntry(q2.writer, q2.cards),
	)
	gameid := game.ID()
	q2.Resolve(gameid)
	q1.Resolve(gameid)
	return
}

func (m *Maker) Queue(user game.Writer, settings QueueSettings) (*Queue, error) {
	username := user.Name()
	var count card.Count
	if settings.Owner == username {
		if account := m.server.Accounts().Get(username); account != nil {
			if settings.DeckID > -1 && settings.DeckID < len(account.Decks) {
				count = account.Decks[settings.DeckID].Cards
			} else {
				return nil, ErrDeckID
			}
		} else {
			return nil, errors.New("account missing")
		}
	} else if settings.Owner == "vii" {
		if settings.DeckID > -1 && settings.DeckID < len(m.server.Content().Decks()) {
			count = m.server.Content().Decks()[settings.DeckID].Cards
		} else {
			return nil, ErrDeckID
		}
	} else {
		return nil, ErrOwner
	}
	rules := RulesFromSettings(settings)

	if err := game.VerifyRulesDeck(rules, count); err != nil {
		return nil, err
	}

	q := NewQueue(user, settings, count)
	m.cache.Set(username, q)

	go m.TryMatch(rules, q)

	return q, nil
}

func (m *Maker) TryMatch(rules game.Rules, q *Queue) {
	m.cache.Sync(func() {
		if matchID := m.live[rules]; matchID == "" {
			m.live[rules] = q.Writer().Name()
		} else if match := m.cache.Get(matchID); match == nil {
			m.live[rules] = q.Writer().Name()
		} else {
			delete(m.live, rules)
			go m.Make(rules, q, match)
		}
	})
}

func (m *Maker) VSAI(user game.Entry, settings QueueSettings) *game.G {
	rules := RulesFromSettings(settings)
	ai := ai.New("A.I.")
	aiEntry := ai.Entry(m.server.Content())
	g := m.server.Games().New(rules, engine.NewRunner(), user, aiEntry)
	ai.Connect(g)
	return g
}

// Get returns the active Queue for the given username
func (m *Maker) Get(username string) *Queue { return m.cache.Get(username) }

func (m *Maker) Cancel(name string) (ok bool) {
	m.cache.Sync(func() {
		if q := m.cache.Get(name); q != nil {
			ok = true
			rules := RulesFromSettings(q.settings)
			delete(m.live, rules)
		}
	})
	m.cache.Remove(name)
	return
}
