package api

import (
	"net/http"

	"github.com/zachtaylor/7elements"
	"ztaylor.me/http/sessions"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	decks, err := vii.DeckService.GetAll()
	if err != nil {
		log.Add("Error", err).Warn("api/ping.json: deckservice getall failed")
	}

	packs, err := vii.PackService.GetAll()
	if err != nil {
		log.Add("Error", err).Warn("api/ping.json: packservice getall failed")
	}

	w.Write([]byte(js.Object{
		"cards":  AllCardsJson(),
		"packs":  packs.Json(),
		"decks":  decks.Json(),
		"online": sessions.Service.Count(),
	}.String()))
}

// func pingHandlerDataHelperGames(username string) js.Object {
// 	gamesdata := js.Object{}
// 	games := vii.GameService.GetPlayerGames(username)
// 	for _, gameid := range games {
// 		game := vii.GameService.Get(gameid)
// 		gamesdata[gameid] = game.Json(username)
// 	}
// 	return gamesdata
// }
