package session

import (
	"github.com/zachtaylor/7elements/account"
	"taylz.io/http/session"
	"taylz.io/yas"
)

type Cache = yas.Cache[*T]

type T struct {
	Account account.Row
	Session *session.T
}
