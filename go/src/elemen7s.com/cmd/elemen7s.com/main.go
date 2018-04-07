package main

import (
	"elemen7s.com"
	"elemen7s.com/db"
	_ "elemen7s.com/db/service"
	_ "elemen7s.com/games"
	_ "elemen7s.com/scripts"
	"elemen7s.com/server"
	"ztaylor.me/buildir/js"
	"ztaylor.me/env"
	"ztaylor.me/log"
)

const PATCH = 19

func main() {
	env.Bootstrap()

	log.SetLevel(env.Default("LOG_LEVEL", "info"))

	db.Open(env.Default("DB_PATH", "elemen7s.db"))

	if patch := db.Patch(); patch != PATCH {
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

	log.StartRoller(env.Default("LOG_PATH", "log/"))

	if env.Name() == "dev" {
		// http.SessionLifetime = 1 * time.Minute
		server.Start(":" + env.Default("PORT", "80"))
	} else if env.Name() == "elemen7s.com" {
		buildirjs.EnableMinify()
		server.StartTLS("7elements.cert", "7elements.key")
	} else {
		log.Add("env", env.Name()).Error("7elements failed to launch, env error")
	}
}
