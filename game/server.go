package game

import (
	"github.com/zachtaylor/7elements/content"
	"github.com/zachtaylor/7elements/internal/log"
	"github.com/zachtaylor/7elements/internal/user"
)

// Server is how service sees its' parent
type Server interface {
	content.Server
	log.Server
	user.Server
}
