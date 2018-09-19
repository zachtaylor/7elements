package main

import (
	"net/http"

	"github.com/zachtaylor/7elements/db"
	_ "github.com/zachtaylor/7elements/scripts"
	"github.com/zachtaylor/7elements/server"
	"ztaylor.me/env"
	"ztaylor.me/log"
)

const PATCH = 2

func main() {
	log.SetLevel(env.Default("LOG_LEVEL", "info"))

	fs := http.Dir(env.Default("WWW_PATH", "www/"))

	if logPath := env.Get("LOG_PATH"); logPath != "" {
		log.StartRoller(logPath)
	}

	if patch, err := db.ConnPatch(); err != nil {
		log.Add("Error", err).Add("Patch", patch).Error("patch error")
		return
	} else if patch != PATCH {
		log.Add("Expected", PATCH).Add("Found", patch).Error("patch mismatch")
		return
	}

	if env.Name() == "dev" {
		// http.SessionLifetime = 1 * time.Minute
		server.Start(fs, ":"+env.Default("PORT", "80"))
	} else if env.Name() == "pro" {
		server.StartTLS(fs, "7elements.cert", "7elements.key")
	} else {
		log.Error("7elements failed to launch, env error")
	}
}
