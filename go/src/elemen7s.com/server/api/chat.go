package api

import (
	"elemen7s.com/chat"
	"ztaylor.me/http"
	"ztaylor.me/log"
)

func ChatHandler(r *http.Request) error {
	log := log.Add("Remote", r.Remote)

	if r.Session == nil {
		return ErrSessionRequired
	} else if channel := r.Data.Sval("channel"); channel == "" {
		log.Warn("chat: channel missing")
	} else if msg := r.Data.Sval("message"); msg == "" {
		log.Warn("chat: message missing")
	} else if channel == "all" {
		chat.GetChannel("all").AddMessage(chat.NewMessage(r.Session.Username, msg))
	} else {
		log.Add("Channel", channel).Warn("chat target missing")
	}

	return nil
}
