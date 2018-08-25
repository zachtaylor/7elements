package api

import (
	"github.com/zachtaylor/7elements"
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
