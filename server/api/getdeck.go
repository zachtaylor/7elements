package api

import (
	"github.com/zachtaylor/7elements/account"
	"ztaylor.me/log"
)

func GetMyDeck(rt *Runtime, log *log.Entry, username string, deckid int) *account.Deck {
	if mydecks, err := rt.Root.AccountsDecks.Find(username); mydecks == nil {
		log.Add("User", username).Add("Error", err).Error("user missing")
	} else if d := mydecks[deckid]; d == nil {
		log.Add("DeckID", deckid).Error("deck missing")
	} else {
		return d
	}
	return nil
}

func GetFreeDeck(rt *Runtime, log *log.Entry, username string, deckid int) *account.Deck {
	if d, err := rt.Root.Decks.Get(deckid); d == nil {
		log.Add("Error", err).Error("free decks missing")
	} else {
		return account.NewDeckWith(d, username)
	}
	return nil
}
