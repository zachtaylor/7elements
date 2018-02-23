package api

import (
	"elemen7s.com/decks"
	"ztaylor.me/http"
	"ztaylor.me/log"
)

func DecksHandler(r *http.Request) error {
	if r.Session == nil {
		return ErrSessionRequired
	}
	log := log.Add("Username", r.Session.Username)
	if decks, err := decks.Get(r.Session.Username); err != nil {
		return err
	} else {
		r.WriteJson(decks.Json())
		log.Debug("/api/decks")
		return nil
	}
}