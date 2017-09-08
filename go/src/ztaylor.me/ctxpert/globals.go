package ctxpert

import (
	"context"
	"time"
)

func New() *Context {
	ctx, cancel := context.WithCancel(context.Background())
	return (&Context{
		ctx:    ctx,
		cancel: cancel,
		done:   make(timeouts, 0),
		always: make(timeouts, 0),
		store:  make(map[string]interface{}),
	}).start()
}

func WithNewTimeout(ctx *Context, d time.Duration) *Context {
	storecopy := ctx.CopyStore()
	ctx.cancel()
	ctx.Lock()
	defer ctx.Unlock()

	context, cancel := context.WithTimeout(context.Background(), d)
	return (&Context{
		ctx:    context,
		cancel: cancel,
		done:   timeouts(ctx.done).copy(),
		always: timeouts(ctx.always).copy(),
		store:  storecopy,
	}).start()
}
