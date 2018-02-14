package games

import (
	"elemen7s.com/cards"
	"elemen7s.com/cards/texts"
	"elemen7s.com/decks"
	"fmt"
	"time"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

type Game struct {
	Id       int
	Cards    map[int]*Card
	Delay    time.Duration
	Timeout  time.Duration
	Events   map[int]Event
	Timeline chan *Event
	*log.Logger
	Active *Event
	*TurnClock
	Seats map[string]*Seat
	*Results
}

func New() *Game {
	id := NewGameId()
	logger := log.NewLogger()
	logger.SetLevel("debug")
	logger.SetFile(fmt.Sprintf("log/game-%d.log", id))
	return &Game{
		Id:       id,
		Cards:    make(map[int]*Card),
		Delay:    1 * time.Second,
		Timeout:  10 * time.Minute,
		Events:   make(map[int]Event),
		Timeline: make(chan *Event),
		Logger:   logger,
		// Active:   nil,
		// TurnClock: nil,
		Seats: make(map[string]*Seat),
		// Results: nil,
	}
}

func (g *Game) Log() log.Log {
	return g.Logger.New()
}

func (g *Game) AddSeat(s *Seat) {
	if g.Seats[s.Username] != nil {
		g.Log().Add("Username", s.Username).Warn("games.AddSeat: username already stored")
	} else if len(g.Events) > 0 {
		g.Log().Add("Username", s.Username).Warn("games.AddSeat: game already in progress")
	} else {
		g.Seats[s.Username] = s
	}
}

func (g *Game) RegisterToken(username string, c *Card) {
	log := g.Log().Add("Username", username).Add("CardName", c.Text.Name)
	if s := g.GetSeat(username); s == nil {
		log.Warn("games.RegisterToken: seat missing")
	} else if c.IsRegistered() {
		log.Add("gcid", c.Id).Warn("games.RegisterToken: card is already registered")
	} else {
		c.Id = len(g.Cards)
		g.Cards[c.Id] = c
		s.Alive[c.Id] = c
	}
}

func (g *Game) Register(deck *decks.Deck, lang string) *Seat {
	log := g.Log().Add("Username", deck.Username)

	if g.Seats[deck.Username] != nil {
		log.Warn("games.Register: username already registered")
		return nil
	}

	s := newSeat()
	s.Username = deck.Username
	s.Deck = NewDeck()
	s.Deck.DeckId = deck.Id

	deckSize := 0

	for cardid, copies := range deck.Cards {
		card := cards.Test(cardid)
		text := texts.GetAll(lang)[cardid]
		if card == nil {
			log.Clone().Add("CardId", cardid).Warn("register: card missing")
			return nil
		}
		if text == nil {
			log.Clone().Add("CardId", cardid).Warn("register: text missing")
			return nil
		}

		for i := 0; i < copies; i++ {
			c := NewCard(card, text)
			c.Id = len(g.Cards)
			c.Username = deck.Username
			g.Cards[c.Id] = c
			s.Deck.Append(c)
		}
		deckSize += copies
		log.Clone().Add("CardId", cardid).Add("Copies", copies).Debug("register card")
	}

	g.Seats[deck.Username] = s
	log.Add("DeckSize", deckSize).Debug("registered seat")
	return s
}

func (g *Game) GetSeat(username string) *Seat {
	return g.Seats[username]
}

func (g *Game) SendCatchup(seat *Seat) {
	seat.Send("game", g.JsonWithPerspective(seat))
	seat.Send(g.Active.Name(), g.Active.Json(g.Active, g, seat))
}

func (g *Game) Broadcast(name string, json js.Object) {
	for _, seat := range g.Seats {
		if player := seat.Player; player != nil {
			go player.Send(name, json)
		}
	}
}

func (g *Game) Receive(username string, j js.Object) {
	if g.Active.Name() != j["event"] {
		g.Log().Add("Seat", username).Add("Event", j["event"]).Add("Active", g.Active.Name()).Warn("receive: out of sync")
	} else {
		g.Active.Receive(g, g.GetSeat(username), j)
	}
}

func (g *Game) Start() {
	go func() {
		for {
			e := g.Active
			start := time.Now()
			select {
			case rcv := <-g.Timeline:
				if rcv != nil {
					rcv.Activate(g)
				} else {
					e.Duration = time.Second
				}
			case <-time.After(time.Second):
				e.Duration -= time.Now().Sub(start)
				if e.Duration < time.Second {
					e.Resolve(g)
				}
			}
		}
	}()
}

func (g *Game) TimelineJoin(e *Event) {
	go func() {
		g.Timeline <- e
	}()
}

func (g *Game) Win(s *Seat) {
	losers := make([]string, 0)
	for _, s2 := range g.Seats {
		if s2.Username != s.Username {
			losers = append(losers, s2.Username)
		}
	}
	g.Results = &Results{
		Losers:  losers,
		Winners: []string{s.Username},
	}
	End(g)
}

func (g *Game) Lose(s *Seat) {
	winners := make([]string, 0)
	for _, s2 := range g.Seats {
		if s2.Username != s.Username {
			winners = append(winners, s2.Username)
		}
	}
	g.Results = &Results{
		Losers:  []string{s.Username},
		Winners: winners,
	}
	End(g)
}

func (g *Game) Json() js.Object {
	event := g.Active
	data := js.Object{
		"gameid": g.Id,
		"timer":  int(event.Duration.Seconds()),
		"mode":   event.ModeName(),
	}

	seatdata := js.Object{}
	for _, seat := range g.Seats {
		seatdata[seat.Username] = seat.Json()
	}
	data["seats"] = seatdata

	return data
}

func (game *Game) JsonWithPerspective(s *Seat) js.Object {
	json := game.Json()
	json["username"] = s.Username
	json["hand"] = s.Hand.Json()
	return json
}
