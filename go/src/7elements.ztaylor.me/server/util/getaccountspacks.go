package serverutil

import (
	"7elements.ztaylor.me"
)

func GetAccountsPacks(username string) ([]*SE.AccountPack, error) {
	if accountspacks := SE.AccountsPacks.Cache[username]; accountspacks == nil {
		if ap, err := SE.AccountsPacks.Get(username); err != nil {
			return nil, err
		} else {
			SE.AccountsPacks.Cache[username] = ap
			return ap, nil
		}
	} else {
		return accountspacks, nil
	}
}
