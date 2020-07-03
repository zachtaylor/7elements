package vii

// import (
// 	"github.com/zachtaylor/7elements/account"
// 	"github.com/zachtaylor/7elements/card"
// 	"github.com/zachtaylor/7elements/card/pack"
// 	"github.com/zachtaylor/7elements/deck"
// 	"ztaylor.me/cast"
// 	"ztaylor.me/log"
// )

// // Runtime holds service refs
// type Runtime struct {
// 	Logger        log.Service
// 	Accounts      account.Service
// 	AccountsCards account.CardService
// 	AccountsDecks account.DeckService
// 	Cards         card.PrototypeService
// 	Decks         deck.PrototypeService
// 	Packs         pack.Service
// }

// // func (rt *Runtime) SendAccountUpdate(sender func(cast.JSON), name string) {
// // 	sender(rt.AccountJSON(name))
// // }

// func (rt *Runtime) FindAccountJSON(name string) cast.JSON {
// 	a, _ := rt.Accounts.Find(name)
// 	if a == nil {
// 		return nil
// 	}
// 	return rt.AccountJSON(a)
// }

// func (rt *Runtime) AccountJSON(a *account.T) cast.JSON {
// 	if a == nil {
// 		return nil
// 	} else if acs, err := rt.AccountsCards.Find(a.Username); err != nil {
// 		return nil
// 	} else if ads, err := rt.AccountsDecks.Find(a.Username); err != nil {
// 		return nil
// 	} else {
// 		return cast.JSON{
// 			"username": a.Username,
// 			"email":    a.Email,
// 			"session":  a.SessionID,
// 			"coins":    a.Coins,
// 			"cards":    acs.JSON(),
// 			"decks":    ads.JSON(),
// 		}
// 	}
// }
