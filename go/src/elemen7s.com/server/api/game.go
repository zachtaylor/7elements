package api

import (
	"elemen7s.com/games"
	"ztaylor.me/http"
)

func GameHandler(r *http.Request) error {
	if gameid := r.Data.Ival("gameid"); gameid < 1 {
		return ErrGameIdRequired
	} else if game := games.Cache.Get(gameid); game == nil {
		return ErrGameMissing
	} else {
		game.Receive(r.Username, r.Data)
		return nil
	}
}
