package main

import (
	"github.com/zachtaylor/7tcg"
	"ztaylor.me/events"
	"ztaylor.me/log"
)

func init() {
	events.On("SessionClose", func(args ...interface{}) {
		username := args[0].(string)
		vii.AccountService.Forget(username)
		vii.AccountCardService.Forget(username)
		vii.AccountDeckService.Forget(username)
		log.Add("Username", username).Info("account cached data cleared")
	})
}
