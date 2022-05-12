package internal

import (
	"github.com/zachtaylor/7elements/content"
	"taylz.io/log"
)

// import (
// 	"github.com/zachtaylor/7elements/account"
// 	"github.com/zachtaylor/7elements/card"
// 	"github.com/zachtaylor/7elements/card/pack"
// 	"github.com/zachtaylor/7elements/db/accounts"
// 	"github.com/zachtaylor/7elements/deck"
// 	"github.com/zachtaylor/7elements/game"
// 	"github.com/zachtaylor/7elements/match"
// 	"github.com/zachtaylor/7elements/wsout"
// 	"taylz.io/db"
// 	"taylz.io/http/session"
// 	"taylz.io/http/user"
// 	"taylz.io/http/websocket"
// )

type Server interface {
	Content() content.T
	Log() log.Writer
}

// Server is the interally exposed server api
// type Server interface {
// 	EnvProd() bool
// 	Log() log.Writer
// 	GetDB() *db.DB
// 	GetFileSystem() http.FileSystem
// 	GetAccounts() *account.Cache
// 	GetSessionManager() session.Manager
// 	// GetWebsocketManager() *websocket.Manager
// 	GetUserManager() user.Manager
// 	HashPassword(string) string
// 	GetGameVersion() *game.Version
// 	GetGameManager() *game.Manager
// 	GetMatchMaker() *match.Maker
// 	Ping()
// }

// // Runtime is the internal-facing API routes consume
// type Runtime interface {
// 	EnvProd() bool
// 	Log() log.Writer
// 	GetDB() *db.DB
// 	GetSessions() *session.Manager
// 	GetSockets() *websocket.Manager
// 	GetUsers() *user.Manager
// 	GetAccounts() *account.Cache
// 	GetCards() card.Prototypes
// 	GetDecks() deck.Prototypes
// 	GetPacks() pack.Prototypes
// 	GetGames() *game.Manager
// 	GetMatchMaker() *match.Maker
// }

// type runtime struct {
// 	envProd    bool
// 	log        log.Writer
// 	db         *db.DB
// 	patchData  []byte
// 	passHash   func(string) string
// 	sessions   *session.Manager
// 	sockets    *websocket.Manager
// 	users      *user.Manager
// 	accounts   *account.Cache
// 	cards      card.Prototypes
// 	decks      deck.Prototypes
// 	packs      pack.Prototypes
// 	games      *game.Manager
// 	matchMaker *match.Maker
// }

// func NewRuntime(
// 	isprod bool,
// 	log log.Writer,
// 	db *db.DB,
// 	patchData []byte,
// 	passHash func(string) string,
// 	sessions *session.Manager,
// 	sockets *websocket.Manager,
// 	users *user.Manager,
// 	accounts *account.Cache,
// 	cards card.Prototypes,
// 	decks deck.Prototypes,
// 	packs pack.Prototypes,
// 	games *game.Manager,
// 	matchMaker *match.Maker,
// ) runtime {
// 	return runtime{
// 		envProd:    isprod,
// 		log:        log,
// 		db:         db,
// 		patchData:  patchData,
// 		passHash:   passHash,
// 		sessions:   sessions,
// 		sockets:    sockets,
// 		users:      users,
// 		accounts:   accounts,
// 		cards:      cards,
// 		decks:      decks,
// 		packs:      packs,
// 		games:      games,
// 		matchMaker: matchMaker,
// 	}
// }

// // Ping sends updated user counts
// func (rt runtime) Ping() {
// 	users, _ := accounts.Count(rt.db)
// 	bytes := wsout.Ping(map[string]any{
// 		"ping":   rt.GetSockets().Count(),
// 		"online": rt.GetSessions().Count(),
// 		"users":  users,
// 	})
// 	rt.GetSockets().Each(func(id string, ws *websocket.T) {
// 		ws.Write(bytes)
// 	})
// }
