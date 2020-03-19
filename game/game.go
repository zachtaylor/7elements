package game

import (
	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/deck"
	"ztaylor.me/cast"
	"ztaylor.me/cast/charset"
	"ztaylor.me/keygen"
	"ztaylor.me/log"
)

type T struct {
	id      string
	in      chan *Request
	lock    cast.Mutex
	chat    *chat.Room
	close   chan bool
	Objects map[string]interface{}
	Seats   map[string]*Seat
	State   *State
	Runtime *Runtime
}

func New(id string, rt *Runtime) *T {
	game := &T{
		id:      id,
		in:      make(chan *Request),
		chat:    rt.chat.New(`game#`+id, 21),
		close:   make(chan bool),
		Objects: make(map[string]interface{}),
		Seats:   make(map[string]*Seat),
		Runtime: rt,
	}
	return game
}

func (game *T) ID() string {
	return game.id
}

func (game *T) GetChat() *chat.Room {
	return game.chat
}

func (game T) String() string {
	return "Game#" + game.id
}

func (game *T) Log() *log.Entry {
	return game.Runtime.logger.New()
}

func (game *T) NewState(r Stater) *State {
	id := keygen.New(3, charset.AlphaCapitalNumeric, keygen.DefaultSettings.Rand)

	return &State{
		id:     id,
		Timer:  game.Runtime.Timeout,
		Reacts: make(map[string]string),
		R:      r,
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
	game.lock.Lock()
	if game.in != nil {
		game.in <- r
	}
	game.lock.Unlock()
}

func (game *T) Monitor() <-chan *Request {
	return game.in
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

func (game *T) Register(ad *vii.AccountDeck) *Seat {
	log := game.Log().Add("Username", ad.Username)

	if game.Seats[ad.Username] != nil {
		log.Warn("register: username already registered")
		return nil
	}

	seat := game.NewSeat()
	seat.Deck = deck.New()
	seat.Username = ad.Username
	seat.Deck.Username = ad.Username
	seat.Deck.AccountDeckID = ad.ID
	deckSize := 0

	for cardid, copies := range ad.Cards {
		proto, _ := game.Runtime.Root.Cards.Get(cardid)
		if proto == nil {
			log.Copy().Add("CardId", cardid).Warn("register: card missing")
			return nil
		}

		for i := 0; i < copies; i++ {
			card := card.New(proto)
			card.Username = ad.Username
			game.RegisterCard(card)
			seat.Deck.Append(card)
		}
		deckSize += copies
	}

	game.Seats[seat.Username] = seat
	log.Add("DeckSize", deckSize).Add("Cards", ad.Cards).Debug("registered seat")
	return seat
}

func (game *T) RegisterObjectKey() (key string) {
	newkey := func() string {
		return keygen.New(4, charset.AlphaCapitalNumeric, keygen.DefaultSettings.Rand)
	}
	key = newkey()
	for _, ok := game.Objects[key]; ok; {
		key = newkey()
		_, ok = game.Objects[key]
	}
	game.Objects[key] = nil
	return
}

func (game *T) RegisterCard(card *card.T) {
	key := game.RegisterObjectKey()
	card.ID = key
	game.Objects[key] = card
}

func (game *T) RegisterToken(token *Token) {
	key := game.RegisterObjectKey()
	token.ID = key
	game.Objects[key] = token
}

// GetCloser returns the game open chan
func (game *T) Done() chan bool {
	return game.close
}

// Close ends the game, freeing resources
func (game *T) Close() {
	game.Runtime.logger.New().
		// With(log.Fields{}).
		Info("close") // add fields later
	game.lock.Lock()
	close(game.in)
	game.in = nil
	close(game.close)
	game.close = nil
	game.lock.Unlock()
	game.chat.Destroy()
	game.Runtime.logger.Close()
}

// WriteJSON calls WriteJSON(data) for all game seats
func (game *T) WriteJSON(json cast.JSON) {
	for _, seat := range game.Seats {
		seat.WriteJSON(json)
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
