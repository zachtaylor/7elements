package games

// import (
// 	"github.com/zachtaylor/7tcg"
// 	"sync"
// 	"time"
// 	"ztaylor.me/js"
// )

// type EMode interface {
// 	Name() string
// 	Json(*Event, *vii.Game, *vii.GameSeat) js.Object
// 	OnActivate(*Event, *vii.Game)
// 	OnResolve(*Event, *vii.Game)
// 	OnReceive(*Event, *vii.Game, *vii.GameSeat, js.Object)
// 	OnSendCatchup(*Event, *vii.Game, *vii.GameSeat)
// }

// type Event struct {
// 	Id       int
// 	Username string
// 	Target   string
// 	Resp     map[string]string
// 	time.Duration
// 	EMode
// 	sync.Mutex
// }

// func NewEvent(username string) *Event {
// 	return &Event{
// 		Username: username,
// 		Resp:     make(map[string]string),
// 		Duration: EventTimeout,
// 	}
// }

// func (e *Event) ModeName() string {
// 	return e.EMode.Name()
// }

// func (e *Event) Activate(g *vii.Game) {
// 	g.Active = e
// 	g.Broadcast(e.Name(), e.EMode.Json(e, g, g.GetSeat(e.Username)))
// 	log := g.Log().Add("Username", e.Username).Add("Mode", e.EMode.Name())
// 	log.Info("games/event: activate")
// 	e.EMode.OnActivate(e, g)
// 	if e.EMode.Name() != "start" && e.EMode.Name() != "sunrise" {
// 		delay(15*time.Second, func() {
// 			e.Lock()
// 			defer e.Unlock()
// 			if g.Active != e {
// 				return
// 			}
// 			if e.CheckPass(g) {
// 				g.Log().Add("Target", e.Target).Info("games/event: autopass")
// 				e.Timeout()
// 			}
// 		})
// 	}
// }

// func (e *Event) Resolve(g *vii.Game) {
// 	g.Log().Add("Username", e.Username).Add("Mode", e.EMode.Name()).Debug("resolve")
// 	e.EMode.OnResolve(e, g)
// }

// func (e *Event) Receive(g *vii.Game, s *vii.GameSeat, j js.Object) {
// 	e.Lock()
// 	defer e.Unlock()
// 	e.EMode.OnReceive(e, g, s, j)
// }

// func (e *Event) RespPass(game *vii.Game, seat *vii.GameSeat) {
// 	log := game.Log().Add("Username", seat.Username).Add("Mode", e.EMode.Name())
// 	if e.Resp[seat.Username] != "" {
// 		AnimateAlertError(seat, game, "pass", "already recorded")
// 		log.Add("Val", e.Resp[seat.Username]).Warn("pass: response already recorded")
// 	} else if e.Resp[seat.Username] == "" {
// 		log.Debug("pass")
// 		e.Resp[seat.Username] = "pass"
// 		game.Broadcast("pass", js.Object{
// 			"gameid":   game.Id,
// 			"target":   e.Target,
// 			"username": seat.Username,
// 		})
// 	}

// 	if e.CheckPass(game) {
// 		e.Timeout()
// 	}
// }

// func (e *Event) SendCatchup(g *vii.Game, s *vii.GameSeat) {
// 	s.Send(e.Name(), e.Json(g, s))
// 	e.EMode.OnSendCatchup(e, g, s)
// 	for username, resp := range e.Resp {
// 		if resp == "pass" {
// 			AnimatePass(s, g, username)
// 		}
// 	}
// }

// func (e *Event) CheckPass(g *vii.Game) bool {
// 	for _, s := range g.Seats {
// 		if e.Resp[s.Username] == "pass" {
// 		} else if s.HasCardsInHand() && s.HasActiveElements() {
// 			return false
// 		}
// 	}
// 	return true
// }

// func (e *Event) Timeout() {
// 	e.Duration = time.Second
// }

// func (e *Event) Json(g *vii.Game, s *vii.GameSeat) js.Object {
// 	return e.EMode.Json(e, g, s)
// }
