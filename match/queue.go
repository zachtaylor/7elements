package match

import (
	"sync"
	"time"

	"github.com/zachtaylor/7elements/deck"
	"github.com/zachtaylor/7elements/game"
	"taylz.io/http/user"
)

type Queue struct {
	start    time.Time
	writer   user.Writer
	settings QueueSettings
	result   string
	once     sync.Once
	done     chan bool
}

func NewQueue(writer user.Writer, settings QueueSettings) *Queue {
	return &Queue{
		start:    time.Now(),
		writer:   writer,
		settings: settings,
		done:     make(chan bool),
	}
}

func (q *Queue) Start() time.Time { return q.start }

func (q *Queue) Writer() user.Writer { return q.writer }

func (q *Queue) Settings() QueueSettings { return q.settings }

func (q *Queue) GetRules() game.Rules { return q.settings.Rules() }

func (q *Queue) Deck() *deck.Prototype { return q.settings.Deck }

// SyncGameID waits to receive the id
func (q *Queue) SyncGameID() string {
	<-q.done
	return q.result
}

// Resolve saves the result and closes the Queue
func (q *Queue) Resolve(gameid string) (ok bool) {
	q.once.Do(func() {
		ok = true
		q.result = gameid
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

func (q *Queue) Data() map[string]interface{} {
	return map[string]interface{}{
		"deckid": q.settings.Deck.ID,
		"owner":  q.settings.Deck.User,
		"hands":  q.settings.Hands,
		"speed":  q.settings.Speed,
		"timer":  int(time.Since(q.start).Seconds()),
	}
}

type QueueSettings struct {
	Deck  *deck.Prototype
	Hands string
	Speed string
}

func NewQueueSettings(deck *deck.Prototype, hands string, speed string) QueueSettings {
	return QueueSettings{
		Deck:  deck,
		Hands: hands,
		Speed: speed,
	}
}

func (s QueueSettings) Rules() game.Rules {
	rules := game.DefaultRules()
	switch s.Hands {
	case "small":
		rules.StartingHand = 3
	case "med":
		rules.StartingHand = 4
	case "large":
		rules.StartingHand = 5
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
