package internal

import (
	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/account"
	"ztaylor.me/cast"
	"ztaylor.me/http/session"
)

func Signup(rt *vii.Runtime, sessions session.Service, username, email, password string) (*account.T, *session.T, error) {
	a := &account.T{
		Username: username,
		Email:    email,
		Password: password,
		Skill:    1000,
		Register: cast.Now(),
	}

	s, err := Login(rt, sessions, a)
	if err != nil {
		return nil, nil, err
	}

	if err := rt.Accounts.Insert(a); err != nil {
		return nil, nil, err
	}

	for i := 1; i < 4; i++ {
		deck := account.NewDeck()
		deck.ID = i
		deck.Username = username
		if err := rt.AccountsDecks.Update(deck); err != nil {
			return nil, nil, err
		}
	}

	// if err := emails.SendValidationEmail(account); err != nil {
	// 	log.Clone().Add("mail-user", options.String("mail-user")).Add("mail-pass", options.String("mail-pass")).Add("mail-host", options.String("mail-host")).Add("Error", err).Error("/api/signup: send validation email")
	// }

	return a, s, nil
}
