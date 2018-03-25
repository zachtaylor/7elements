package games

import (
	"elemen7s.com"
	"elemen7s.com/chat"
	"elemen7s.com/decks"
	"fmt"
	"time"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

const EventTimeout = 5 * time.Minute

type Game struct {
	Id       int
	Cards    map[int]*Card
	Events   map[int]Event
	Timeline chan *Event
	Chat     chat.Channel
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
	go func() {
		logger.SetFile(fmt.Sprintf("log/game-%d.log", id))
	}()
	<-time.After(time.Second)
	logger.New().Add("GameId", id).Info("game started")
	return &Game{
		Id:       id,
		Cards:    make(map[int]*Card),
		Events:   make(map[int]Event),
		Timeline: make(chan *Event),
		Chat:     chat.NewChannel(fmt.Sprintf("game#%d", id)),
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
	log := g.Log().Add("Username", username).Add("CardName", c.CardText.Name)
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
		card, _ := vii.CardService.GetCard(cardid)
		if card == nil {
			log.Clone().Add("CardId", cardid).Warn("register: card missing")
			return nil
		}
		text, err := vii.CardTextService.GetCardText(lang, cardid)
		if text == nil {
			log.Clone().Add("CardId", cardid).Add("Error", err).Warn("register: text missing")
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
	seat.Send("game", g.StateJson(seat.Username))
	g.Active.SendCatchup(g, seat)
}

func (g *Game) Broadcast(name string, json js.Object) {
	for _, seat := range g.Seats {
		if player := seat.Player; player != nil {
			go player.Send(name, json)
		}
	}
}

func (g *Game) Receive(username string, json js.Object) {
	switch json["event"] {
	case "chat":
		BroadcastAnimateAlertChat(g, username, json.Sval("message"))
		break
	case g.Active.Name():
		g.Active.Receive(g, g.GetSeat(username), json)
		break
	default:
		g.Log().Add("Seat", username).Add("Event", json["event"]).Add("Active", g.Active.Name()).Warn("receive: out of sync")
	}
}

func (g *Game) Start() {
	go g.Watch()
}

func (g *Game) Watch() {
	for {
		e := g.Active
		start := time.Now()
		select {
		case rcv, ok := <-g.Timeline:
			if rcv != nil {
				rcv.Activate(g)
			} else if ok {
				e.Duration = time.Second
				g.Log().Add("Mode", g.Active.Name()).Debug("games/watch: close short circuit")
			} else {
				g.Log().Add("Mode", g.Active.Name()).Debug("games/watch: finished")
				return
			}
		case <-time.After(time.Second):
			e.Duration -= time.Now().Sub(start)
			if e.Duration < time.Second {
				e.Resolve(g)
			}
		}
	}
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

func (g *Game) StateJson(username string) js.Object {
	var json js.Object
	if seat := g.GetSeat(username); seat != nil {
		json = seat.Json(true)
		opponentsdata := make([]string, 0)
		for _, seat2 := range g.Seats {
			if seat2.Username != username {
				opponentsdata = append(opponentsdata, seat2.Username)
			}
		}
		json["opponents"] = opponentsdata
	} else {
		json = js.Object{
			"error": "username not found: " + username,
		}
	}

	json["gameid"] = g.Id
	json["timer"] = int(g.Active.Duration.Seconds())
	json["mode"] = g.Active.ModeName()

	seatdata := js.Object{}
	for _, seat := range g.Seats {
		seatdata[seat.Username] = seat.Json(false)
	}
	json["seats"] = seatdata

	return json
}
