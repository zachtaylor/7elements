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
	})
}
