package api

import (
	"github.com/zachtaylor/7elements"
	"ztaylor.me/http"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func PingHandler(r *http.Request) error {
	decks, err := vii.DeckService.GetAll()
	if err != nil {
		log.Add("Error", err).Warn("api/ping.json: deckservice getall failed")
	}

	packs, err := vii.PackService.GetAll()
	if err != nil {
		log.Add("Error", err).Warn("api/ping.json: packservice getall failed")
	}

	r.WriteJson(js.Object{
		"cards":  AllCardsJson(),
		"packs":  packs.Json(),
		"decks":  decks.Json(),
		"online": http.SessionService.Count(),
	})

	return nil
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
