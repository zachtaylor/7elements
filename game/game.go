package game

import (
	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/chat"
	"ztaylor.me/cast"
	"ztaylor.me/keygen"
	"ztaylor.me/log"
)

type T struct {
	id      string
	in      chan *Request
	chat    *chat.Room
	close   chan bool
	Cards   Cards
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
		Cards:   make(Cards),
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

func (game *T) NewState(event Event) *State {
	id := keygen.NewVal()
	return &State{
		id:     id,
		Timer:  game.Runtime.Timeout,
		Reacts: make(map[string]string),
		Event:  event,
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
	if game.in != nil {
		game.in <- r
	}
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

func (game *T) Register(deck *vii.AccountDeck) *Seat {
	log := game.Log().Add("Username", deck.Username)

	if game.Seats[deck.Username] != nil {
		log.Warn("register: username already registered")
		return nil
	}

	seat := game.NewSeat()
	seat.Deck = NewDeck()
	seat.Username = deck.Username
	seat.Deck.Username = deck.Username
	seat.Deck.AccountDeckID = deck.ID
	deckSize := 0

	for cardid, copies := range deck.Cards {
		card, _ := game.Runtime.Root.Cards.Get(cardid)
		if card == nil {
			log.Copy().Add("CardId", cardid).Warn("register: card missing")
			return nil
		}

		for i := 0; i < copies; i++ {
			card := NewCard(card)
			card.Username = deck.Username
			game.RegisterCard(card)
			game.Cards[card.Id] = card
			seat.Deck.Append(card)
		}
		deckSize += copies
	}

	game.Seats[seat.Username] = seat
	log.Add("DeckSize", deckSize).Add("Cards", deck.Cards).Debug("registered seat")
	return seat
}

func (game *T) RegisterCard(card *Card) {
	card.Id = keygen.NewVal()
	for game.Cards[card.Id] != nil {
		card.Id = keygen.NewVal()
	}
	game.Cards[card.Id] = card
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
	close(game.in)
	game.in = nil
	close(game.close)
	game.close = nil
	game.chat.Destroy()
	game.Runtime.logger.Close()
}

// PerspectiveJSON returns JSON representation of a game
func (game *T) PerspectiveJSON(seat *Seat) cast.JSON {
	if game == nil {
		return nil
	}
	seats := cast.JSON{}
	for _, s := range game.Seats {
		seats[s.Username] = s.JSON()
	}
	return cast.JSON{
		"id":       game.ID(),
		"life":     seat.Life,
		"hand":     seat.Hand.JSON(),
		"state":    game.State.JSON(),
		"elements": seat.Elements.JSON(),
		"username": seat.Username,
		"opponent": game.GetOpponentSeat(seat.Username).Username,
		"seats":    seats,
	}
}
