package api

import (
	"time"

	"github.com/zachtaylor/7elements"
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
	} else if accountdecks, err := vii.AccountDeckService.Get(r.Username); err != nil {
		return err
	} else {
		r.WriteJson(js.Object{
			"username":    r.Username,
			"email":       account.Email,
			"sessionlife": r.Expire.Sub(time.Now()).String(),
			"coins":       account.Coins,
			"cards":       accountcards.Json(),
			"decks":       accountdecks.Json(),
		})
		return nil
	}
}
