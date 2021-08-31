package apihttp

// import (
// 	"github.com/zachtaylor/7elements"
// 	"ztaylor.me/http"
// 	"ztaylor.me/js"
// 	"ztaylor.me/log"
// )

// func CoinsHandler(r *http.Quest) error {
// 	if r.Session == nil {
// 		return ErrSessionRequired
// 	} else if account, err := vii.AccountService.Get(r.Username); account == nil {
// 		return err
// 	} else {
// 		account.Coins += 10
// 		r.WriteJson(js.Object{
// 			"username": account.Username,
// 			"coins":    account.Coins,
// 		})

// 		log.WithFields(websocket.MsgData{"Remote": r.Remote,
// 			"Username": r.Username,
// 			"Coins":    account.Coins,
// 		}).Info("coins")
// 		return nil
// 	}
// }
