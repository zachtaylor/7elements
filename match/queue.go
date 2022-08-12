package match

import (
	"sync"
	"time"

	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
)

type Queue struct {
	start    time.Time
	writer   game.Writer
	settings QueueSettings
	cards    card.Count
	gameid   string
	once     sync.Once
	done     chan bool
}

func NewQueue(writer game.Writer, settings QueueSettings, cards card.Count) *Queue {
	return &Queue{
		start:    time.Now(),
		writer:   writer,
		settings: settings,
		cards:    cards,
		done:     make(chan bool),
	}
}

func (q *Queue) Start() time.Time { return q.start }

func (q *Queue) Writer() game.Writer { return q.writer }

func (q *Queue) Settings() QueueSettings { return q.settings }

// func (q *Queue) GetRules() game.Rules { return q.settings.Rules() }

// func (q *Queue) Deck() *deck.Prototype { return q.settings.Deck }

// Sync waits to receive gameID, playerID
func (q *Queue) Sync() string {
	<-q.done
	return q.gameid
}

// Resolve saves the result and closes the Queue
func (q *Queue) Resolve(gameid string) (ok bool) {
	q.once.Do(func() {
		ok = true
		q.gameid = gameid
		close(q.done)
	})
	return
}

// Cancel closes the Queue without setting results
func (q *Queue) Cancel() (ok bool) {
	q.once.Do(func() {
		ok = true
		close(q.done)
	})
	return
}

func (q *Queue) Data() map[string]any {
	return map[string]any{
		"owner":  q.settings.Owner,
		"deckid": q.settings.DeckID,
		"hands":  q.settings.Hands,
		"speed":  q.settings.Speed,
		"timer":  int(time.Since(q.start).Seconds()),
	}
}

type QueueSettings struct {
	Owner  string
	DeckID int
	Hands  string
	Speed  string
}

func NewQueueSettings(owner string, deckID int, hands, speed string) QueueSettings {
	return QueueSettings{
		Owner:  owner,
		DeckID: deckID,
		Hands:  hands,
		Speed:  speed,
	}
}

func RulesFromSettings(s QueueSettings) game.Rules {
	rules := game.DefaultRules()
	switch s.Hands {
	case "small":
		rules.PlayerHand = 3
	case "med":
		rules.PlayerHand = 4
	case "large":
		rules.PlayerHand = 5
	}
	switch s.Speed {
	case "fast":
		rules.Timeout = 30 * time.Second
	case "med":
		rules.Timeout = 60 * time.Second
	case "slow":
		rules.Timeout = 90 * time.Second
	}
	return rules
}
