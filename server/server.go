package server

import (
	"os"
	"strings"
	"time"

	"github.com/zachtaylor/7elements/db"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/internal/user"
	"github.com/zachtaylor/7elements/server/internal"
	"taylz.io/http/hash"
	"taylz.io/http/session"
	"taylz.io/http/websocket"
	"taylz.io/log"
	"taylz.io/types"
)

func NewRuntime(env map[string]string, ex_patch int) (internal.Server, error) {
	var isprod bool
	if env["ENV"] == "prod" {
		isprod = true
	}

	// database
	db, err := db.OpenEnv(env, ex_patch)
	if err != nil {
		return nil, err
	}

	passhash := hash.NewMD5SaltHash(env["DB_PWSALT"])

	// logging
	var logger *log.T
	filePath := types.NewSource(0).File()
	for i := 0; i < 2; i++ {
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

	// server
	sessionSettings := session.DefaultSettings()
	if !isprod {
		sessionSettings.Lifetime = 10 * time.Minute
		sessionSettings.GC = 15 * time.Second
	} else {
		sessionSettings.Secure = true
		// sessionSettings.Lifetime = 1 * types.Hour // default is 1 hour
	}

	origins := strings.Split(env["ORIGINS"], ",")

	websocketSettings := websocket.Settings{
		OriginPatterns: origins,
	}

	userSettings := user.Settings{
		ReadSpeedLimit: time.Second,
		Websocket:      websocketSettings,
	}

	gameSettings := game.NewSettings(env["LOG_DIR"])

	return internal.NewRuntime(
		isprod,
		db,
		passhash,
		logger,
		sessionSettings,
		userSettings,
		gameSettings,
		origins,
		Routes,
		WSRoutes(),
	)
}
