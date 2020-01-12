package game

import (
	"time"

	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/chat"
	"ztaylor.me/cast"
	"ztaylor.me/log"
)

// Runtime (game) refers for the engine
type Runtime struct {
	Root    *vii.Runtime
	Service Service
	Timeout time.Duration
	logger  log.Service
	chat    chat.Service
}

// NewRuntime creates a new game.Runtime from base Runtime
func NewRuntime(root *vii.Runtime, service Service, timeout time.Duration, w cast.WriteCloser, chat chat.Service) *Runtime {
	logger := log.NewService(log.LevelDebug, log.DefaultFormatWithoutColor(), w)
	logger.Formatter().CutSourcePath(0)
	return &Runtime{
		Root:    root,
		Service: service,
		Timeout: timeout,
		logger:  logger,
		chat:    chat,
	}
}
