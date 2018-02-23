package api

import (
	"elemen7s.com/accounts"
	"elemen7s.com/accountscards"
	"elemen7s.com/decks"
	"elemen7s.com/games"
	"fmt"
	"time"
	"ztaylor.me/http"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func PingHandler(r *http.Request) error {
	if r.Session == nil {
		r.WriteJson(js.Object{
			"uri": "/ping",
			"data": js.Object{
				"online": http.SessionCount(),
				"cards":  AllCardsJson("en-US"),
			},
		})
		return nil
	} else if account, err := accounts.Get(r.Username); account == nil {
		return err
	} else if decks, err := decks.Get(r.Username); decks == nil {
		return err
	} else if accountcards, err := accountscards.Get(r.Username); accountcards == nil {
		return err
	} else {
		cardsdata := AllCardsJson("en-US")
		for cardid, cards := range accountcards {
			if k := fmt.Sprintf("%d", cardid); cardsdata[k] == nil {
				log.Add("Key", k).Add("CardId", cardid).Add("Username", r.Username).Add("Copies", len(cards)).Warn("copies of missing card")
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

func pingHandlerDataHelperGames(name string) map[int]js.Object {
	gamesdata := make(map[int]js.Object)
	for _, gameid := range games.Cache.GetPlayerGames(name) {
		if game := games.Cache.Get(gameid); game == nil {
		} else if seat := game.GetSeat(name); seat != nil {
			gamedata := seat.Json()
			gamedata["gameid"] = gameid
			gamedata["timer"] = int(game.Active.Duration.Seconds())
			opponentsdata := make([]string, 0)
			for _, seat2 := range game.Seats {
				if seat2.Username != name {
					opponentsdata = append(opponentsdata, seat2.Username)
				}
			}
			gamedata["opponents"] = opponentsdata
			gamesdata[gameid] = gamedata
		}
	}
	return gamesdata
}