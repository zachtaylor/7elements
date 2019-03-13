package game

import (
	"github.com/zachtaylor/7elements"
	"ztaylor.me/keygen"
	"ztaylor.me/log"
)

type T struct {
	id       string
	in       chan *Request
	Cards    Cards
	Seats    map[string]*Seat
	State    *State
	Settings *Settings
	Logger   *log.Logger
}

func New(id string, logger *log.Logger, settings *Settings) *T {
	return &T{
		id:       id,
		in:       make(chan *Request),
		Cards:    make(Cards),
		Seats:    make(map[string]*Seat),
		Settings: settings,
		Logger:   logger,
	}
}

func (game *T) ID() string {
	return game.id
}

func (game T) String() string {
	return "T#" + game.id
}

func (game *T) Log() log.Log {
	return game.Logger.New()
}

func (game *T) NewState(seat string, event Event) *State {
	id := keygen.NewVal()
	return &State{
		id:     id,
		Seat:   seat,
		Timer:  game.Settings.Timeout,
		Reacts: make(map[string]string),
		Event:  event,
	}
}

// Request buffers a *game.Request, returning immediately
func (game *T) Request(username string, uri string, data vii.Json) {
	go game.request(&Request{
		Username: username,
		URI:      uri,
		Data:     data,
	})
}
func (game *T) request(r *Request) {
	game.in <- r
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
		card, _ := vii.CardService.Get(cardid)
		if card == nil {
			log.Clone().Add("CardId", cardid).Warn("register: card missing")
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

func (game *T) WriteJson(json vii.Json) {
	for _, seat := range game.Seats {
		seat.WriteJson(json)
	}
}

func (game *T) Json(name string) vii.Json {
	seat := game.GetSeat(name)
	seats := vii.Json{}
	for _, s := range game.Seats {
		seats[s.Username] = s.Json(false)
	}
	return vii.Json{
		"id":       game.id,
		"life":     seat.Life,
		"hand":     seat.Hand.Json(),
		"state":    game.State.Json(game),
		"elements": seat.Elements.Json(),
		"username": name,
		"opponent": game.GetOpponentSeat(name).Username,
		"seats":    seats,
	}
}

// Stop ends the game and removes it from the service
func (game *T) Stop() {
	Service.Forget(game.id)
	log.Protect(func() {
		close(game.in)
	})
}
