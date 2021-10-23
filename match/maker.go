package match

import (
	"errors"

	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/ai"
	"taylz.io/http/user"
)

var (
	ErrDuplicate = errors.New("queue exists")
	ErrNotFound  = errors.New("matching error")
	ErrMatchSync = errors.New("match sync error")
)

// Maker is a match manager
type Maker struct {
	settings Settings
	cache    *Cache
	live     map[game.Rules]string
}

func NewMaker(settings Settings) *Maker {
	return &Maker{
		settings: settings,
		cache:    NewCache(),
		live:     make(map[game.Rules]string),
	}
}

func (m *Maker) Make(rules game.Rules, q1, q2 *Queue) (err error) {
	game := m.settings.Games.New(rules, q1, q2)
	gameid := game.ID()
	if !q2.Resolve(gameid) {
		err = ErrMatchSync
	} else if !q1.Resolve(gameid) {
		err = ErrMatchSync
	}
	return
}

func (m *Maker) Queue(user user.Writer, settings QueueSettings) (q *Queue, err error) {

	username := user.Name()
	rules := settings.Rules()

	// TODO 2nd player doesn't release the lock or what ?
	m.cache.Sync(func(get CacheGetter, set CacheSetter) {
		if q = get(username); q != nil {
			err = ErrDuplicate
			return
		}
		q = NewQueue(user, settings)
		if foundUsername := m.live[rules]; foundUsername == "" {
			err = ErrNotFound
		} else if foundWaiting := get(foundUsername); foundWaiting == nil {
			err = ErrNotFound
		} else if err = m.Make(rules, q, foundWaiting); err != nil {
			// err != nil
		} else {
			set(foundUsername, nil)
			delete(m.live, rules)
		}

		if err != nil {
			set(username, q)
			m.live[rules] = username
			err = nil
		}
	})

	return
}

func (m *Maker) VSAI(user user.Writer, settings QueueSettings) *game.T {
	return m.settings.Games.New(settings.Rules(), game.NewEntry(settings.Deck, user), ai.New("A.I.").Entry(m.settings.Logger, m.settings.Cards, m.settings.Decks))
}

// Get returns the active Queue for the given username
func (m *Maker) Get(username string) *Queue { return m.cache.Get(username) }

func (m *Maker) Cancel(name string) (err error) {
	m.cache.Sync(func(get CacheGetter, set CacheSetter) {
		if q := get(name); q == nil {
			err = ErrNotFound
		} else {
			set(name, nil)
			delete(m.live, q.GetRules())
		}
	})
	return
}
