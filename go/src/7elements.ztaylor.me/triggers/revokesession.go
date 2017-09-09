package triggers

import (
	"7elements.ztaylor.me/accounts"
	"7elements.ztaylor.me/accountscards"
	"7elements.ztaylor.me/decks"
	"ztaylor.me/events"
	"ztaylor.me/log"
)

func init() {
	events.On("SessionRevoke", func(args ...interface{}) {
		username := args[0].(string)
		accounts.Forget(username)
		accountscards.Forget(username)
		decks.Forget(username)
		log.Add("Username", username).Debug("sessionrevoke: uncache account")
	})
}
