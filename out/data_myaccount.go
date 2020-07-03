package out

import "github.com/zachtaylor/7elements/player"

// Account attempts to map `Target` to `*player.T`, to send `Account` data using `MyAccount`
func Account(t Target) {
	if player, _ := t.(*player.T); player == nil {
	} else {
		MyAccount(player)
	}
}

// MyAccount writes account data to "/data/myaccount"
func MyAccount(player *player.T) {
	if player == nil {
		return
	}
	player.Send("/data/myaccount", player.Account.JSON())
}
