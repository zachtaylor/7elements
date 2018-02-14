package api

import (
	"elemen7s.com/accounts"
	"elemen7s.com/accountscards"
	"fmt"
	"ztaylor.me/http"
	"ztaylor.me/js"
)

func MyCardsHandler(r *http.Request) error {
	if r.Session == nil {
		return ErrSessionRequired
	} else if account, err := accounts.Get(r.Username); account == nil {
		return err
	} else if accountcards, err := accountscards.Get(r.Username); err != nil {
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
