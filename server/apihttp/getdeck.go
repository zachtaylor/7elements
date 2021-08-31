package apihttp

// import (
// 	"github.com/zachtaylor/7elements/account"
// 	"github.com/zachtaylor/7elements/server/runtime"
// 	"ztaylor.me/log"
// )

// func GetMyDeck(t *runtime.T, log *log.Entry, username string, deckid int) *account.Deck {
// 	if mydecks, err := rt.Root.AccountsDecks.Find(username); mydecks == nil {
// 		log.Add("User", username).Add("Error", err).Error("user missing")
// 	} else if d := mydecks[deckid]; d == nil {
// 		log.Add("DeckID", deckid).Error("deck missing")
// 	} else {
// 		return d
// 	}
// 	return nil
// }

// func GetFreeDeck(t *runtime.T, username string, deckid int) *account.Deck {
// 	if d, err := rt.Root.Decks.Get(deckid); d == nil {
// 		log.Add("Error", err).Error("free decks missing")
// 	} else {
// 		return account.NewDeckWith(d, username)
// 	}
// 	return nil
// }
