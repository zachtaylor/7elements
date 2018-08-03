package api

import (
	"github.com/zachtaylor/7tcg"
	"time"
	"ztaylor.me/http"
	"ztaylor.me/js"
)

func MyAccountHandler(r *http.Request) error {
	if r.Session == nil {
		return ErrSessionRequired
	} else if account, err := vii.AccountService.Get(r.Username); account == nil {
		return err
	} else if accountcards, err := vii.AccountCardService.Get(r.Username); err != nil {
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
