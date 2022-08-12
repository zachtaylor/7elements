package game

import (
	"github.com/zachtaylor/7elements/card"
	"taylz.io/log"
	"taylz.io/yas"
)

type G struct {
	id      string
	keygen  func() string
	done    chan struct{}
	request chan *Request
	data    map[string]any
	seat    yas.Set[string]
	dirty   yas.Set[string]
	rules   Rules
	log     log.Writer
}

func New(id string, rules Rules, keygen func() string, log log.Writer) *G {
	return &G{
		id:      id,
		keygen:  keygen,
		done:    make(chan struct{}),
		request: make(chan *Request, 7),
		data:    map[string]any{},
		seat:    yas.Set[string]{},
		dirty:   yas.Set[string]{},
		rules:   rules,
		log:     log,
	}
}

func (g *G) ID() string      { return g.id }
func (g *G) Log() log.Writer { return g.log }

func (g *G) AddRequest(r *Request) {
	select {
	case <-g.done:
	case g.request <- r:
	}
}

func (g *G) RequestChan() <-chan *Request { return g.request }

func (g *G) Close() {
	close(g.done)
	g.log.Close()
}

func (g *G) Object(id string) any     { return g.data[id] }
func (g *G) Card(id string) *Card     { return Get[*card.Prototype](g, id) }
func (g *G) Player(id string) *Player { return Get[PlayerContext](g, id) }
func (g *G) State(id string) *State   { return Get[StateContext](g, id) }
func (g *G) Token(id string) *Token   { return Get[TokenContext](g, id) }

func (g *G) Players() []string { return g.seat.Slice() }
func (g *G) PlayerCount() int  { return len(g.seat) }

func (g *G) NewPriority(playerID string) Priority {
	priority := Priority{playerID}
	for v := range g.seat {
		if v != playerID {
			priority = append(priority, v)
		}
	}
	return priority
}

func (g *G) MarkUpdate(id string) { g.dirty.Add(id) }
func (g *G) ReadUpdates() (dirtyIDs []string) {
	dirtyIDs = g.dirty.Slice()
	g.dirty = yas.Set[string]{}
	return
}

func (g *G) Rules() Rules { return g.rules }

func Get[T Target](g *G, id string) *Object[T] {
	if obj, _ := g.data[id].(*Object[T]); obj != nil {
		return obj
	}
	return nil
}

// func (g *G) Keygen() string { return g.keygen() }

func (g *G) NewPlayer(ctx PlayerContext) *Player {
	id := ""
	for ok := true; ok; _, ok = g.data[id] {
		id = g.keygen()
	}
	player := NewObject(id, "vii", ctx)
	g.data[id] = player
	g.seat.Add(id)
	return player
}

func (g *G) NewState(playerID string, ctx StateContext) *State {
	id := ""
	for ok := true; ok; _, ok = g.data[id] {
		id = g.keygen()
	}
	state := NewObject(id, playerID, ctx)
	g.data[id] = state
	return state
}

func (g *G) NewCard(playerID string, proto *card.Prototype) *Card {
	id := ""
	for ok := true; ok; _, ok = g.data[id] {
		id = g.keygen()
	}
	card := NewObject(id, playerID, proto)
	g.data[id] = card
	return card
}

func (g *G) NewToken(playerID string, ctx TokenContext) *Token {
	id := ""
	for ok := true; ok; _, ok = g.data[id] {
		id = g.keygen()
	}
	token := NewObject(id, playerID, ctx)
	g.data[id] = token
	return token
}

func (g *G) Write(bytes []byte) {
	for _, playerID := range g.Players() {
		if player := g.Player(playerID); player != nil {
			player.T.Writer.Write(bytes)
		}
	}
}

// func (t *T) Save(obj any) (id string) {
// 	t.sync.Lock()
// 	for ok := true; ok; _, ok = t.data[id] {
// 		id = t.keygen()
// 	}
// 	t.data[id] = Object{
// 		id: id,
// 		it: obj,
// 	}
// 	t.sync.Unlock()
// 	return
// }

// func (t *T) SaveCard(card *Card) {
// 	card.ID = t.objSave(card)
// }

// func (t *T) SaveToken(token *token.T) {
// 	token.ID = t.objSave(token)
// }

// func (t *T) GetState(key string) *State {
// 	if state, ok := t.obj[key].(*State); ok {
// 		return state
// 	}
// 	return nil
// }

// func (t *T) GetCard(key string) *card.T {
// 	if card, ok := t.obj[key].(*card.T); ok {
// 		return card
// 	}
// 	return nil
// }

// func (t *T) GetToken(key string) *token.T {
// 	if token, ok := t.obj[key].(*token.T); ok {
// 		return token
// 	}
// 	return nil
// }
