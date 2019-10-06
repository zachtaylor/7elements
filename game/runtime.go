package game

import (
	"time"

	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/chat"
	"ztaylor.me/log"
)

type Runtime struct {
	Root    *vii.Runtime
	Service Service
	Timeout time.Duration
	logger  log.Service
	chat    chat.Service
}

func NewRuntime(root *vii.Runtime, service Service, timeout time.Duration, logger log.Service, chat chat.Service) *Runtime {
	return &Runtime{
		Root:    root,
		Service: service,
		Timeout: timeout,
		logger:  logger,
		chat:    chat,
	}
}
