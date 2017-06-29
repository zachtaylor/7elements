package serverutil

import (
	"7elements.ztaylor.me"
)

func GetAccountsCards(username string) (SE.CardCollection, error) {
	if cardcollection := SE.AccountsCards.Cache[username]; cardcollection == nil {
		if cc, err := SE.AccountsCards.Get(username); err != nil {
			return nil, err
		} else {
			SE.AccountsCards.Cache[username] = cc
			return cc, nil
		}
	} else {
		return cardcollection, nil
	}
}

func FlushAccountsCards(username string) error {
	if cardcollection := SE.AccountsCards.Cache[username]; cardcollection != nil {
		if err := SE.AccountsCards.Delete(username); err != nil {
			return err
		} else if err := SE.AccountsCards.Insert(username); err != nil {
			return err
		}
	}

	return nil
}
