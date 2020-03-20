package vii

import (
	"github.com/zachtaylor/7elements/account"
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/card/pack"
	"github.com/zachtaylor/7elements/deck"
	"ztaylor.me/cast"
	"ztaylor.me/log"
)

// Runtime holds service refs
type Runtime struct {
	Logger        log.Service
	Accounts      account.Service
	AccountsCards account.CardService
	AccountsDecks account.DeckService
	Cards         card.PrototypeService
	Decks         deck.PrototypeService
	Packs         pack.Service
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
