package engine

import (
	"time"

	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/update"
	"ztaylor.me/cast"
)

// Run the engine
func Run(g *game.T) {
	tStart := time.Now()
	one2nd := time.NewTicker(time.Second)
	var tStep time.Time

	enginelog := g.Runtime.Root.Logger.New().With(cast.JSON{
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
			g.Log().With(cast.JSON{
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
			update.State(g)
			connect(g, nil)
		} else if e := g.State.R.GetNext(g); e != nil { // next state
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

	enginelog.Add("Data", g.State.R.JSON()).Add("Time", time.Now().Sub(tStart)).Info("end")
	for _, seat := range g.Seats {
		if seat.Receiver != nil {
			seat.Receiver.WriteJSON(update.Build("/game", nil))
		}
		seat.Receiver = nil
	}
}
