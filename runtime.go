package vii

import (
	"github.com/zachtaylor/7elements/card"
	"ztaylor.me/cast"
	"ztaylor.me/log"
)

// Runtime holds service refs
type Runtime struct {
	Logger        log.Service
	Accounts      AccountService
	AccountsCards AccountCardService
	AccountsDecks AccountDeckService
	Cards         card.PrototypeService
	Decks         DeckService
	Packs         PackService
}

func (rt *Runtime) SendAccountUpdate(sender func(cast.JSON), name string) {
	sender(rt.AccountJSON(name))
}

func (rt *Runtime) AccountJSON(name string) cast.JSON {
	if a, _ := rt.Accounts.Find(name); a == nil {
		return nil
	} else if acs, err := rt.AccountsCards.Find(a.Username); err != nil {
		return nil
	} else if ads, err := rt.AccountsDecks.Find(a.Username); err != nil {
		return nil
	} else {
		return cast.JSON{
			"username": a.Username,
			"email":    a.Email,
			"session":  a.SessionID,
			"coins":    a.Coins,
			"cards":    acs.JSON(),
			"decks":    ads.JSON(),
		}
	}
}

func (rt *Runtime) JSON() cast.JSON {
	decks, _ := rt.Decks.GetAll()
	packs, _ := rt.Packs.GetAll()
	users, _ := rt.Accounts.GetCount()
	return cast.JSON{
		"cards": rt.Cards.GetAll().JSON(),
		"packs": packs.JSON(),
		"decks": decks.JSON(),
		"users": users,
	}
}
