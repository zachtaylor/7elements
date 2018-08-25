package api

import (
	"github.com/zachtaylor/7elements/chat"
	"ztaylor.me/http"
	"ztaylor.me/js"
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
		r.Agent.WriteJson(js.Object{
			"uri": "/notification",
			"data": js.Object{
				"class":    "error",
				"username": channel,
				"message":  "Channel not supported yet, sorry!",
			},
		})
		log.Add("Channel", channel).Warn("chat target missing")
	}
	return nil
}
