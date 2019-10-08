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
	dbe "ztaylor.me/db/env"
	"ztaylor.me/db/mysql"
	"ztaylor.me/env"
	"ztaylor.me/http/session"
	"ztaylor.me/log"
)

const Patch = 2

var stdout = log.StdOutService(log.LevelDebug)

func main() {
	env := env.Global()
	envName := env.Get("ENV")

	rt := &api.Runtime{
		Root:       &vii.Runtime{},
		Chat:       chat.MemService{},
		Salt:       env.Get("DB_PWSALT"),
		FileSystem: http.Dir(env.Default("WWW_PATH", "www/")),
	}
	rt.Games = engine.NewService(rt.Root, rt.Chat)

	if conn, err := mysql.Open(dbe.BuildDSN(env)); err != nil {
		stdout.New().Add("Error", err).Error("db error")
		return
	} else if patch, err := db.Patch(conn); patch != Patch {
		stdout.New().Add("Expected", Patch).Add("Found", patch).Add("Error", err).Error("patch mismatch")
		return
	} else {
		rt.Root.Accounts = db.NewAccountService(conn)
		rt.Root.AccountsCards = db.NewAccountCardService(conn)
		rt.Root.AccountsDecks = db.NewAccountDeckService(conn)
		rt.Root.Cards = db.NewCardService(conn)
		rt.Root.Decks = db.NewDeckService(conn)
		rt.Root.Packs = db.NewPackService(conn)
	}

	logLevel, _ := log.GetLevel(env.Default("LOG_LEVEL", "info"))

	if envName == "dev" {
		rt.Sessions = session.NewCache(7 * time.Minute)
		rt.Root.Logger = log.StdOutService(logLevel)
		server.Start(rt, ":"+env.Default("PORT", "80"))
	} else if envName == "pro" {
		rt.Sessions = session.NewCache(7 * time.Hour)
		rt.Root.Logger = log.DailyRollingService(logLevel, env.Default("LOG_PATH", "./"))
		server.StartTLS(rt, "7elements.cert", "7elements.key")
	} else {
		stdout.New().Error("7elements failed to launch, env error")
	}
}
