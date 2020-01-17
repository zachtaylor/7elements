package main

import (
	"net/http"
	"time"

	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/db"
	"github.com/zachtaylor/7elements/game/engine"
	_ "github.com/zachtaylor/7elements/scripts"
	"github.com/zachtaylor/7elements/server"
	"github.com/zachtaylor/7elements/server/api"
	"ztaylor.me/cast"
	dbe "ztaylor.me/db/env"
	"ztaylor.me/db/mysql"
	"ztaylor.me/env"
	"ztaylor.me/http/session"
	"ztaylor.me/log"
)

const Patch = 3

var stdout = log.StdOutService(log.LevelDebug)

func main() {
	env := env.Global()
	envName := env.Get("ENV")

	stdout.Formatter().CutSourcePath(2)
	stdout.New().With(cast.JSON{
		"Patch":   Patch,
		"ENV":     envName,
		"DB_USER": env.Get("DB_USER"),
	}).Source().Debug()
	tstart := cast.Now()

	rt := &api.Runtime{
		Root:       &vii.Runtime{},
		Chat:       chat.MemService{},
		Salt:       env.Get("DB_PWSALT"),
		FileSystem: http.Dir(env.Default("WWW_PATH", "www/")),
		Ping:       &api.Ping{},
	}
	// games is a circular ref
	rt.Games = engine.NewService(rt.Root, rt.Chat)

	if conn, err := mysql.Open(dbe.BuildDSN(env)); err != nil {
		stdout.New().Add("Error", err).Error("db error")
		return
	} else if patch, err := db.Patch(conn); patch != Patch {
		stdout.New().With(cast.JSON{
			"ExpectedPatch": Patch,
			"ActualPatch":   patch,
			"Error":         err,
		}).Source().Error("patch error")
		return
	} else {
		rt.Root.Accounts = db.NewAccountService(conn)
		rt.Root.AccountsCards = db.NewAccountCardService(conn)
		rt.Root.AccountsDecks = db.NewAccountDeckService(conn)
		rt.Root.Cards = db.NewCardService(conn)
		rt.Root.Decks = db.NewDeckService(conn)
		rt.Root.Packs = db.NewPackService(conn)
		stdout.New().
			Add("Cards", len(rt.Root.Cards.GetAll())).
			Add("Speed", cast.Now().Sub(tstart)).
			Source().Debug("loaded")
	}

	logLevel, _ := log.GetLevel(env.Default("LOG_LEVEL", "info"))

	if envName == "dev" {
		rt.Sessions = session.NewCache(7 * time.Minute)
		rt.Root.Logger = stdout
		server.Start(rt, ":"+env.Default("PORT", "80"))
	} else if envName == "pro" {
		rt.Sessions = session.NewCache(7 * time.Hour)
		rt.Root.Logger = log.DailyRollingService(logLevel, env.Default("LOG_PATH", "./"))
		server.StartTLS(rt, "7elements.cert", "7elements.key")
	} else {
		stdout.New().Error("7elements failed to launch, env error")
	}
}
