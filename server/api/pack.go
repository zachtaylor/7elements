package api

// import (
// 	"github.com/zachtaylor/7elements"
// 	"ztaylor.me/http"
// 	"ztaylor.me/js"
// 	"ztaylor.me/log"
// )

// func PackHandler(r *http.Request) error {
// 	if r.Session == nil {
// 		return ErrSessionRequired
// 	} else if account, err := vii.AccountService.Get(r.Username); account == nil {
// 		return err
// 	} else if account.Coins < 7 {
// 		r.WriteJson(js.Object{
// 			"username": account.Username,
// 			"coins":    account.Coins,
// 			"packs":    account.Packs,
// 		})
// 		return ErrInsufficientFunds
// 	} else {
// 		account.Coins -= 7
// 		account.Packs++
// 		r.WriteJson(js.Object{
// 			"username": account.Username,
// 			"coins":    account.Coins,
// 			"packs":    account.Packs,
// 		})
// 		log.WithFields(log.Fields{
// 			"Remote":   r.Remote,
// 			"Username": r.Username,
// 			"Coins":    account.Coins,
// 			"Packs":    account.Packs,
// 		}).Info("pack")
// 		return nil
// 	}
// }
