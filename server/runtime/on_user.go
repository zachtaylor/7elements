package runtime

import "taylz.io/http/user"

func (rt *T) OnUser(name string, oldUser, newUser *user.T) {
	go rt.onUser(name, oldUser, newUser)
}

func (rt *T) onUser(name string, oldUser, newUser *user.T) {
	log := rt.Logger.Add("User", name)
	if oldUser == nil && newUser != nil {
		log.Debug("new")
	} else if oldUser != nil && newUser == nil {
		log.Debug("old")
		if rt.MatchMaker.Get(name) != nil {
			log.Warn("expire queue")
			rt.MatchMaker.Cancel(name)
		}
		rt.Accounts.Remove(name)
	} else {
		log.With(map[string]interface{}{
			"Old": oldUser,
			"New": newUser,
		}).Warn("weird")
	}
	go rt.Ping()
}
