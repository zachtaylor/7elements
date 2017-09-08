package api

import (
	"7elements.ztaylor.me/accounts"
	"7elements.ztaylor.me/accountscards"
	"7elements.ztaylor.me/decks"
	"7elements.ztaylor.me/games"
	"7elements.ztaylor.me/server/sessionman"
	"net/http"
	"time"
	"ztaylor.me/json"
	"ztaylor.me/log"
	// "strconv"
)

var NavbarJsonHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log := log.Add("RemoteAddr", r.RemoteAddr).Add("Referrer", r.Referer())
	if r.Method != "GET" {
		log.Warn("navbar.json: only GET supported")
	}

	session, err := sessionman.ReadRequestCookie(r)
	if session == nil {
		sessionman.EraseSessionId(w)
		w.WriteHeader(401)

		if err != nil {
			log.Add("Error", err)
			w.Write([]byte(err.Error()))
		} else {
			w.Write([]byte("session missing"))
		}

		log.Debug("navbar.json: no session")
		return
	}

	log.Add("Username", session.Username)
	account := accounts.Test(session.Username)
	if account == nil {
		w.WriteHeader(500)
		w.Write([]byte("account missing"))
		log.Warn("navbar.json: account missing")
		return
	}

	decks, err := decks.Get(session.Username)
	if decks == nil {
		w.WriteHeader(500)
		w.Write([]byte("decks missing"))
		log.Add("Error", err).Error("navbar.json: failed finding decks")
		return
	}

	accountcards, err := accountscards.Get(session.Username)
	if accountcards == nil {
		w.WriteHeader(500)
		w.Write([]byte("cards missing"))
		log.Add("Error", err).Error("navbar.json: failed finding cards")
		return
	}

	gamesdata := make(map[int]interface{})
	for gameid, ok := range games.GetActiveGames(session.Username) {
		if game := games.Cache[gameid]; !ok || game == nil {
		} else if seat := game.GetSeat(session.Username); seat != nil {
			gamedata := seat.Json()
			gamedata["gameid"] = gameid
			gamedata["timer"] = int(game.Context.Timer().Seconds() + 1)
			opponentsdata := make([]string, 0)
			for _, seat2 := range game.Seats {
				if seat2.Username != seat.Username {
					opponentsdata = append(opponentsdata, seat2.Username)
				}
			}

			gamedata["opponents"] = opponentsdata
			gamesdata[gameid] = gamedata
		}
	}

	session.Refresh()
	json.Json{
		"username":       account.Username,
		"email":          account.Email,
		"session-expire": session.Expire.Sub(time.Now()).String(),
		"coins":          account.Coins,
		"packs":          account.Packs,
		"cards":          accountcards.Json(),
		"decks":          decks.Json(),
		"games":          gamesdata,
	}.Write(w)
	log.Info("navbar.json")
})
