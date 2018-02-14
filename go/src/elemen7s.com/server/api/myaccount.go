package api

import (
	"elemen7s.com/accounts"
	"elemen7s.com/accountscards"
	"time"
	"ztaylor.me/http"
	"ztaylor.me/js"
)

func MyAccountHandler(r *http.Request) error {
	if r.Session == nil {
		return ErrSessionRequired
	} else if account, err := accounts.Get(r.Username); account == nil {
		return err
	} else if accountcards, err := accountscards.Get(r.Username); err != nil {
		return err
	} else {
		r.WriteJson(js.Object{
			"username":     r.Username,
			"language":     account.Language,
			"email":        account.Email,
			"session-life": r.Expire.Sub(time.Now()).String(),
			"coins":        account.Coins,
			"packs":        account.Packs,
			"cards":        accountcards.Json(),
		})
		return nil
	}
}
