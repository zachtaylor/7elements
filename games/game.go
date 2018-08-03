package games

// import (
// 	"github.com/zachtaylor/7tcg/chat"
// 	"fmt"
// 	"time"
// 	"ztaylor.me/js"
// 	"ztaylor.me/keygen"
// 	"ztaylor.me/log"
// )

// const EventTimeout = 5 * time.Minute

// type Game struct {
// 	Id       int
// 	Cards    map[string]*vii.GameCard
// 	Active   *Event
// 	Events   map[int]Event
// 	Timeline chan *Event
// 	Chat     chat.Channel
// 	Seats    map[string]*vii.GameSeat
// 	*log.Logger
// 	*TurnClock
// 	*Results
// }

// func New() *vii.Game {
// 	id := NewGameId()
// 	logger := log.NewLogger()
// 	logger.SetLevel("debug")
// 	go func() {
// 		logger.SetFile(fmt.Sprintf("log/game-%d.log", id))
// 	}()
// 	<-time.After(time.Second)
// 	logger.New().Add("GameId", id).Info("game started")
// 	return &Game{
// 		Id:       id,
// 		Cards:    make(map[string]*vii.GameCard),
// 		Events:   make(map[int]Event),
// 		Timeline: make(chan *Event),
// 		Chat:     chat.NewChannel(fmt.Sprintf("game#%d", id)),
// 		Logger:   logger,
// 		// Active:   nil,
// 		// TurnClock: nil,
// 		Seats: make(map[string]*vii.GameSeat),
// 		// Results: nil,
// 	}
// }

// func (g *vii.Game) Log() log.Log {
// 	return g.Logger.New()
// }

// func (g *vii.Game) NewCardId() string {
// 	key := keygen.NewVal()
// 	for g.Cards[key] != nil {
// 		key = keygen.NewVal()
// 	}
// 	return key
// }

// func (g *vii.Game) AddSeat(s *vii.GameSeat) {
// 	if g.Seats[s.Username] != nil {
// 		g.Log().Add("Username", s.Username).Warn("games.AddSeat: username already stored")
// 	} else if len(g.Events) > 0 {
// 		g.Log().Add("Username", s.Username).Warn("games.AddSeat: game already in progress")
// 	} else {
// 		g.Seats[s.Username] = s
// 	}
// }

// func (g *vii.Game) RegisterToken(username string, c *vii.GameCard) {
// 	log := g.Log().Add("Username", username).Add("CardName", c.CardText.Name)
// 	if s := g.GetSeat(username); s == nil {
// 		log.Warn("games.RegisterToken: seat missing")
// 	} else if c.IsRegistered() {
// 		log.Add("gcid", c.Id).Warn("games.RegisterToken: card is already registered")
// 	} else {
// 		c.IsToken = true
// 		c.Username = username
// 		c.Id = g.NewCardId()
// 		g.Cards[c.Id] = c
// 		s.Alive[c.Id] = c
// 	}
// }

// func (g *vii.Game) Register(deck *vii.AccountDeck, lang string) *vii.GameSeat {
// 	log := g.Log().Add("Username", deck.Username)

// 	if g.Seats[deck.Username] != nil {
// 		log.Warn("games.Register: username already registered")
// 		return nil
// 	}

// 	s := newSeat()
// 	s.Deck = vii.NewGameDeck()
// 	s.Username = deck.Username
// 	s.Deck.Username = deck.Username
// 	s.Deck.AccountDeckId = deck.Id
// 	s.Deck.AccountDeckVersion = deck.Version

// 	deckSize := 0

// 	for cardid, copies := range deck.Cards {
// 		card, _ := vii.CardService.GetCard(cardid)
// 		if card == nil {
// 			log.Clone().Add("CardId", cardid).Warn("register: card missing")
// 			return nil
// 		}
// 		text, err := vii.CardTextService.GetCardText(lang, cardid)
// 		if text == nil {
// 			log.Clone().Add("CardId", cardid).Add("Error", err).Warn("register: text missing")
// 			return nil
// 		}

// 		for i := 0; i < copies; i++ {
// 			c := vii.NewGameCard(card, text)
// 			c.Id = g.NewCardId()
// 			c.Username = deck.Username
// 			g.Cards[c.Id] = c
// 			s.Deck.Append(c)
// 		}
// 		deckSize += copies
// 		log.Clone().Add("CardId", cardid).Add("Copies", copies).Debug("register card")
// 	}

// 	g.Seats[deck.Username] = s
// 	log.Add("DeckSize", deckSize).Debug("registered seat")
// 	return s
// }

