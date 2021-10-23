package game

import (
	"os"
	"strings"

	"github.com/zachtaylor/7elements/deck"
	"taylz.io/keygen"
	"taylz.io/log"
	"taylz.io/types"
)

type Manager struct {
	settings Settings
	cache    *Cache
}

func NewManager(settings Settings) *Manager {
	return &Manager{
		settings: settings,
		cache:    NewCache(),
	}
}

// func (m *Manager) Get(id string) *T

func (m *Manager) sourcePath() string {
	sourcePath := types.NewSource(0).File()
	sourcePath = sourcePath[:strings.LastIndex(sourcePath, "/")]
	return sourcePath
}

func (m *Manager) New(rules Rules, e1, e2 Enterer) (game *T) {
	m.cache.Sync(func(get CacheGetter, set CacheSetter) {
		var key, file string
		for {
			key = keygen.New(21)
			file = m.settings.LogDir + key + ".log"
			if get(key) != nil {
				// continue
			} else if _, err := os.Stat(file); err != nil {
				break // err = file doesn't exist
			}
		}
		f, _ := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		logger := log.Lining(log.LevelLiner(log.LevelTrace, log.IOLiner(&log.ColorFormat{
			SrcFmt:  log.ClassicSourceFormatter(m.sourcePath()),
			TimeFmt: log.DefaultTimeFormatter(),
		}, f)))
		// chat := t.Settings.Chats.New("game#"+key, 7)
		game = New(key, m.settings.Engine, rules, logger)
		set(key, game)
	})

	deck1, err := deck.BuildNew(e1.Writer().Name(), e1.Deck(), m.settings.Cards)
	if err != nil {
		return
	}
	deck2, err := deck.BuildNew(e2.Writer().Name(), e2.Deck(), m.settings.Cards)
	if err != nil {
		return
	}

	game.Register(deck1, e1.Writer())
	game.Register(deck2, e2.Writer())

	go Sandbox(m.settings.Logger, game)
	return
}

func (m *Manager) Get(id string) (game *T) { return m.cache.dat[id] }
