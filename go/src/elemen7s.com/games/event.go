package games

import (
	"sync"
	"time"
	"ztaylor.me/js"
)

type EMode interface {
	Name() string
	Json(*Event, *Game, *Seat) js.Object
	OnActivate(*Event, *Game)
	OnResolve(*Event, *Game)
	OnReceive(*Event, *Game, *Seat, js.Object)
}

type Event struct {
	Id       int
	Username string
	Target   string
	Resp     map[string]string
	time.Duration
	EMode
	sync.Mutex
}

func NewEvent(username string) *Event {
	return &Event{
		Username: username,
		Resp:     make(map[string]string),
		Duration: 5 * time.Minute,
	}
}

func (e *Event) ModeName() string {
	return e.EMode.Name()
}

func (e *Event) Activate(g *Game) {
	g.Active = e
	g.Broadcast(e.Name(), e.EMode.Json(e, g, g.GetSeat(e.Username)))
	log := g.Log().Add("Username", e.Username).Add("Mode", e.EMode.Name())
	log.Debug("games/event: activate")
	e.EMode.OnActivate(e, g)
	if e.EMode.Name() != "start" && e.EMode.Name() != "sunrise" {
		delay(15*time.Second, func() {
			e.Lock()
			defer e.Unlock()
			if g.Active != e {
				return
			}
			if e.CheckPass(g) {
				log.Info("games/event: autopass")
				e.Timeout()
			}
		})
	}
}

func (e *Event) Resolve(g *Game) {
	g.Log().Add("Username", e.Username).Add("Mode", e.EMode.Name()).Debug("resolve")
	e.EMode.OnResolve(e, g)
}

// func (e *Event) chat(game *Game, seat *Seat, msg *wsocks.Message) {
// 	go game.Broadcast("alert", js.Object{
// 		"class":    "tip",
// 		"gameid":   game.Id,
// 		"username": seat.Username,
// 		"message":  m.Data["message"],
// 	})
// }

func (e *Event) Receive(g *Game, s *Seat, j js.Object) {
	e.Lock()
	defer e.Unlock()
	if j["resp"] == "pass" {
		e.RespPass(g, s)
	} else if j["resp"] == "play" {
		TryPlay(e, g, s, j, e.EMode.Name() != "main" || s.Username != e.Username)
	} else {
		e.EMode.OnReceive(e, g, s, j)
	}
}

func (e *Event) RespPass(game *Game, seat *Seat) {
	log := game.Log().Add("Username", seat.Username).Add("Mode", e.EMode.Name())
	if e.Resp[seat.Username] != "" {
		AnimateAlertError(seat, game, "pass", "already recorded")
		log.Add("Val", e.Resp[seat.Username]).Warn("pass: response already recorded")
	} else if e.Resp[seat.Username] == "" {
		log.Debug("pass")
		e.Resp[seat.Username] = "pass"
		game.Broadcast("pass", js.Object{
			"gameid":   game.Id,
			"target":   e.Target,
			"username": seat.Username,
		})
	}

	if e.CheckPass(game) {
		e.Timeout()
	}
}

func (e *Event) SendCatchup(g *Game, seat *Seat) {
	seat.Send(e.Name(), e.Json(g, seat))
	for username, resp := range e.Resp {
		if resp == "pass" {
			AnimatePass(seat, g, username)
		}
	}
}

func (e *Event) CheckPass(g *Game) bool {
	for _, s := range g.Seats {
		if e.Resp[s.Username] == "pass" {
		} else if s.HasCardsInHand() && s.HasActiveElements() {
			return false
		}
	}
	return true
}

func (e *Event) Timeout() {
	e.Duration = time.Second
}

func (e *Event) Json(g *Game, s *Seat) js.Object {
	return e.EMode.Json(e, g, s)
}
