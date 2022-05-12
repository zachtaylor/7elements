package server

import (
	"os"
	"strings"

	"github.com/zachtaylor/7elements/account"
	"github.com/zachtaylor/7elements/db"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/engine"
	"github.com/zachtaylor/7elements/match"
	"github.com/zachtaylor/7elements/server/internal"
	"github.com/zachtaylor/7elements/server/runtime"
	"taylz.io/http"
	"taylz.io/http/hash"
	"taylz.io/http/session"
	"taylz.io/http/user"
	"taylz.io/http/websocket"
	"taylz.io/keygen"
	"taylz.io/log"
	"taylz.io/types"
)

type T struct {
	runtime runtime.T
	fork    *http.Fork
	forkws  *websocket.MessageFork
}

func New(env map[string]string, ex_patch int) (*T, error) {
	var isprod bool
	if env["ENV"] == "prod" {
		isprod = true
	}

	// database
	db, err := db.OpenEnv(env, ex_patch)
	if err != nil {
		return nil, err
	}
	version, err := game.BuildVersion(db)
	if err != nil {
		return nil, err
	}

	passhash := hash.NewMD5SaltHash(env["DB_PWSALT"])

	// logging
	var logger *log.T
	filePath := types.NewSource(0).File()
	for i := 0; i < 3; i++ {
		filePath = filePath[:strings.LastIndex(filePath, "/")]
	}
	loglvl, err := log.GetLevel(env["LOG_LEVEL"])
	if err != nil {
		return nil, err
	} else if !isprod {
		logger = log.Lining(log.IOLiner(&log.ColorFormat{
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
		logger = log.Lining(log.LevelLiner(loglvl, log.IOLiner(&log.ColorFormat{
			SrcFmt:  log.ClassicSourceFormatter(filePath),
			TimeFmt: log.DefaultTimeFormatter(),
		}, log.DailyRotatingFile(env["LOG_PATH"]))))
	}

	// game deps
	gameRT := game.NewRuntime(db, engine.New())

	// server
	sessionSettings := session.DefaultSettings(keygen.NewFunc(8))
	if !isprod {
		sessionSettings.Lifetime = 1 * types.Minute
		sessionSettings.GC = 15 * types.Second
	} else {
		sessionSettings.Secure = true
		// sessionSettings.Lifetime = 1 * types.Hour // default is 1 hour
	}
	sessions := session.NewManager(sessionSettings)

	// wsHandler := &websocket.Fork{}

	sockets := websocket.NewManager(websocket.NewSettings(keygen.NewFunc(12)))

	users := user.NewManager(user.NewSettings(sessions, sockets))

	// chatKeygen := keygen.NewFunc(4)

	gameKeygen := keygen.NewFunc(21)
	games := game.NewManager(game.NewSettings(env["LOG_PATH"]+"game/", cards, logger, engine.New(), gameKeygen))

	matchMaker := match.NewMaker(match.NewSettings(logger, cards, decks, games))

	runtime := internal.NewRuntime(
		isprod,
		logger,
		db,
		glob,
		passhash,
		sessions,
		sockets,
		users,
		account.NewCache(),
		cards,
		decks,
		packs,
		games,
		matchMaker,
	)

	fork := &http.Fork{}

	server := T{
		runtime: runtime,
		fork:    fork,
		forkws:  &websocket.Fork{},
	}

	// todo observe various caches for logging and desyncing account data in time

	// todo  replace runtime/parse.go
}