// func (g *vii.Game) GetSeat(username string) *vii.GameSeat {
// 	return g.Seats[username]
// }

// func (g *vii.Game) SendCatchup(seat *vii.GameSeat) {
// 	seat.Send("game", g.StateJson(seat.Username))
// 	g.Active.SendCatchup(g, seat)
// }

// func (g *vii.Game) Broadcast(name string, json js.Object) {
// 	for _, seat := range g.Seats {
// 		if player := seat.Player; player != nil {
// 			go player.Send(name, json)
// 		}
// 	}
// }

// func (g *vii.Game) Receive(username string, json js.Object) {
// 	switch json["event"] {
// 	case "pass":
// 		TryPass(g, g.GetSeat(username), json)
// 		break
// 	case "chat":
// 		BroadcastAnimateAlertChat(g, username, json.Sval("message"))
// 		break
// 	case "play":
// 		TryPlay(g, g.GetSeat(username), json, g.Active.ModeName() != "main" || username != g.Active.Username)
// 		break
// 	case "trigger":
// 		TryTrigger(g, g.GetSeat(username), json)
// 		break
// 	case g.Active.Name():
// 		g.Active.Receive(g, g.GetSeat(username), json)
// 		break
// 	default:
// 		g.Log().Add("Seat", username).Add("Event", json["event"]).Add("Active", g.Active.Name()).Warn("receive: out of sync")
// 	}
// }

// func (g *vii.Game) PowerScript(seat *vii.GameSeat, power *vii.Power, target interface{}) {
// 	if script := Scripts[power.Script]; script == nil {
// 		log.Add("Script", power.Script).Warn("script missing")
// 	} else {
// 		script(g, seat, target)
// 	}
// }

// func (g *vii.Game) Start() {
// 	go g.Watch()
// }

// func (g *vii.Game) Watch() {
// 	for {
// 		e := g.Active
// 		start := time.Now()
// 		select {
// 		case rcv, ok := <-g.Timeline:
// 			if rcv != nil {
// 				rcv.Activate(g)
// 			} else if ok {
// 				e.Duration = time.Second
// 				g.Log().Add("Mode", g.Active.Name()).Debug("games/watch: close short circuit")
// 			} else {
// 				g.Log().Add("Mode", g.Active.Name()).Debug("games/watch: finished")
// 				return
// 			}
// 		case <-time.After(time.Second):
// 			e.Duration -= time.Now().Sub(start)
// 			if e.Duration < time.Second {
// 				e.Resolve(g)
// 			}
// 		}
// 	}
// }

// func (g *vii.Game) TimelineJoin(e *Event) {
// 	go func() {
// 		g.Timeline <- e
// 	}()
// }

// func (g *vii.Game) Win(s *vii.GameSeat) {
// 	losers := make([]string, 0)
// 	for _, s2 := range g.Seats {
// 		if s2.Username != s.Username {
// 			losers = append(losers, s2.Username)
// 		}
// 	}
// 	g.Results = &Results{
// 		Losers:  losers,
// 		Winners: []string{s.Username},
// 	}
// 	End(g)
// }

// func (g *vii.Game) Lose(s *vii.GameSeat) {
// 	winners := make([]string, 0)
// 	for _, s2 := range g.Seats {
// 		if s2.Username != s.Username {
// 			winners = append(winners, s2.Username)
// 		}
// 	}
// 	g.Results = &Results{
// 		Losers:  []string{s.Username},
// 		Winners: winners,
// 	}
// 	End(g)
// }

// func (g *vii.Game) StateJson(username string) js.Object {
// 	var json js.Object
// 	if seat := g.GetSeat(username); seat != nil {
// 		json = seat.Json(true)
// 		opponentsdata := make([]string, 0)
// 		for _, seat2 := range g.Seats {
// 			if seat2.Username != username {
// 				opponentsdata = append(opponentsdata, seat2.Username)
// 			}
// 		}
// 		json["opponents"] = opponentsdata
// 	} else {
// 		json = js.Object{
// 			"error": "username not found: " + username,
// 		}
// 	}

// 	json["gameid"] = g.Id
// 	json["timer"] = int(g.Active.Duration.Seconds())
// 	json["mode"] = g.Active.ModeName()

// 	seatdata := js.Object{}
// 	for _, seat := range g.Seats {
// 		seatdata[seat.Username] = seat.Json(false)
// 	}
// 	json["seats"] = seatdata

// 	return json
// }
