package persistence

import (
	"7elements.ztaylor.me"
)

func init() {
	SE.Accounts.Get = AccountsGet
	SE.Accounts.Insert = AccountsInsert
	SE.Accounts.Delete = AccountsDelete
	SE.AccountsCards.Get = AccountsCardsGet
	SE.AccountsCards.Delete = AccountsCardsDelete
	SE.Cards.LoadCache = CardsLoadCache
	SE.Cards.Insert = CardsInsert
	SE.Cards.Delete = CardsDelete
	SE.CardTexts.LoadCache = CardTextsLoadCache
}
