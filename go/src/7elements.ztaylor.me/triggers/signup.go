package triggers

import (
	"7elements.ztaylor.me"
	"7elements.ztaylor.me/event"
	"7elements.ztaylor.me/log"
	"time"
)

func init() {
	event.On("Signup", func(args ...interface{}) {
		username := args[0].(string)
		log.Add("Username", username)

		register := time.Now()

		SE.AccountsPacks.Cache[username] = []*SE.AccountPack{
			&SE.AccountPack{Register: register},
			&SE.AccountPack{Register: register},
			&SE.AccountPack{Register: register},
			&SE.AccountPack{Register: register},
			&SE.AccountPack{Register: register},
			&SE.AccountPack{Register: register},
			&SE.AccountPack{Register: register},
		}

		if err := SE.AccountsPacks.Insert(username); err != nil {
			log.Add("Error", err).Error("signup: grant 7 packs")
		} else {
			log.Debug("signup: grant 7 packs")
		}

		SE.AccountsDecks.Cache[username] = map[int]*SE.AccountDeck{
			1: &SE.AccountDeck{Id: 1, Cards: make(map[int]int)},
			2: &SE.AccountDeck{Id: 2, Cards: make(map[int]int)},
			3: &SE.AccountDeck{Id: 3, Cards: make(map[int]int)},
		}

		if err := SE.AccountsDecks.Insert(username, 0); err != nil {
			log.Add("Error", err).Error("signup: grant 3 decks")
		} else {
			log.Debug("signup: grand 3 decks")
		}
	})
}
