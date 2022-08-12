package api

import (
	"errors"

	"github.com/zachtaylor/7elements/account"
	"github.com/zachtaylor/7elements/db/accounts"
	"github.com/zachtaylor/7elements/server/internal"
	"taylz.io/http/session"
)

var (
	// ErrLoginSession is returned as extra information
	ErrLoginSession = errors.New("session exists for name")
	// ErrLoginUsername is returned as rejection
	ErrLoginUsername = errors.New("username not found")
	// ErrLoginPassword is returned as rejection
	ErrLoginPassword = errors.New("wrong password")
)

// Login saves runtime account data
//
// returns (*account.T, *session.T, ErrLoginSession) when a session is joined
func Login(server internal.Server, username, password string) (account *account.T, session *session.T, err error) {
	if account = server.Accounts().Get(username); account != nil {
		if account.Password != password {
			account = nil
			err = ErrLoginPassword
		} else {
			session = server.Sessions().Get(account.SessionID)
			err = ErrLoginSession
		}
	} else if account, err = accounts.Get(server.DB(), username); err != nil {
		account = nil
	} else if account == nil {
		err = ErrLoginUsername
	} else if account.Password != password {
		account = nil
		err = ErrLoginPassword
	} else {
		if user := server.Users().Get(username); user != nil {
			session = user.Session()
			err = ErrLoginSession
		} else {
			session = server.Sessions().Must(username)
		}
		account.SessionID = session.ID()
		server.Accounts().Set(username, account)
	}
	return
}
