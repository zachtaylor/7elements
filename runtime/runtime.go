package runtime

import (
	"net/http"

	"github.com/zachtaylor/7elements/account"
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/card/pack"
	"github.com/zachtaylor/7elements/chat"
	db7 "github.com/zachtaylor/7elements/db"
	"github.com/zachtaylor/7elements/deck"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/player"
	"ztaylor.me/cast"
	"ztaylor.me/db"
	pkg_env "ztaylor.me/db/env"
	"ztaylor.me/db/mysql"
	"ztaylor.me/http/session"
	"ztaylor.me/log"
)

// T is a server runtime
type T struct {
	// IsDevEnv sets development environment
	IsDevEnv bool
	// PassSalt is the password salt
	PassSalt string

	Accounts   account.Service
	Cards      card.PrototypeService
	Chats      chat.Service
	Decks      deck.PrototypeService
	FileSystem http.FileSystem
	Games      *game.Cache
	Logger     log.Service
	Packs      pack.Service
	Players    *player.Cache
}

func New(env pkg_env.Service) (*T, error) {
	t := &T{}
	// logger
	if loglvl, _ := log.GetLevel(env["LOG_LEVEL"]); false {
		// no path
	} else if env["ENV"] == "dev" {
		t.Logger = log.StdOutService(loglvl, log.DefaultFormatWithColor())
		t.Logger.Format().CutPathSourceParent(1)
	} else {
		t.Logger = log.DailyRollingService(loglvl, log.DefaultFormatWithoutColor(), env["LOG_PATH"])
		t.Logger.Format().CutPathSourceParent(1)
	}
	// sessions
	if env["ENV"] == "dev" {
		t.Sessions = session.NewCache(1 * cast.Hour)
	} else {
		t.Sessions = session.NewCache(10 * cast.Minute)
	}
	// db
	if conn, err := mysql.Open(pkg_env.BuildDSN(env.Match("DB_"))); err != nil {
		return nil, &cast.Error{"db connection failed", err}
	} else if _, err = db.Patch(conn); err != nil {
		return nil, &cast.Error{"db patch failed", err}
	} else {
		t.Accounts = db7.NewAccountService(conn)
		// rt.AccountsCards = db.NewAccountCardService(conn)
		// rt.AccountsDecks = db.NewAccountDeckService(conn)
		t.Cards = db7.NewCardService(conn)
		t.Decks = db7.NewDeckService(conn)
		t.Packs = db7.NewPackService(conn)
	}

	// chat
	t.Chats = &chat.MemService{}
	// game engine
	t.Games = game.NewCache(game.NewCacheSettings(t.Accounts, t.Cards, t.Chats))
	// filesystem
	t.FileSystem = http.Dir(env["WWW_PATH"])
	// combined runtime
	t.Players = player.NewCache()
	// players.New(t.Logger, t.Accounts, websocket.NewCache(t.Sessions))

	return t, nil
}

func (t *T) Log() *log.Entry { return t.Settings.Logger.New() }

// func (t *T) AccountJSON(rt  a *account.T) cast.JSON {
// 	if a == nil {
// 		return nil
// 	} else if acs, err := rt.AccountsCards.Find(a.Username); err != nil {
// 		return nil
// 	} else if ads, err := rt.AccountsDecks.Find(a.Username); err != nil {
// 		return nil
// 	} else {
// 		return cast.JSON{
// 			"username": a.Username,
// 			"email":    a.Email,
// 			"session":  a.SessionID,
// 			"coins":    a.Coins,
// 			"cards":    acs.JSON(),
// 			"decks":    ads.JSON(),
// 		}
// 	}
// }
