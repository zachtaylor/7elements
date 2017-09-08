package ctxpert

import (
	"context"
	"sync"
	"time"
)

type Context struct {
	sync.Mutex
	ctx    context.Context
	cancel context.CancelFunc
	done   []Timeout
	always []Timeout
	store  map[string]interface{}
	isDone bool
}

func (ctx *Context) Done(f Timeout) *Context {
	ctx.Lock()
	defer ctx.Unlock()
	if ctx.isDone { // error
		return nil
	}
	ctx.done = append(ctx.done, f)
	return ctx
}

func (ctx *Context) Always(f Timeout) *Context {
	ctx.Lock()
	defer ctx.Unlock()
	if ctx.isDone { // error
		return nil
	}
	ctx.always = append(ctx.always, f)
	return ctx
}

func (ctx *Context) Store(key string, val interface{}) *Context {
	ctx.Lock()
	defer ctx.Unlock()
	if ctx.isDone { // error
		return nil
	}
	ctx.store[key] = val
	return ctx
}

func (ctx *Context) CopyStore() map[string]interface{} {
	ctx.Lock()
	defer ctx.Unlock()
	if ctx.isDone { // error
		return nil
	}
	cp := make(map[string]interface{})
	for k, v := range ctx.store {
		cp[k] = v
	}
	return cp
}

func (ctx *Context) Get(key string) interface{} {
	ctx.Lock()
	defer ctx.Unlock()
	if ctx.isDone { // error
		return nil
	}
	return ctx.store[key]
}

func (ctx *Context) Cancel() {
	ctx.Lock()
	if ctx.isDone {
		ctx.Unlock()
		return
	}
	ctx.Unlock()
	for _, f := range ctx.always {
		f(ctx)
	}
	ctx.cancel()
}

func (ctx *Context) Timer() time.Duration {
	if t, ok := ctx.ctx.Deadline(); ok {
		return t.Sub(time.Now())
	}
	return -1
}

func (ctx *Context) start() *Context {
	go func() {
		<-ctx.ctx.Done()
		ctx.Lock()
		ctx.isDone = true
		ctx.Unlock()

		if ctx.ctx.Err().Error() != "context canceled" {
			for _, f := range ctx.done {
				f(ctx)
			}
			for _, f := range ctx.always {
				f(ctx)
			}
		}
		ctx.done = nil
		ctx.always = nil
		ctx.store = nil
	}()
	return ctx
}
