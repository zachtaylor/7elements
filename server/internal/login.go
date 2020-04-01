package internal

import (
	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/account"
	"ztaylor.me/cast"
	"ztaylor.me/http/session"
)

func Login(rt *vii.Runtime, sessions session.Service, a *account.T) (*session.T, error) {
	rt.Accounts.Cache(a)
	log := rt.Logger.New().Add("Username", a.Username).Source()
	a.LastLogin = cast.Now()
	if err := rt.Accounts.UpdateLogin(a); err != nil {
		return nil, err
	}
	if a.SessionID != "" {
		if s := sessions.Get(a.SessionID); s != nil {
			log.Debug("found")
			return s, nil
		}
	}
	s := sessions.Grant(a.Username)
	a.SessionID = s.ID()
	go loginWaiter(rt, s)
	return s, nil
}

func loginWaiter(rt *vii.Runtime, s *session.T) {
	for {
		if _, ok := <-s.Done(); !ok {
			break
		}
	}
	rt.Logger.New().With(cast.JSON{
		"SessionID": s.ID(),
		"Username":  s.Name(),
	}).Source().Info("done")
	rt.Accounts.Forget(s.Name())
	rt.AccountsCards.Forget(s.Name())
	rt.AccountsDecks.Forget(s.Name())
}
