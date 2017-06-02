package persistence

import (
	"7elements.ztaylor.me"
)

func init() {
	SE.Accounts.Get = AccountsGet
	SE.Accounts.Insert = AccountsInsert
	SE.Accounts.Delete = AccountsDelete
	SE.AccountsCards.Get = AccountsCardsGet
	SE.AccountsCards.Insert = AccountsCardsInsert
	SE.AccountsCards.Delete = AccountsCardsDelete
	SE.AccountsPacks.Get = AccountsPacksGet
	SE.AccountsPacks.Insert = AccountsPacksInsert
	SE.AccountsPacks.Delete = AccountsPacksDelete
	SE.Cards.LoadCache = CardsLoadCache
	SE.Cards.Insert = CardsInsert
	SE.Cards.Delete = CardsDelete
	SE.CardTexts.LoadCache = CardTextsLoadCache
}
