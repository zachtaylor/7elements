package runtime

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/zachtaylor/7elements/account"
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/db"
	"github.com/zachtaylor/7elements/db/cards"
	"github.com/zachtaylor/7elements/db/decks"
	"github.com/zachtaylor/7elements/db/packs"
	"github.com/zachtaylor/7elements/gameserver"
	"github.com/zachtaylor/7elements/gameserver/queue"
	"taylz.io/db/patch"
	"taylz.io/env"
	"taylz.io/http"
	"taylz.io/http/hash"
	"taylz.io/http/session"
	"taylz.io/http/user"
	"taylz.io/http/websocket"
	"taylz.io/keygen"
	"taylz.io/log"
	"taylz.io/types"
)

func Parse(env env.Service, ex_patch int) (*T, error) {
	t := &T{}
	// isdevenv
	if env["ENV"] == "dev" {
		t.IsDevEnv = true
	}
	// passhash
	t.PassHash = hash.NewMD5SaltHash(env["DB_PWSALT"])
	// db
	if conn, err := db.OpenEnv(env); err != nil {
		return nil, types.WrapErr(err, "db connection failed")
	} else if patch, err := patch.Get(conn); err != nil {
		return nil, types.WrapErr(err, "db patch failed")
	} else if patch != ex_patch {
		return nil, types.NewErr("Patch mismatch: " + strconv.FormatInt(int64(patch), 10))
	} else {
		t.DB = conn
	}
	// logger
	filePath := types.NewSource(0).File()
	for i := 0; i < 3; i++ {
		filePath = filePath[:strings.LastIndex(filePath, "/")]
	}
	loglvl, err := log.GetLevel(env["LOG_LEVEL"])
	if err != nil {
		return nil, err
	} else if t.IsDevEnv {
		t.Logger = log.Lining(log.IOLiner(&log.ColorFormat{
			Colors:     log.DefaultColorMap(),
			ColorMsg:   true,
			ColorField: true,
			SrcFmt: log.RestringSourceFormatter(
				log.DetailSourceFormatter(),
				log.RestringerMiddleware(log.RestringerCutPrefixes{filePath}, log.RestringerLenExact(40)),
			),
			TimeFmt: log.DefaultTimeFormatter(),
		}, os.Stdout))
	} else {
		t.Logger = log.Lining(log.LevelLiner(loglvl, log.IOLiner(&log.ColorFormat{
			SrcFmt:  log.ClassicSourceFormatter(filePath),
			TimeFmt: log.DefaultTimeFormatter(),
		}, log.DailyRotatingFile(env["LOG_PATH"]))))
	}
	// accounts
	t.Accounts = account.NewCache()
	// Cards
	if t.Cards = cards.GetAll(t.DB); t.Cards == nil {
		return nil, errors.New("failed to load cards")
	}
	if t.Decks, err = decks.GetAll(t.DB); t.Decks == nil {
		return nil, types.WrapErr(err, "failed to load decks")
	}
	// Packs
	packs, err := packs.GetAll(t.DB)
	if err != nil {
		return nil, types.WrapErr(err, "failed to load packs")
	}
	t.Packs = packs
	t.GlobalData() // build glob
	// server
	t.Server = &http.Fork{}
	t.WSServer = &websocket.Fork{}
	sessionSettings := session.DefaultSettings(keygen.NewFunc(8))
	if t.IsDevEnv {
		sessionSettings.Lifetime = 10 * types.Minute
	} else {
		sessionSettings.Secure = true
		sessionSettings.Lifetime = 1 * types.Hour
	}
	t.Sessions = session.NewManager(sessionSettings)
	t.Sockets = websocket.NewManager(websocket.NewSettings(t.Sessions, keygen.NewFunc(12), t.WSServer))
	t.Users = user.NewManager(user.Settings{
		Sessions: t.Sessions,
		Sockets:  t.Sockets,
	})
	t.Sessions.Observe(t.OnSession)
	t.Users.Observe(t.OnUser)
	t.Sockets.Observe(t.OnWebsocket)
	// chats
	t.Chats = chat.NewManager(chat.Settings{
		Users:  t.Users,
		Keygen: keygen.NewFunc(4),
	})
	// game engine
	t.Games = gameserver.New(gameserver.Settings{
		Accounts: t.Accounts,
		Cards:    t.Cards,
		Chats:    t.Chats,
		DB:       t.DB,
		Logger:   t.Logger,
	}, gameserver.NewCache())
	// queue
	t.Queue = queue.NewServer(queue.Settings{
		Games:  t.Games,
		Logger: t.Logger,
		// Users:  t.Server.Users,
	}, queue.NewCache())

	return t, nil
}
