package main

import (
	"7elements.ztaylor.me/cards"
	"7elements.ztaylor.me/db"
	_ "7elements.ztaylor.me/games"
	"7elements.ztaylor.me/options"
	"7elements.ztaylor.me/server"
	_ "7elements.ztaylor.me/triggers"
	"net/http"
	"ztaylor.me/http/sessions"
	"ztaylor.me/log"
)

const patch uint64 = 9

func main() {
	go sessions.SessionClock()

	log.Add("Patch", patch).Add("Patch-path", options.String("patch-path")).Info("starting 7Elements server...")

	db.Patch(options.String("patch-path"))
	if dbPatch, err := db.GetPatch(); err != nil {
		log.Add("patch-path", options.String("patch-path")).Add("Error", err).Error("patch read error")
		return
	} else if patch != dbPatch {
		log.Add("Expected", patch).Add("Found", dbPatch).Error("patch mismatch")
		return
	}

	if err := cards.LoadCache(); err != nil {
		log.Add("Error", err).Error("cannot load card cache, aborting...")
		return
	} else if err := cards.LoadBodyCache(); err != nil {
		log.Add("Error", err).Error("cannot load card bodies cache, aborting...")
		return
	} else if err := cards.LoadTextsCache("en-US"); err != nil {
		log.Add("Error", err).Error("cannot load card texts cache, aborting...")
		return
	}

	log.Add("port", options.String("port")).Add("server-path", options.String("server-path")).Info("7Elements server ready!")

	if options.Bool("use-https") {
		go http.ListenAndServe(":80", http.HandlerFunc(redirectHttps))
		log.Error(http.ListenAndServeTLS(":443", "7elements.cert", "7elements.key", &server.Router))
	} else {
		log.Error(http.ListenAndServe(":"+options.String("port"), &server.Router))
	}
}

func redirectHttps(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req,
		"https://"+req.Host+req.URL.String(),
		http.StatusMovedPermanently)
}
