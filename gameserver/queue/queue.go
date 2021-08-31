package queue

import (
	"github.com/zachtaylor/7elements/account"
	"github.com/zachtaylor/7elements/gameserver"
	"taylz.io/types"
)

//go:generate go-gengen -p=queue -k=string -v=*T

type T struct {
	Account *account.T
	Entry   *gameserver.Entry
	Start   types.Time
	done    chan string
	cancel  chan bool
}

func (t *T) Finish(id string) {
	t.done <- id
	t.close()
}

func (t *T) Done() <-chan string {
	return t.done
}

func (t *T) Cancel() <-chan bool {
	return t.cancel
}

func (t *T) Close() {
	t.cancel <- true
	t.close()
}

func (t *T) close() {
	close(t.done)
	close(t.cancel)
}
