package gameserver

import (
	"github.com/zachtaylor/7elements/account"
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/chat"
	"taylz.io/db"
	"taylz.io/log"
)

type Settings struct {
	Accounts *account.Cache
	Cards    card.Prototypes
	Chats    *chat.Manager
	DB       *db.DB
	Logger   *log.T
}

func NewSettings(accounts *account.Cache, cards card.Prototypes, chats *chat.Manager, db *db.DB, logger *log.T) Settings {
	return Settings{
		Accounts: accounts,
		Cards:    cards,
		Chats:    chats,
		DB:       db,
		Logger:   logger,
	}
}

// // NewCacheSettings creates a new CacheSettings
// func NewCacheSettings(cards card.Prototypes, chats chat.Service, db *db.DB, engine Engine, keygen keygen.I, logger log.T, sockets *websocket.Cache) Settings {
// 	return CacheSettings{
// 		Cards:   cards,
// 		Chats:   chats,
// 		DB:      db,
// 		Engine:  engine,
// 		Keygen:  keygen,
// 		Logger:  logger,
// 		Sockets: sockets,
// 	}
// }
