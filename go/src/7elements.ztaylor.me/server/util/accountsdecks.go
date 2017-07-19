package serverutil

import (
	"7elements.ztaylor.me"
)

func GetAccountsDecks(username string) (map[int]*SE.AccountDeck, error) {
	if SE.AccountsDecks.Cache[username] == nil {
		if decks, err := SE.AccountsDecks.Get(username); err != nil {
			return nil, err
		} else {
			SE.AccountsDecks.Cache[username] = decks
		}
	}
	return SE.AccountsDecks.Cache[username], nil
}
