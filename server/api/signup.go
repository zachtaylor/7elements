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

func Signup(server internal.Server, username, email, password string) (account *account.T, session *session.T, err error) {
	if err = CheckUsername(username); err != nil {
		return
	} else if err = CheckEmail(email); err != nil {
		return
	} else if _, err = accounts.Get(server.GetDB(), username); err != sql.ErrNoRows {
		return nil, nil, ErrSignupExists
	} else {
		session = server.GetSessionManager().Must(username)
	}

	account = account.Make(username, email, password, session.ID())
	for i := 0; i < 3; i++ {
		proto := deck.NewPrototype(username)
		proto.ID = i + 1
		account.Decks[proto.ID] = proto
	}

	err = accounts.Insert(server.GetDB(), account)
	if err != nil {
		server.GetSessionManager().Remove(session.ID())
		return nil, nil, err
	}
	server.GetAccounts().Set(username, account)

	go server.Ping()

	return account, session, nil
}
