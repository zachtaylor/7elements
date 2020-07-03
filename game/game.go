package game

import (
	"time"

	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/deck"
	"github.com/zachtaylor/7elements/out"
	"ztaylor.me/cast"
	"ztaylor.me/cast/charset"
	"ztaylor.me/keygen"
	"ztaylor.me/log"
)

type T struct {
	Settings Settings
	id       string
	in       chan *Request
	close    chan bool
	chanμ    cast.Mutex
	obj      map[string]interface{}
	Seats    map[string]*Seat
	State    *State
}

func New(settings Settings, id string) *T {
	game := &T{
		Settings: settings,
		id:       id,
		in:       make(chan *Request),
		close:    make(chan bool),
		obj:      make(map[string]interface{}),
		Seats:    make(map[string]*Seat),
	}
	return game
}

func (game *T) ID() string      { return game.id }
func (game T) String() string   { return "Game#" + game.id }
func (game *T) Log() *log.Entry { return game.Settings.Logger.New() }

func (game *T) NewState(r Stater, d time.Duration) *State {
	id := keygen.New(3, charset.AlphaCapitalNumeric, keygen.DefaultSettings.Rand)

	return &State{
		id:    id,
		r:     r,
		Timer: d,
		// Timer:  game.Runtime.Timeout,
		Reacts: make(map[string]string),
	}
}

// Request buffers a *game.Request, returning immediately
func (game *T) Request(username string, uri string, data cast.JSON) {
	go game.request(&Request{
		Username: username,
		URI:      uri,
		Data:     data,
	})
}
func (game *T) request(r *Request) {
	game.chanμ.Lock()
	if game.in != nil {
		game.in <- r
	}
	game.chanμ.Unlock()
}

func (game *T) Monitor() <-chan *Request {
	return game.in
}

// Done returns the game open chan
func (game *T) Done() chan bool {
	return game.close
}

// Close ends the game, freeing resources
func (game *T) Close() {
	// game.Runtime.logger.New().
	// 	// With(log.Fields{}).
	// 	Info("close") // add fields later
	game.chanμ.Lock()
	if game.in != nil {
		close(game.in)
		game.in = nil
		close(game.close)
		game.close = nil
	}
	game.chanμ.Unlock()
	game.Settings.Logger.Close()
	game.Settings.Chat.Destroy()
}

// Send broadcoasts a message to all players
func (game *T) Send(uri string, data cast.JSON) {
	for _, seat := range game.Seats {
		seat.Player.Send(uri, data)
	}
}

func (game *T) GetSeat(name string) *Seat {
	for k, seat := range game.Seats {
		if k == name {
			return seat
		}
	}
	return nil
}

func (game *T) GetOpponentSeat(name string) *Seat {
	for k, seat := range game.Seats {
		if k != name {
			return seat
		}
	}
	return nil
}

func (game *T) Register(deck *deck.T) *Seat {
	log := game.Log().Add("Username", deck.User)

	if game.Seats[deck.User] != nil {
		log.Warn("register: username already registered")
		return nil
	}

	seat := game.NewSeat()
	seat.Username = deck.User
	seat.Deck = deck

	for _, card := range seat.Deck.Cards {
		game.RegisterCard(card)
	}

	game.Seats[seat.Username] = seat
	log.With(cast.JSON{
		"DeckSize": seat.Deck.Count(),
		"Username": deck.User,
	}).Info(seat.Deck.Proto.Cards)
	return seat
}

func (game *T) RegisterObjectKey() (key string) {
	newkey := func() string {
		return keygen.New(4, charset.AlphaCapitalNumeric, keygen.DefaultSettings.Rand)
	}
	key = newkey()
	for _, ok := game.obj[key]; ok; {
		key = newkey()
		_, ok = game.obj[key]
	}
	game.obj[key] = nil
	return
}

func (game *T) RegisterCard(card *card.T) {
	key := game.RegisterObjectKey()
	card.ID = key
	game.obj[key] = card
}

func (game *T) RegisterToken(token *Token) {
	key := game.RegisterObjectKey()
	token.ID = key
	game.obj[key] = token
}

func (game *T) GetCard(key string) *card.T {
	obj := game.obj[key]
	if card, ok := obj.(*card.T); ok {
		return card
	}
	return nil
}

func (game *T) GetToken(key string) *Token {
	obj := game.obj[key]
	if token, ok := obj.(*Token); ok {
		return token
	}
	return nil
}

func (game *T) ResolveState() {
	log := game.Log().Add("State", game.State)
	game.State.Timer = 0
	states := game.State.Finish(game) // combine states

	if game.State.Stack != nil {
		log.Add("New", game.State).Debug("stackpop")
		game.State = game.State.Stack
		game.State.Reactivate(game)
	} else {
		log.Add("New", game.State).Debug("getnext")
		game.State = game.NewState(game.State.GetNextStater(game), game.Settings.Timeout)
		states = append(states, game.State.Activate(game)...) // combine states
	}

	game.Stack(states) // stack new states
	out.GameState(game, game.State.JSON())
}

// Stack adds new States as Stacked States
func (game *T) Stack(stack []Stater) {
	if len(stack) < 1 {
		return
	}
	log := game.Log().With(cast.JSON{
		"State": game.State,
		"Stack": stack,
	})
	log.Trace()
	next := make([]Stater, 0)
	for _, r := range stack {
		state := game.NewState(r, game.Settings.Timeout)
		state.Stack = game.State
		game.State = state

		if addnext := game.State.Activate(game); len(addnext) > 0 {
			game.Log().With(cast.JSON{
				"State": game.State,
				"Stack": stack,
			}).Debug("activate trigger")
			next = append(next, addnext...)
		}
	}
	if len(next) > 0 {
		log.Debug("Hold my cards, lads, I'm going in!")
		game.Stack(next)
	}
}

// JSON returns JSON representation of a game
func (game *T) JSON(seat *Seat) cast.JSON {
	if game == nil {
		return nil
	}
	seats := cast.JSON{}
	for _, s := range game.Seats {
		seats[s.Username] = s.JSON()
	}
	return cast.JSON{
		"id": game.ID(),
		// "life":     seat.Life,
		"hand":  seat.Hand.JSON(),
		"state": game.State.JSON(),
		// "elements": seat.Elements.JSON(),
		"username": seat.Username,
		// "opponent": game.GetOpponentSeat(seat.Username).Username,
		"seats": seats,
	}
}
