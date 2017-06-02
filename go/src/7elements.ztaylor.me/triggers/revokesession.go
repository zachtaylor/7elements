package triggers

import (
	"7elements.ztaylor.me"
	"7elements.ztaylor.me/event"
	"7elements.ztaylor.me/log"
)

func init() {
	event.On("RevokeSession", func(args ...interface{}) {
		username := args[0].(string)
		delete(SE.Accounts.Cache, username)
		log.Add("Username", username).Debug("revokesession: uncache account")
	})

	event.On("RevokeSession", func(args ...interface{}) {
		username := args[0].(string)
		delete(SE.AccountsCards.Cache, username)
		log.Add("Username", username).Debug("revokesession: uncache accountscards")
	})

	event.On("RevokeSession", func(args ...interface{}) {
		username := args[0].(string)
		delete(SE.AccountsPacks.Cache, username)
		log.Add("Username", username).Debug("revokesession: uncache accountspacks")
	})
}
