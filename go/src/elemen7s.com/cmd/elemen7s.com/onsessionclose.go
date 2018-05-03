package main

import (
	"elemen7s.com"
	"elemen7s.com/accounts"
	"elemen7s.com/accountscards"
	"ztaylor.me/events"
	"ztaylor.me/log"
)

func init() {
	events.On("SessionClose", func(args ...interface{}) {
		username := args[0].(string)
		accounts.Forget(username)
		accountscards.Forget(username)
		vii.AccountDeckService.Forget(username)
		log.Add("Username", username).Info("account cached data cleared")
	})
}
