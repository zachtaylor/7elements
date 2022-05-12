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
	if account = server.GetAccounts().Get(username); account != nil {
		if account.Password != password {
			account = nil
			err = ErrLoginPassword
		} else {
			session = server.GetSessionManager().Get(account.SessionID)
			err = ErrLoginSession
		}
	} else if account, err = accounts.Get(server.GetDB(), username); err != nil {
		account = nil
	} else if account == nil {
		err = ErrLoginUsername
	} else if account.Password != password {
		account = nil
		err = ErrLoginPassword
	} else {
		if session = server.GetSessionManager().GetName(username); session != nil {
			err = ErrLoginSession
		} else {
			session = server.GetSessionManager().Must(username)
		}
		account.SessionID = session.ID()
		server.GetAccounts().Set(username, account)
	}
	return
}
