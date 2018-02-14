package api

import (
	"elemen7s.com/games"
	"ztaylor.me/http"
	"ztaylor.me/log"
)

func JoinHandler(r *http.Request) error {
	if gameid := r.Data.Ival("gameid"); gameid < 1 {
		return ErrGameIdRequired
	} else if game := games.Cache.Get(int(gameid)); game == nil {
		return ErrGameMissing
	} else if seat := game.GetSeat(r.Username); seat == nil {
		log.Add("GameId", game.Id).Add("Username", r.Username).Warn("/api/join: not participating in game")
	} else if seat.Player != nil {
		log.Add("GameId", game.Id).Add("Username", r.Username).Warn("/api/join: seat already occupied")
	} else {
		games.ConnectPlayerAgent(seat, r.Agent)
		game.SendCatchup(seat)
	}
	return nil
}
