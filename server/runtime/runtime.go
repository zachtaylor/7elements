package runtime

import (
	"github.com/zachtaylor/7elements/account"
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/card/pack"
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/deck"
	"github.com/zachtaylor/7elements/gameserver"
	"github.com/zachtaylor/7elements/gameserver/queue"
	"taylz.io/db"
	"taylz.io/http"
	"taylz.io/http/hash"
	"taylz.io/http/session"
	"taylz.io/http/user"
	"taylz.io/http/websocket"
	"taylz.io/log"
	"taylz.io/types"
)

// T is a server runtime
type T struct {
	// IsDevEnv sets development environment
	IsDevEnv bool
	// PassHash hashes passwords
	PassHash hash.Func
	// DB is database connection
	DB *db.DB
	// Logger is system log util
	Logger *log.T
	// Accounts stores accounts
	Accounts *account.Cache
	// Cards stores cards
	Cards card.Prototypes
	// Decks stores decks
	Decks deck.Prototypes
	// Packs stores packs
	Packs pack.Prototypes
	// Chats is a chat.Service
	Chats *chat.Manager
	// Games is a game.Server
	Games *gameserver.T
	// Server is handler
	Server *http.Fork
	// WSServer is ws handler
	WSServer *websocket.Fork
	// Sessions is *session.Manager
	Sessions *session.Manager
	// Sockets is *websocket.Manager
	Sockets *websocket.Manager
	// Users is *user.Manager
	Users *user.Manager
	// Queue is a queue.Server
	Queue *queue.Server
	// glob is cached global data
	glob types.Bytes
}

func (t *T) Log() log.Writer { return t.Logger.New() }

// func (t *T) AccountJSON(a *account.T) websocket.MsgData {
// 	if a == nil {
// 		return nil
// 	}
// 	return websocket.MsgData{
// 		"username": a.Username,
// 		"email":    a.Email,
// 		"session":  a.SessionID,
// 		"coins":    a.Coins,
// 		"cards":    a.Cards.JSON(),
// 		"decks":    a.Decks.JSON(),
// 	}
// }
