package triggers

import (
	"7elements.ztaylor.me/accounts"
	"7elements.ztaylor.me/accountscards"
	"7elements.ztaylor.me/decks"
	"7elements.ztaylor.me/event"
	"ztaylor.me/log"
)

func init() {
	event.On("SessionRevoke", func(args ...interface{}) {
		username := args[0].(string)
		accounts.Forget(username)
		accountscards.Forget(username)
		decks.Forget(username)
		log.Add("Username", username).Debug("sessionrevoke: uncache account")
	})
}
