package main

import (
	"7elements.ztaylor.me"
	"7elements.ztaylor.me/log"
	"7elements.ztaylor.me/options"
	"7elements.ztaylor.me/persistence"
	"7elements.ztaylor.me/server"
	"7elements.ztaylor.me/server/routes/openpack"
	"7elements.ztaylor.me/server/sessionman"
	_ "7elements.ztaylor.me/triggers"
	"net/http"
)

const patch uint64 = 13

func main() {
	go sessionman.SessionClock()

	log.Add("Patch", patch).Add("Patch-path", options.String("patch-path")).Info("starting 7Elements server...")

	persistence.Patch(options.String("patch-path"))
	if dbPatch, err := persistence.GetPatch(); err != nil {
		log.Add("patch-path", options.String("patch-path")).Add("Error", err).Error("patch read error")
		return
	} else if patch != dbPatch {
		log.Add("Expected", patch).Add("Found", dbPatch).Error("patch mismatch")
		return
	}

	if err := SE.Cards.LoadCache(); err != nil {
		log.Add("Error", err).Error("cannot load card cache, aborting...")
		return
	} else if err := SE.CardTexts.LoadCache("en-US"); err != nil {
		log.Add("Error", err).Error("cannot load card texts cache, aborting...")
		return
	}

	openpack.ShufflePacks()

	log.Add("port", options.String("port")).Add("server-path", options.String("server-path")).Info("7Elements server ready!")

	if options.Bool("use-https") {
		log.Error(http.ListenAndServeTLS(":443", "7elements.cert", "7elements.key", &server.Router))
	} else {
		log.Error(http.ListenAndServe(":"+options.String("port"), &server.Router))
	}
}
