package game

import (
	"github.com/zachtaylor/7elements/out"
	"ztaylor.me/cast"
)

type State struct {
	id     string
	r      Stater
	Timer  cast.Duration
	Reacts map[string]string
	Stack  *State
}

func (s *State) ID() string {
	return s.id
}

func (s *State) Name() string {
	return s.r.Name()
}

func (s *State) Seat() string {
	return s.r.Seat()
}

func (s *State) GetNextStater(game *T) Stater {
	return s.r.GetNext(game)
}

func (s *State) Connect(game *T, seat *Seat) {
	log := game.Log().With(cast.JSON{
		"State": game.State,
		"Seat":  seat,
	}).Tag("connect")
	// out.GameData(s.Player, g, s)
	if connector, ok := s.r.(ConnectStater); ok {
		connector.OnConnect(game, seat)
		log.Trace()
	} else {
		log.Trace("none")
	}
}

func (s *State) Disconnect(game *T, seat *Seat) {
	log := game.Log().With(cast.JSON{
		"State": game.State,
		"Seat":  seat,
	}).Tag("disconnect")

	if disconnector, ok := s.r.(DisconnectStater); ok {
		disconnector.OnDisconnect(game, seat)
		log.Trace()
	} else {
		log.Trace("none")
	}
}

func (s *State) Activate(game *T) (events []Stater) {
	log := game.Log().With(cast.JSON{
		"State": game.State,
	}).Tag("activate")

	if activator, ok := s.r.(ActivateStater); ok {
		events = activator.OnActivate(game)
		log.Trace()
	} else {
		log.Trace("none")
	}
	s.Reactivate(game)
	return
}

func (s *State) Reactivate(game *T) {
	game.Log().With(cast.JSON{
		"State": game.State,
	}).Trace("reactivate")
	game.State.Timer = game.Settings.Timeout
	game.State.Reacts = make(map[string]string)
	out.GameState(game, game.State.JSON())
	s.Connect(game, nil)
}

func (s *State) Finish(game *T) (events []Stater) {
	log := game.Log().With(cast.JSON{
		"State": game.State,
	}).Tag("finish")

	if finisher, _ := game.State.r.(FinishStater); finisher != nil {
		events = finisher.Finish(game)
		log.Trace()
	} else {
		log.Trace("none")
	}
	return
}

func (s *State) Request(game *T, seat *Seat, json cast.JSON) {
	log := game.Log().With(cast.JSON{
		"Username": seat.Username,
		"State":    game.State,
	})

	if requester, ok := game.State.r.(RequestStater); ok {
		requester.Request(game, seat, json)
		log.Trace("ok")
	} else {
		log.Trace("none")
	}
}

func (s *State) String() string {
	return `game.State{#` + s.id + `(` + s.Name() + `:` + s.Seat() + `)` + `}`
}

// JSON returns a a representation of game state
func (s *State) JSON() cast.JSON {
	reactsJSON := cast.JSON{}
	for k, v := range s.Reacts {
		reactsJSON[k] = v
	}
	return cast.JSON{
		"id":     s.ID(),
		"seat":   s.Seat,
		"name":   s.Name(),
		"data":   s.r.JSON(),
		"timer":  int(s.Timer.Seconds()),
		"reacts": reactsJSON,
	}
}
