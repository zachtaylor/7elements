package api

import (
	"elemen7s.com"
	"ztaylor.me/http"
)

func GameHandler(r *http.Request) error {
	if gameid := r.Data.Sval("gameid"); gameid != "" {
		return ErrGameIdRequired
	} else if game := vii.GameService.Get(gameid); game == nil {
		return ErrGameMissing
	} else {
		game.In <- &vii.GameRequest{
			Username: r.Username,
			Data:     r.Data,
		}
		return nil
	}
}
