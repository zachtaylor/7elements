package game

import (
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/deck"
	"github.com/zachtaylor/7elements/game/seat"
	"taylz.io/http/websocket"
	"taylz.io/log"
)

type T struct {
	id    string
	in    chan *Request
	close chan bool
	eng   Engine
	chat  *chat.Room
	log   log.Writer
	obj   map[string]interface{}
	rules Rules
	Seats *seat.List
	State *State
}

func New(id string, engine Engine, chat *chat.Room, log log.Writer, rules Rules) *T {
	return &T{
		id:    id,
		in:    make(chan *Request),
		close: make(chan bool),
		eng:   engine,
		chat:  chat,
		log:   log,
		obj:   make(map[string]interface{}),
		rules: rules,
		Seats: seat.NewList(),
	}
}

func (*T) NewWinLoss(winner, loser string) Resulter {
	return &WinLossResult{
		winner: winner,
		loser:  loser,
	}
}

func (*T) NewDraw() Resulter { return &DrawResult{} }

func (game *T) NewState(phase Phaser) *State { return NewState(game.rules.Timeout, phase) }

func (game *T) ID() string { return game.id }

func (game *T) Log() log.Writer { return game.log.New() }

func (game *T) Engine() Engine { return game.eng }

func (game *T) Rules() Rules { return game.rules }

func (game *T) String() string { return "Game#" + game.id }

func (game *T) Close() {
	if game.in != nil {
		close(game.in)
		game.in = nil
		close(game.close)
		game.close = nil
	}
	game.log.Close()
	game.chat.Destroy()
}

// RequestChan returns the raw ordered game input
func (game *T) RequestChan() <-chan *Request { return game.in }

// Request starts a go routine to call RequestSync
func (game *T) Request(username string, uri string, data websocket.MsgData) {
	go game.RequestSync(&Request{
		Username: username,
		URI:      uri,
		Data:     data,
	})
}

// RequestSync waits to request the game engine
func (game *T) RequestSync(r *Request) {
	if game.in != nil {
		game.in <- r
	}
}

func (game *T) Chat(source, message string) { game.chat.Add(source, message) }

func (game *T) Phase() string { return game.State.Phase.Name() }

func (game *T) Data(seat *seat.T) websocket.MsgData {
	return websocket.MsgData{
		"id":       game.ID(),
		"stateid":  game.State.ID(),
		"username": seat.Username,
		"seats":    game.Seats.Keys(),
	}
}

func (t *T) Register(deck *deck.T, writer websocket.Writer) *seat.T {
	log := t.Log().Add("Username", deck.User)

	if t.Seats.Get(deck.User) != nil {
		log.Warn("username already registered")
		return nil
	}

	seat := seat.New(t.rules.StartingLife, deck, writer)

	for _, card := range seat.Deck.Cards {
		t.RegisterCard(card)
	}

	t.Seats.Add(seat)
	log.With(websocket.MsgData{
		"DeckSize": seat.Deck.Count(),
		"Username": deck.User,
	}).Info("register")
	return seat
}
