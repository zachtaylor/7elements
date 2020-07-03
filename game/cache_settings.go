package game

import (
	"github.com/zachtaylor/7elements/account"
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/chat"
	"ztaylor.me/http/websocket"
	"ztaylor.me/keygen"
	"ztaylor.me/log"
)

// CacheSettings controls Cache options
type CacheSettings struct {
	Accounts account.Service
	Cards    card.PrototypeService
	Chats    chat.Service
	Keygen   keygen.Keygener
	Engine   Engine
	Logger   log.Service
	Sockets  websocket.Service
}

// NewCacheSettings creates a new CacheSettings
func NewCacheSettings(accounts account.Service, cards card.PrototypeService, chats chat.Service, keygen keygen.Keygener, engine Engine, logger log.Service, sockets websocket.Service) CacheSettings {
	return CacheSettings{
		Accounts: accounts,
		Cards:    cards,
		Chats:    chats,
		Keygen:   keygen,
		Engine:   engine,
		Logger:   logger,
		Sockets:  sockets,
	}
}
