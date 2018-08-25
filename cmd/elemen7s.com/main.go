package main

import (
	"github.com/zachtaylor/7elements/db"
	_ "github.com/zachtaylor/7elements/scripts"
	"github.com/zachtaylor/7elements/server"
	"ztaylor.me/env"
	"ztaylor.me/log"
)

const PATCH = 1

func main() {
	log.SetLevel(env.Default("LOG_LEVEL", "info"))

	if logPath := env.Get("LOG_PATH"); logPath != "" {
		log.StartRoller(logPath)
	}

	if patch, err := db.Patch(); err != nil {
		log.Add("Error", err).Add("Patch", patch).Error("patch error")
		return
	} else if patch != PATCH {
		log.Add("Expected", PATCH).Add("Found", patch).Error("patch mismatch")
		return
	}

	if err := vii.CardService.Start(); err != nil {
		log.Add("Error", err).Error("cannot load card cache, aborting...")
		return
	} else if vii.CardTextService == nil {
		log.Error("vii.CardTextService must not be nil")
		return
	} else if err := vii.CardTextService.Start(); err != nil {
		log.Add("Error", err).Error("cannot load card texts cache, aborting...")
		return
	}

	if env.Name() == "dev" {
		// http.SessionLifetime = 1 * time.Minute
		server.Start(":" + env.Default("PORT", "80"))
	} else if env.Name() == "pro" {
		server.StartTLS("7elements.cert", "7elements.key")
	} else {
		log.Add("env", env.Name()).Error("7elements failed to launch, env error")
	}
}
