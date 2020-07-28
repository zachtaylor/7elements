package game

import (
	"time"

	"github.com/zachtaylor/7elements/account"
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/chat"
	"ztaylor.me/cast"
	"ztaylor.me/http/websocket"
	"ztaylor.me/log"
)

// Settings is game configuration
type Settings struct {
	Engine   Engine
	Timeout  time.Duration
	Accounts account.Service
	Cards    card.PrototypeService
	Chat     *chat.Room
	Logger   log.Service
	Sockets  *websocket.Cache
}

// NewSettings creates a Settings struct
func NewSettings(e Engine, timeout time.Duration, a account.Service, c card.PrototypeService, chat *chat.Room, logWriter cast.WriteCloser, s *websocket.Cache) Settings {
	logger := log.NewService(log.LevelDebug, log.DefaultFormatWithoutColor(), logWriter)
	logger.Format().CutPathSourceParent(1)
	return Settings{
		Engine:   e,
		Timeout:  timeout,
		Accounts: a,
		Cards:    c,
		Chat:     chat,
		Logger:   logger,
		Sockets:  s,
	}
}
