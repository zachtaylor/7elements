package api

import (
	"fmt"
	"time"

	"github.com/zachtaylor/7elements"
	"ztaylor.me/http"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func PingHandler(r *http.Request) error {
	if r.Session == nil {
		log.Info("/ping")
		r.WriteJson(js.Object{
			"uri": "/ping",
			"data": js.Object{
				"online": http.SessionCount(),
				"cards":  AllCardsJson("en-US"),
			},
		})
		log.Info("/ping")
		return nil
	} else if account, err := vii.AccountService.Get(r.Username); account == nil {
		return err
	} else if decks, err := vii.AccountDeckService.Get(r.Username); decks == nil {
		return err
	} else if accountcards, err := vii.AccountCardService.Get(r.Username); accountcards == nil {
		return err
	} else {
		cardsdata := AllCardsJson("en-US")
		for cardid, cards := range accountcards {
			if k := fmt.Sprintf("%d", cardid); cardsdata[k] == nil {
				log.WithFields(log.Fields{
					"Key":      k,
					"CardId":   cardid,
					"Username": r.Username,
					"Copies":   len(cards),
				}).Warn("/ping: account contains copies of missing card")
			} else {
				cardsdata[k].(js.Object)["copies"] = len(cards)
			}
		}

		r.Session.Refresh()
		r.WriteJson(js.Object{
			"uri": "/ping",
			"data": js.Object{
				"username":       r.Username,
				"email":          account.Email,
				"session-expire": r.Session.Expire.Sub(time.Now()).String(),
				"coins":          account.Coins,
				"packs":          account.Packs,
				"cards":          cardsdata,
				"decks":          decks.Json(),
				"online":         http.SessionCount(),
				"games":          pingHandlerDataHelperGames(r.Username),
			},
		})
		return nil
	}
}

func pingHandlerDataHelperGames(username string) js.Object {
	gamesdata := js.Object{}
	games := vii.GameService.GetPlayerGames(username)
	for _, gameid := range games {
		game := vii.GameService.Get(gameid)
		gamesdata[gameid] = game.Json(username)
	}
	return gamesdata
}
