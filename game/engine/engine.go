package engine

import (
	"time"

	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/log"
)

// Run the engine
func Run(g *game.T) {
	tStart := time.Now()
	one2nd := time.NewTicker(time.Second)
	var tStep time.Time

	enginelog := g.Runtime.Root.Logger.New().With(log.Fields{
		"GameID": g.ID(),
	}).Tag("engine")
	enginelog.Info("start")

	activate(g) // bootstrap
loop: // nested break
	for { // event loop
		tStep = time.Now()
		select { // read request chan OR tick timer
		case r, ok := <-g.Monitor():
			if !ok {
				g.Log().Info("closed")
				break loop // nested break
			}
			g.Log().With(log.Fields{
				"Path":     r.URI,
				"Username": r.Username,
			}).Debug("engine: request")
			stack(g, Request(g, g.GetSeat(r.Username), r.URI, r.Data))
			continue
		case <-one2nd.C:
		} // time passed, requests were processed
		g.State.Timer -= time.Now().Sub(tStep)

		log := g.Log().Add("State", g.State.Print()).Tag("engine/tick")

		if len(g.State.Reacts) == len(g.Seats) {
			log.Debug("resolve")
		} else if g.State.Timer < time.Second {
			log.Warn("timeout")
		} else {
			continue
		} // move on

		// change state
		g.State.Timer = 0
		events := finish(g)                       // save event stack
		if state := g.State.Stack; state != nil { // stack pop
			log.Add("NextState", state).Debug("stackpop")
			g.State = state
			g.State.Timer = g.Runtime.Timeout
			g.State.Reacts = make(map[string]string)
			g.SendAll(game.BuildStateUpdate(g.State))
		} else if e := g.State.Event.GetNext(g); e != nil { // next state
			log.Add("Next", e).Debug("getnext")
			g.State = g.NewState(e)
			events = append(events, activate(g)...)
		} else {
			log.Info("exit")
			break loop // nested break
		}

		// stack events
		if len(events) > 0 {
			g.Log().Add("Events", events).Debug("stacking")
			stack(g, events)
		}
	} // event loop

	enginelog.Add("Data", g.State.Event.JSON()).Add("Time", time.Now().Sub(tStart)).Info("end")
	for _, seat := range g.Seats {
		if seat.Receiver != nil {
			seat.Receiver.WriteJSON(game.BuildGameUpdate(nil, seat.Username))
		}
		seat.Receiver = nil
	}
}

func connect(g *game.T, seat *game.Seat) {
	log := g.Log().With(log.Fields{
		"State": g.State.Print(),
		"Seat":  seat, // seat could be nil
	}).Tag("engine/connect")
	if connector, _ := g.State.Event.(game.ConnectEventer); connector != nil {
		log.Debug()
		connector.OnConnect(g, seat)
	} else {
		log.Debug("empty")
	}
}

func stack(g *game.T, events []game.Event) {
	if events == nil {
		return
	}
	g.Log().With(log.Fields{
		"State": g.State,
		"Stack": events,
	}).Debug("engine: stack")
	for _, e := range events {
		s := g.NewState(e)
		s.Stack = g.State
		g.State = s
		stack(g, activate(g))
	}
}

func finish(g *game.T) []game.Event {
	log := g.Log().Add("State", g.State.Print()).Tag("engine/finish")
	if finisher, _ := g.State.Event.(game.FinishEventer); finisher != nil {
		log.Debug()
		return finisher.Finish(g)
	}
	log.Debug("empty")
	return nil
}

func activate(g *game.T) []game.Event {
	g.State.Timer = g.Runtime.Timeout
	g.State.Reacts = make(map[string]string)
	g.SendAll(game.BuildStateUpdate(g.State))
	if activator, ok := g.State.Event.(game.ActivateEventer); ok {
		g.Log().Add("State", g.State.Print()).Debug("engine/activate")
		if events := activator.OnActivate(g); len(events) > 0 {
			return events
		}
	} else {
		g.Log().Add("State", g.State.Print()).Debug("engine/activate: none")
	}
	return nil
}
