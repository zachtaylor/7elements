package api

import (
	"database/sql"
	"errors"

	"github.com/zachtaylor/7elements/account"
	"github.com/zachtaylor/7elements/db/accounts"
	"github.com/zachtaylor/7elements/deck"
	"github.com/zachtaylor/7elements/server/runtime"
	"taylz.io/http/session"
)

var (
	ErrSignupExists = errors.New("username exists")
)

func Signup(rt *runtime.T, username, email, password string) (account *account.T, session *session.T, err error) {
	if err = CheckUsername(username); err != nil {
		return
	} else if err = CheckEmail(email); err != nil {
		return
	} else if _, err = accounts.Get(rt.DB, username); err != sql.ErrNoRows {
		return nil, nil, ErrSignupExists
	} else {
		session = rt.Sessions.Must(username)
	}

	account = account.Make(username, email, password, session.ID())
	for i := 0; i < 3; i++ {
		proto := deck.NewPrototype(username)
		proto.ID = i + 1
		account.Decks[proto.ID] = proto
	}

	err = accounts.Insert(rt.DB, account)
	if err != nil {
		rt.Sessions.Remove(session.ID())
		return nil, nil, err
	}
	rt.Accounts.Set(username, account)

	return account, session, nil
}
