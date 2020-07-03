package runtime

import "ztaylor.me/cast"

func (t *T) GlobalJSON() cast.JSON {
	decks, _ := t.Decks.GetUser("vii")
	packs, _ := t.Packs.GetAll()
	users, _ := t.Accounts.Count()
	return cast.JSON{
		"cards": t.Cards.GetAll().JSON(),
		"packs": packs.JSON(),
		"decks": decks.JSON(),
		"users": users,
	}
}
