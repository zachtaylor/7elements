package main

import (
	"net/http"
	"time"

	"github.com/zachtaylor/7elements/db"
	_ "github.com/zachtaylor/7elements/scripts"
	"github.com/zachtaylor/7elements/server"
	"ztaylor.me/env"
	"ztaylor.me/http/sessions"
	"ztaylor.me/log"
)

const Patch = 2

func main() {
	env := env.Global()

	log.SetLevel(env.Default("LOG_LEVEL", "info"))

	fs := http.Dir(env.Default("WWW_PATH", "www/"))

	if logPath := env.Get("LOG_PATH"); logPath != "" {
		log.StartRoller(logPath)
	}

	if patch, err := db.OpenEnv(env); err != nil {
		log.Add("Error", err).Add("Patch", patch).Error("patch error")
		return
	} else if patch != Patch {
		log.Add("Expected", Patch).Add("Found", patch).Error("patch mismatch")
		return
	}

	if envName := env.Get("ENV"); envName == "dev" {
		go sessions.Service.Watch(1 * time.Minute)
		server.Start(fs, env.Get("DB_PWSALT"), ":"+env.Default("PORT", "80"))
	} else if envName == "pro" {
		go sessions.Service.Watch(6 * time.Hour)
		server.StartTLS(fs, env.Get("DB_PWSALT"), "7elements.cert", "7elements.key")
	} else {
		log.Error("7elements failed to launch, env error")
	}
}
