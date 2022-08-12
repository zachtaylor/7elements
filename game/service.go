package game

import (
	"os"
	"strings"
	"sync"

	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/content"
	"taylz.io/keygen"
	"taylz.io/log"
	"taylz.io/types"
)

type Service struct {
	Settings
	cache  *Cache
	sync   sync.Mutex
	keygen func() string
	server Server
}

func NewService(settings Settings, keygen func() string, server Server) *Service {
	return &Service{
		Settings: settings,
		cache:    NewCache(),
		keygen:   keygen,
		server:   server,
	}
}

// func (m *Service) Get(id string) *T

func (m *Service) sourcePath() string {
	sourcePath := types.NewSource(0).File()
	sourcePath = sourcePath[:strings.LastIndex(sourcePath, "/")]
	return sourcePath
}

func (m *Service) New(rules Rules, runner Runner, e1, e2 Entry) (game *G) {
	m.sync.Lock()
	defer m.sync.Unlock()
	var key, file string
	for {
		key = m.keygen()
		file = m.LogDir + "game/" + key + ".log"
		if m.Get(key) != nil {
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
	game = New(key, rules, keygen.DefaultFunc(), logger)
	m.cache.Set(key, game)

	m.NewPlayer(game, e1.Writer, e1.cardCount, m.server.Content())
	m.NewPlayer(game, e2.Writer, e2.cardCount, m.server.Content())

	go Sandbox(game, m.server.Log(), runner)
	return
}

func (m *Service) NewPlayer(game *G, w Writer, cardCount card.Count, content content.T) (player *Player) {
	player = game.NewPlayer(NewPlayerContext(w, game.Rules().PlayerLife))
	future := make([]string, 0, cardCount.Count())
	cards := content.Cards()
	for cardID, count := range cardCount {
		if proto := cards[cardID]; proto != nil {
			for i := 0; i < count; i++ {
				card := game.NewCard(player.id, proto)
				future = append(future, card.id)
			}
		}
	}
	player.T.Future = future
	return player
}

func (m *Service) Get(id string) (game *G) { return m.cache.Get(id) }
