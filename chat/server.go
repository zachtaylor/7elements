package chat

import (
	"github.com/zachtaylor/7elements/internal/log"
	"github.com/zachtaylor/7elements/internal/user"
)

type Server interface {
	log.Server
	user.Server
}
