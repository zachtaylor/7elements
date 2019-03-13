package api

// import (
// 	"github.com/zachtaylor/7elements"
// 	"ztaylor.me/http"
// )

// func GameHandler(r *http.Quest) error {
// 	if gameid := r.Data.Sval("gameid"); gameid != "" {
// 		return ErrGameIdRequired
// 	} else if game := game.Service.Get(gameid); game == nil {
// 		return ErrGameMissing
// 	} else {
// 		game.In <- &game.Request{
// 			Username: r.Username,
// 			Data:     r.Data,
// 		}
// 		return nil
// 	}
// }
