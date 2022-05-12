package runtime

import (
	"encoding/json"
	"errors"

	"github.com/zachtaylor/7elements/account"
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/card/pack"
	"github.com/zachtaylor/7elements/db/cards"
	"github.com/zachtaylor/7elements/db/decks"
	"github.com/zachtaylor/7elements/db/packs"
	"github.com/zachtaylor/7elements/deck"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/engine"
	"github.com/zachtaylor/7elements/match"
	"taylz.io/db"
	"taylz.io/http"
	"taylz.io/http/hash"
	"taylz.io/http/session"
	"taylz.io/http/user"
	"taylz.io/http/websocket"
	"taylz.io/keygen"
	"taylz.io/log"
	"taylz.io/types"
)

// T is a server runtime
type T struct {
	// EnvProd disables development environment
	EnvProd bool
	// FileSystem is the file-backed backup path resolver
	FileSystem http.FileSystem
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
	// Handler is http handler
	Handler *http.Fork
	// WSHandler is ws handler
	WSHandler *websocket.Fork
	// Sessions is *session.Manager
	Sessions *session.Manager
	// Sockets is *websocket.Manager
	Sockets *websocket.Manager
	// Users is *user.Manager
	Users *user.Manager
	// Chats is a chat.Service
	// Chats *chat.Manager
	// Games is a game.Server
	Games *game.Manager
	// MatchMaker is a *matchmaker.T
	MatchMaker *match.Maker
	// glob is cached global data
	glob types.Bytes
}

// New is the canonical way to create a runtime
func New(
	isprod bool,
	fs http.FileSystem,
	passhash hash.Func,
	db *db.DB,
	log *log.T,
	sessionSettings session.Settings,
	wsKeygen keygen.Func,
	chatRoomKeygen keygen.Func,
) (rt *T, err error) {

	rt = &T{
		EnvProd:    isprod,
		FileSystem: fs,
		PassHash:   passhash,
		DB:         db,
		Logger:     log,
		Handler:    &http.Fork{},
		WSHandler:  &websocket.Fork{},
	}

	if rt.Cards = cards.GetAll(db); rt.Cards == nil {
		return nil, errors.New("failed to load cards")
	} else if rt.Decks, err = decks.GetAll(db); rt.Decks == nil {
		return nil, err
	} else if rt.Packs, err = packs.GetAll(db); err != nil {
		return nil, err
	} else if rt.glob, err = json.Marshal(map[string]any{
		"cards": rt.Cards.Data(),
		"decks": rt.Decks.Data(),
		"packs": rt.Packs.Data(),
	}); err != nil {
		return nil, err
	}

	rt.Accounts = account.NewCache()

	rt.Sessions = session.NewManager(sessionSettings)

	rt.Sockets = websocket.NewManager(websocket.NewSettings(rt.Sessions, wsKeygen, rt.WSHandler))

	rt.Users = user.NewManager(user.NewSettings(rt.Sessions, rt.Sockets))

	// rt.Chats = chat.NewManager(chat.NewSettings(rt.Users, chatRoomKeygen))

	rt.Games = game.NewManager(game.NewSettings("./game/", rt.Cards, rt.Logger, engine.New(), keygen.NewFunc(21)))

	rt.MatchMaker = match.NewMaker(match.NewSettings(rt.Logger, rt.Cards, rt.Decks, rt.Games))

	rt.Sessions.Observe(rt.OnSession)
	rt.Sockets.Observe(rt.OnSocket)
	rt.Users.Observe(rt.OnUser)

	return
}

func (rt *T) GlobalData() []byte { return rt.glob }

func (rt *T) Log() log.Writer { return rt.Logger }

// func (rt *T) UpdateSession(id string) error {
// 	if s := rt.Sessions.Get(id); s != nil {
// 		s.Update()
// 	}
// 	return session.ErrExpired
// }
