package api

import (
	"elemen7s.com"
	"fmt"
	"ztaylor.me/http"
	"ztaylor.me/js"
)

func MyCardsHandler(r *http.Request) error {
	if r.Session == nil {
		return ErrSessionRequired
	} else if account, err := vii.AccountService.Get(r.Username); account == nil {
		return err
	} else if accountcards, err := vii.AccountCardService.Get(r.Username); err != nil {
		return err
	} else {
		j := js.Object{}
		for cardId, list := range accountcards {
			j[fmt.Sprintf("%d", cardId)] = len(list)
		}
		r.WriteJson(j)
		return nil
	}
}
