package vii

import (
	"time"

	"ztaylor.me/keygen"
	"ztaylor.me/log"
)

type Game struct {
	Key      string
	Cards    GameCards
	In       chan *GameRequest
	Seats    map[string]*GameSeat
	State    *GameState
	Settings *GameSettings
	*log.Logger
}

func NewGame(settings *GameSettings) *Game {
	return &Game{
		Key:      keygen.NewVal(),
		Cards:    make(GameCards),
		In:       make(chan *GameRequest),
		Seats:    make(map[string]*GameSeat),
		Settings: settings,
		Logger:   log.NewLogger(),
	}
}

func (game Game) String() string {
	return game.Key
}

func (game *Game) Log() log.Log {
	return game.Logger.New()
}

func (game *Game) GetSeat(name string) *GameSeat {
	for k, seat := range game.Seats {
		if k == name {
			return seat
		}
	}
	return nil
}

func (game *Game) GetOpponentSeat(name string) *GameSeat {
	for k, seat := range game.Seats {
		if k != name {
			return seat
		}
	}
	return nil
}

func (game *Game) Register(deck *AccountDeck) *GameSeat {
	log := game.Log().Add("Username", deck.Username)

	if game.Seats[deck.Username] != nil {
		log.Warn("register: username already registered")
		return nil
	}

	seat := game.NewSeat()
	seat.Deck = NewGameDeck()
	seat.Username = deck.Username
	seat.Deck.Username = deck.Username
	seat.Deck.AccountDeckID = deck.ID

	deckSize := 0

	for cardid, copies := range deck.Cards {
		card, _ := CardService.Get(cardid)
		if card == nil {
			log.Clone().Add("CardId", cardid).Warn("register: card missing")
			return nil
		}

		for i := 0; i < copies; i++ {
			card := NewGameCard(card)
			card.Username = deck.Username
			game.RegisterCard(card)
			game.Cards[card.Id] = card
			seat.Deck.Append(card)
		}
		deckSize += copies
		log.Clone().Add("CardId", cardid).Add("Copies", copies).Debug("register card")
	}

	game.Seats[seat.Username] = seat
	log.Add("DeckSize", deckSize).Debug("registered seat")
	return seat
}

func (game *Game) RegisterCard(card *GameCard) {
	card.Id = keygen.NewVal()
	for game.Cards[card.Id] != nil {
		card.Id = keygen.NewVal()
	}
	game.Cards[card.Id] = card
}

func (game *Game) WriteJson(json Json) {
	for _, seat := range game.Seats {
		seat.WriteJson(json)
	}
}

func (game *Game) Json(name string) Json {
	seat := game.GetSeat(name)
	seats := Json{}
	for _, s := range game.Seats {
		seats[s.Username] = s.Json(false)
	}
	return Json{
		"id":       game.Key,
		"life":     seat.Life,
		"hand":     seat.Hand.Json(),
		"state":    game.State.Json(game),
		"elements": seat.Elements.Json(),
		"username": name,
		"opponent": game.GetOpponentSeat(name).Username,
		"seats":    seats,
	}
}

var GameService interface {
	New() *Game
	Get(id string) *Game
	Forget(id string)
	Watch(*Game)
	GetPlayerGames(name string) []string
	GetPlayerSearch(name string) *GameSearch
	StartPlayerSearch(deck *AccountDeck) *GameSearch
}

type GameSearch struct {
	Deck     *AccountDeck
	Start    time.Time
	Done     chan string
	Settings GameSearchSettings
}

type GameSearchSettings struct {
	UseP2P bool
}

func NewGameSearch(deck *AccountDeck) *GameSearch {
	return &GameSearch{
		Deck:     deck,
		Start:    time.Now(),
		Done:     make(chan string),
		Settings: GameSearchSettings{},
	}
}

type GameRequest struct {
	Username string
	Data     Json
}

func (r GameRequest) String() string {
	return r.Username + ":" + r.Data.Sval("event")
}

type GameResults struct {
	Winner string
	Loser  string
}

type GameSettings struct {
	Timeout time.Duration
}

func NewDefaultGameSettings() *GameSettings {
	settings := &GameSettings{}
	settings.Timeout = 7 * time.Minute
	return settings
}

type GameState struct {
	Seat   string
	Timer  time.Duration
	Reacts map[string]string
	Event  GameEvent
}

func NewGameState(seat string, settings *GameSettings, event GameEvent) *GameState {
	return &GameState{
		Seat:   seat,
		Timer:  settings.Timeout,
		Reacts: make(map[string]string),
		Event:  event,
	}
}

func (s *GameState) EventName() string {
	return s.Event.Name()
}

func (s *GameState) Json(game *Game) Json {
	return Json{
		"gameid": game.Key,
		"event":  s.EventName(),
		"seat":   s.Seat,
		"timer":  int(s.Timer.Seconds()),
		"reacts": s.Reacts,
		"data":   s.Event.Json(game),
	}
}

type GameEvent interface {
	// Name is the refferential name of the game event
	Name() string
	// OnStart is called by the engine when the event timer starts
	OnStart(*Game)
	// OnReconnect is called by the engine whenever a GameSeat (re)joins
	OnReconnect(*Game, *GameSeat)
	// NextEvent is called by the engine when this event must pass on
	// Returns the next GameEvent
	NextEvent(*Game) GameEvent
	// Json() create a representation of this GameState extra data
	Json(*Game) Json
	// Receive is called when data is sent to this GameState
	Receive(*Game, *GameSeat, Json)
}
