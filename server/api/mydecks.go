package api

// import (
// 	"github.com/zachtaylor/7elements"
// 	"ztaylor.me/http"
// 	"ztaylor.me/log"
// )

// func MyDecksHandler(r *http.Quest) error {
// 	if r.Session == nil {
// 		return ErrSessionRequired
// 	}
// 	log := log.Add("Username", r.Session.Username)
// 	if decks, err := vii.AccountDeckService.Get(r.Session.Username); err != nil {
// 		return err
// 	} else {
// 		r.WriteJson(decks.Json())
// 		log.Debug("/api/mydecks.json")
// 		return nil
// 	}
// }
