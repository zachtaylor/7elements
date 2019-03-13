package engine

import (
	"time"

	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/log"
)

// Watch starts a goroutine to run the engine
func Watch(game *game.T) {
	go game.Log().Protect(func() {
		Run(game)
	})
}

func Run(g *game.T) {
	if g == nil {
		log.Error("engine/run: game is nil")
		return
	} else if len(g.Seats) < 2 {
		log.Add("Seats", len(g.Seats)).Error("engine/run: game missing participants")
		return
	}

	// bootstrap
	p1 := g.GetOpponentSeat("").Username
	g.State = g.NewState(p1, Start(g))
	activate(g)
	one2nd := time.NewTicker(time.Second)
	var tStart time.Time

	// ticker
	for {
		tStart = time.Now()
		select {
		case r, ok := <-g.Monitor():
			log := g.Log().Add("State", g.State.EventName())
			if !ok {
				log.Debug("timeline finished")
				return
			}
			log.Clone().Add("Username", r.Username).Add("URI", r.URI).Debug("request")
			if state := Request(g, g.GetSeat(r.Username), r.URI, r.Data); state != nil {
				log.Debug("state stack")
				g.State = state
				animate.GameState(g)
				continue
			}
		case <-one2nd.C:
		}

		g.State.Timer -= time.Now().Sub(tStart)

		if len(g.State.Reacts) == len(g.Seats) {
			g.Log().Add("State", g.State.ID()).Debug("resolve")
			resolve(g)
		} else if g.State.Timer < time.Second {
			g.Log().Add("State", g.State.ID()).Add("Event", g.State.EventName()).Warn("time ran out")
			resolve(g)
		}
	}
}

func resolve(g *game.T) {
	g.State.Timer = 0
	finish(g)
	if state := stackpop(g); state != nil {
		g.State = state
		animate.GameState(g)
		return
	} else if state := g.State.Event.GetNext(g); state == nil {
		g.Log().Add("State", g.State.ID()).Error("engine/resolve: GetNext() failed")
	} else {
		g.State = state
		animate.GameState(g)
		activate(g)
	}
}

func finish(g *game.T) {
	log := g.Log().Add("State", g.State)
	if finisher, _ := g.State.Event.(game.FinishEventer); finisher == nil {
		log.Debug("engine/finish: no finisher")
	} else {
		log.Debug("engine/finish")
		finisher.Finish(g)
	}
}

func stackpop(g *game.T) (s *game.State) {
	log := g.Log().Add("State", g.State)
	if stacker, _ := g.State.Event.(game.StackEventer); stacker == nil {
		log.Debug("engine/stackpop: empty")
	} else if state := stacker.GetStack(g); state == nil {
		log.Warn("engine/stackpop: failed")
	} else {
		log.Debug("engine/stackpop")
		s = state
	}
	return
}

func activate(g *game.T) {
	if activator, _ := g.State.Event.(game.ActivateEventer); activator == nil {
		g.Log().Add("State", g.State).Debug("engine/activate: no activator")
	} else {
		g.Log().Add("State", g.State).Debug("engine/activate")
		activator.OnActivate(g)
	}
}
