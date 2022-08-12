package match

import (
	"github.com/zachtaylor/7elements/account"
	"github.com/zachtaylor/7elements/game"
)

type Server interface {
	game.Server
	Accounts() *account.Cache
	Games() game.Manager
}
