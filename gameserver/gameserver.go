package gameserver

import (
	"os"

	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/engine"
	"taylz.io/keygen"
	"taylz.io/log"
)

type T struct {
	Settings Settings
	*Cache
}

func New(settings Settings, games *Cache) *T {
	return &T{
		Settings: settings,
		Cache:    games,
	}
}

func (t *T) New(rules game.Rules, syslog *log.T, entry1, entry2 *Entry) (g *game.T) {
	t.Sync(func(get CacheGetter, set CacheSetter) {
		var key string
		for {
			key = keygen.New(21)
			if get(key) != nil {
				// continue
			} else if _, err := os.Stat("log/game/" + key + ".log"); err != nil {
				break // err = file doesn't exist
			}
		}
		f, _ := os.OpenFile("log/game/"+key+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		logger := log.Lining(log.LevelLiner(log.LevelTrace, log.IOLiner(&log.ColorFormat{
			SrcFmt:  log.ClassicSourceFormatter(game.SourcePath()),
			TimeFmt: log.DefaultTimeFormatter(),
		}, f)))
		chat := t.Settings.Chats.New("game#"+key, 7)
		g = game.New(key, engine.New(), chat, logger, rules)
		set(key, g)
	})

	g.Register(entry1.Deck, entry1.Writer)
	g.Register(entry2.Deck, entry2.Writer)

	go engine.Run(syslog, g)
	t.Settings.Logger.New().Add("ID", g.ID()).Info("game started")

	return g
}
