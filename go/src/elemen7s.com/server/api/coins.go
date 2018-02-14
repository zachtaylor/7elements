package api

import (
	"elemen7s.com/accounts"
	"ztaylor.me/http"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func CoinsHandler(r *http.Request) error {
	if r.Session == nil {
		return ErrSessionRequired
	} else if account, err := accounts.Get(r.Username); account == nil {
		return err
	} else {
		account.Coins += 10
		r.WriteJson(js.Object{
			"username": account.Username,
			"coins":    account.Coins,
		})

		log.WithFields(log.Fields{"Remote": r.Remote,
			"Username": r.Username,
			"Coins":    account.Coins,
		}).Info("coins")
		return nil
	}
}
