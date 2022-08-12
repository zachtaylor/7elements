package api

import (
	"database/sql"
	"errors"

	"github.com/zachtaylor/7elements/account"
	"github.com/zachtaylor/7elements/db/accounts"
	"github.com/zachtaylor/7elements/deck"
	"github.com/zachtaylor/7elements/server/internal"
	"taylz.io/http/session"
)

var (
	ErrSignupExists = errors.New("username exists")
)

func Signup(server internal.Server, username, email, password string) (a *account.T, s *session.T, err error) {
	if err = CheckUsername(username); err != nil {
		return
	} else if err = CheckEmail(email); err != nil {
		return
	} else if _, err = accounts.Get(server.DB(), username); err != sql.ErrNoRows {
		return nil, nil, ErrSignupExists
	} else {
		s = server.Sessions().Must(username)
	}

	a = account.Make(username, email, password, s.ID())
	for i := 0; i < 3; i++ {
		proto := deck.NewPrototype(username)
		proto.ID = i + 1
		a.Decks[proto.ID] = proto
	}

	err = accounts.Insert(server.DB(), a)
	if err != nil {
		server.Sessions().Remove(s.ID())
		return nil, nil, err
	}
	server.Accounts().Set(username, a)

	go server.Ping()

	return a, s, nil
}
