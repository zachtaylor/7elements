package runtime

import (
	"math/rand"
	"net/http"

	"github.com/zachtaylor/7elements/db/accounts_decks"
	"github.com/zachtaylor/7elements/game/engine"

	"github.com/zachtaylor/7elements/account"
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/card/pack"
	"github.com/zachtaylor/7elements/chat"
	db7 "github.com/zachtaylor/7elements/db"
	"github.com/zachtaylor/7elements/deck"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/player"
	"ztaylor.me/cast"
	"ztaylor.me/cast/charset"
	"ztaylor.me/db"
	pkg_env "ztaylor.me/db/env"
	"ztaylor.me/db/mysql"
	"ztaylor.me/http/mux"
	"ztaylor.me/http/session"
	"ztaylor.me/http/websocket"
	"ztaylor.me/keygen"
	"ztaylor.me/log"
)

// T is a server runtime
type T struct {
	// IsDevEnv sets development environment
	IsDevEnv bool
	// PassSalt sets the password salt
	PassSalt string

	Logger     log.Service
	Sessions   session.Service
	Sockets    *websocket.Cache
	Accounts   account.Service
	Cards      card.PrototypeService
	Chats      chat.Service
	Decks      deck.PrototypeService
	Packs      pack.Service
	FileSystem http.FileSystem
	Games      *game.Cache
	Players    *player.Cache
	Server     *mux.Mux
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
	t.Sockets = websocket.NewCache(t.Sessions)
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
	t.Games = game.NewCache(game.NewCacheSettings(
		t.Accounts,
		t.Cards,
		t.Chats,
		&keygen.Settings{
			KeySize: 7,
			CharSet: charset.AlphaCapitalNumeric,
			Rand:    rand.New(rand.NewSource(cast.Now().UnixNano())),
		},
		engine.New(),
		t.Logger,
		t.Sockets,
	))
	// filesystem
	t.FileSystem = http.Dir(env["WWW_PATH"])
	// combined runtime
	t.Players = player.NewCache(t.Logger, t.Sessions, t.Sockets, t.Accounts)

	return t, nil
}

func (t *T) Log() *log.Entry { return t.Logger.New() }

func (t *T) GlobalJSON() cast.JSON {
	decks, _ := accounts_decks.Get("vii")
	packs, _ := t.Packs.GetAll()
	users, _ := t.Accounts.Count()
	return cast.JSON{
		"cards": t.Cards.GetAll().JSON(),
		"packs": packs.JSON(),
		"decks": decks.JSON(),
		"users": users,
	}
}

func (t *T) AccountJSON(a *account.T) cast.JSON {
	if a == nil {
		return nil
	}
	return cast.JSON{
		"username": a.Username,
		"email":    a.Email,
		"session":  a.SessionID,
		"coins":    a.Coins,
		"cards":    a.Cards.JSON(),
		"decks":    a.Decks.JSON(),
	}
}
